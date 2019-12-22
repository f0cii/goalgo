package goalgo

import (
	stdlog "log"

	"github.com/frankrap/goalgo/log"
)

type GRPCLog struct {
	SID int
}

func (l *GRPCLog) Log(e *log.Entry) error {
	//stdlog.Println(e)
	err := GetClient().Log(l.SID, e.ID, e.Time.Unix(), int32(e.Level), e.Message)
	if err != nil {
		stdlog.Println(err)
	}
	return err
}
