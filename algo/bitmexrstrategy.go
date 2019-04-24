package algo

import (
	stdlog "log"

	"github.com/sumorf/bitmex-api"
	"github.com/sumorf/goalgo"
	"github.com/sumorf/goalgo/log"
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
			ex = bitmex.New(bitmex.HostReal, p.AccessKey, p.SecretKey)
		case "bitmex_test":
			ex = bitmex.New(bitmex.HostTestnet, p.AccessKey, p.SecretKey)
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
