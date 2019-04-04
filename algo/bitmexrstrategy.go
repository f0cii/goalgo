package algo

import (
	stdlog "log"

	"github.com/sumorf/bitmexwrap/bitmex"
	"github.com/sumorf/goalgo"
	"github.com/sumorf/goalgo/log"
)

// BitMEXRStrategy BitMEX策略基类，此版本不启动WS
type BitMEXRStrategy struct {
	goalgo.BaseStrategy
	Exchange  *bitmex.Bitmex
	Exchanges []*bitmex.Bitmex
}

func (s *BitMEXRStrategy) Setup(params []goalgo.ExchangeParams) error {
	stdlog.Printf("BitMEXRStrategy Setup")
	s.Exchanges = []*bitmex.Bitmex{}
	for _, p := range params {
		stdlog.Print(p)
		var ex *bitmex.Bitmex
		switch p.Name {
		case "bitmex":
			ex = bitmex.NewBitmex(p.AccessKey, p.SecretKey)
		case "bitmex_test":
			ex = bitmex.NewBitmexTest(p.AccessKey, p.SecretKey)
		default:
			log.Errorf("交易所设置错误 %v", p.Name)
		}
		s.Exchanges = append(s.Exchanges, ex)

		if ex == nil {
			log.Errorf("创建交易所失败 ex == nil")
			continue
		}

		proxy := s.GetProxy()
		if proxy != "" {
			ex.SetProxy(proxy)
		}
	}
	if len(s.Exchanges) > 0 {
		s.Exchange = s.Exchanges[0]
	}
	return nil
}
