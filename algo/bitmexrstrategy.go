package algo

import (
	stdlog "log"

	"github.com/frankrap/bitmex-api"
	"github.com/frankrap/goalgo"
	"github.com/frankrap/goalgo/log"
)

// BitMEXRStrategy BitMEX策略基类，此版本不启动WS
type BitMEXRStrategy struct {
	goalgo.BaseStrategy
	Exchange  *bitmex.BitMEX
	Exchanges []*bitmex.BitMEX
}

func (s *BitMEXRStrategy) Setup(params []goalgo.ExchangeParams) error {
	stdlog.Printf("BitMEXRStrategy Setup")
	s.Exchanges = []*bitmex.BitMEX{}
	for _, p := range params {
		stdlog.Print(p)
		var ex *bitmex.BitMEX
		switch p.Name {
		case "bitmex":
			ex = bitmex.New(nil, bitmex.HostReal, p.AccessKey, p.SecretKey, false)
		case "bitmex_test":
			ex = bitmex.New(nil, bitmex.HostTestnet, p.AccessKey, p.SecretKey, false)
		default:
			log.Errorf("交易所设置错误 %v", p.Name)
		}
		s.Exchanges = append(s.Exchanges, ex)

		if ex == nil {
			log.Errorf("创建交易所失败 ex == nil")
			continue
		}
	}
	if len(s.Exchanges) > 0 {
		s.Exchange = s.Exchanges[0]
	}
	return nil
}
