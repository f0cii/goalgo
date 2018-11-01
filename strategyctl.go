package goalgo

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

const (
	PluginMapStrategyCtlKey = "strategyCtl"
)

// pluginMap is the map of plugins we can dispense.
var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  plugin.CoreProtocolVersion,
	MagicCookieKey:   "STRATEGY_PLUGIN",
	MagicCookieValue: "QtMWr4VtAC*r",
}

var PluginMap = map[string]plugin.Plugin{
	PluginMapStrategyCtlKey: &StrategyPlugin{},
}

// StrategyCtl 插件接口
type StrategyCtl interface {
	GetState() RobotStatus
	GetOptions() (optionMap map[string]*OptionInfo)
	SetOptions(options map[string]interface{}) plugin.BasicError
	QueueCommand(command string) plugin.BasicError
	Start() plugin.BasicError
	Stop() plugin.BasicError
	Pause() plugin.BasicError
}

// StrategyRPC Here is an implementation that talks over RPC
type StrategyRPC struct{ client *rpc.Client }

// GetState ...
func (g *StrategyRPC) GetState() RobotStatus {
	var resp RobotStatus
	err := g.client.Call("Plugin.GetState", new(interface{}), &resp)
	if err != nil {
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		return RobotStatusDisabled
	}

	return resp
}

// GetOptions ...
func (g *StrategyRPC) GetOptions() (optionMap map[string]*OptionInfo) {
	var resp map[string]*OptionInfo
	err := g.client.Call("Plugin.GetOptions", new(interface{}), &resp)
	if err != nil {
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		return map[string]*OptionInfo{}
	}

	return resp
}

// SetOptions ...
func (g *StrategyRPC) SetOptions(options map[string]interface{}) plugin.BasicError {
	var resp plugin.BasicError
	err := g.client.Call("Plugin.SetOptions", options, &resp)
	if err != nil {
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		return plugin.BasicError{Message: err.Error()}
	}

	return resp
}

// QueueCommand ...
func (g *StrategyRPC) QueueCommand(command string) plugin.BasicError {
	var resp plugin.BasicError
	err := g.client.Call("Plugin.QueueCommand", command, &resp)
	if err != nil {
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		return plugin.BasicError{Message: err.Error()}
	}

	return resp
}

// Start ...
func (g *StrategyRPC) Start() plugin.BasicError {
	var resp plugin.BasicError
	err := g.client.Call("Plugin.Start", new(interface{}), &resp)
	if err != nil {
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		//panic(err)
		return plugin.BasicError{Message: err.Error()}
	}

	return resp
}

// Stop ...
func (g *StrategyRPC) Stop() plugin.BasicError {
	var resp plugin.BasicError
	err := g.client.Call("Plugin.Stop", new(interface{}), &resp)
	if err != nil {
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		//panic(err)
		return plugin.BasicError{Message: err.Error()}
	}

	return resp
}

// Pause ...
func (g *StrategyRPC) Pause() plugin.BasicError {
	var resp plugin.BasicError
	err := g.client.Call("Plugin.Pause", new(interface{}), &resp)
	if err != nil {
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		//panic(err)
		return plugin.BasicError{Message: err.Error()}
	}

	return resp
}

// StrategyRPCServer Here is the RPC server that GreeterRPC talks to, conforming to
// the requirements of net/rpc
type StrategyRPCServer struct {
	// This is the real implementation
	Impl StrategyCtl
}

func (s *StrategyRPCServer) GetState(args interface{}, resp *RobotStatus) error {
	*resp = s.Impl.GetState()
	return nil
}

func (s *StrategyRPCServer) GetOptions(args interface{}, resp *map[string]*OptionInfo) error {
	*resp = s.Impl.GetOptions()
	return nil
}

func (s *StrategyRPCServer) SetOptions(args map[string]interface{}, resp *plugin.BasicError) error {
	*resp = s.Impl.SetOptions(args)
	return nil
}

func (s *StrategyRPCServer) QueueCommand(args string, resp *plugin.BasicError) error {
	*resp = s.Impl.QueueCommand(args)
	return nil
}

func (s *StrategyRPCServer) Start(args interface{}, resp *plugin.BasicError) error {
	*resp = s.Impl.Start()
	return nil
}

func (s *StrategyRPCServer) Stop(args interface{}, resp *plugin.BasicError) error {
	*resp = s.Impl.Stop()
	return nil
}

func (s *StrategyRPCServer) Pause(args interface{}, resp *plugin.BasicError) error {
	*resp = s.Impl.Pause()
	return nil
}

// This is the implementation of plugin.Plugin so we can serve/consume this
//
// This has two methods: Server must return an RPC server for this plugin
// type. We construct a GreeterRPCServer for this.
//
// Client must return an implementation of our interface that communicates
// over an RPC client. We return GreeterRPC for this.
//
// Ignore MuxBroker. That is used to create more multiplexed streams on our
// plugin connection and is a more advanced use case.
type StrategyPlugin struct {
	// Impl Injection
	Impl StrategyCtl
}

// Server server
func (p *StrategyPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &StrategyRPCServer{Impl: p.Impl}, nil
}

// Client client
func (StrategyPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &StrategyRPC{client: c}, nil
}
