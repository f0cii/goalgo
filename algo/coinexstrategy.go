package algo

import (
	"github.com/sumorf/coinex/bitmex"
	"github.com/sumorf/goalgo"
	"github.com/sumorf/goalgo/log"
)

type CoinEXStrategy struct {
	goalgo.BaseStrategy
	Exchange  *bitmex.Bitmex
	Exchanges []*bitmex.Bitmex
}

func (s *CoinEXStrategy) Setup(params []goalgo.ExchangeParams) error {
	//log.Printf("CoinEXStrategy Setup")
	s.Exchanges = []*bitmex.Bitmex{}
	for _, p := range params {
		//log.Print(p)
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
			continue
		}

		err := ex.StartWS()
		if err != nil {
			log.Errorf("StartWS error: %v", err)
		}
	}
	if len(s.Exchanges) > 0 {
		s.Exchange = s.Exchanges[0]
	}
	return nil
}
