package goalgo

import (
	"log"
	"sync"
	"time"

	"errors"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	instance      *Client
	clientTimeout = time.Second * 5
	once          sync.Once
)

// Client RPC客户端对象
type Client struct {
	conn    *grpc.ClientConn
	client  RobotCtlClient
	timeout time.Duration
}

func (c Client) Close() {
	c.conn.Close()
}

func (c *Client) GetRobotExchangeInfo(uid string, id string) ([]*RobotExchangeInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	request := &RobotExchangeInfoRequest{
		Uid:     uid,
		RobotId: id,
	}
	r, err := c.client.GetRobotExchangeInfo(ctx, request)
	if err != nil {
		log.Printf("Log: %v", err)
		return nil, err
	}
	return r.GetExchanges(), nil
}

func (c *Client) Log(sid int, id uint64, tm int64, level int32, message string) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	request := &LogRequest{
		Sid:     int32(sid),
		Id:      id,
		Time:    tm,
		Level:   level,
		Message: message,
	}
	r, err := c.client.Log(ctx, request)
	if err != nil {
		log.Printf("Log: %v", err)
		return err
	}
	if r.Success {
		return nil
	}
	return errors.New(r.Message)
}

func (c *Client) UpdateStatus(robotID string, status RobotStatus) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	request := &UpdateStatusRequest{
		RobotId: robotID,
		Status:  int32(status),
	}
	r, err := c.client.UpdateStatus(ctx, request)
	if err != nil {
		log.Printf("Log: %v", err)
		return err
	}
	if r.Success {
		return nil
	}
	return errors.New(r.Message)
}

// GetClient 获得RPC客户端对象
func GetClient() *Client {
	once.Do(func() {
		instance = newClient()
	})
	return instance
}

func newClient() *Client {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	//defer conn.Close()
	c := NewRobotCtlClient(conn)
	return &Client{client: c, conn: conn, timeout: clientTimeout}
}
