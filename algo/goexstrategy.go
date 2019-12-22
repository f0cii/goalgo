package algo

import (
	"log"

	"github.com/frankrap/goalgo"

	"github.com/nntaoli-project/GoEx"
	"github.com/nntaoli-project/GoEx/builder"
)

var (
	goExExchangeNameMap = map[string]string{
		"okex":     "okex.com",
		"huobi":    "huobi.pro",
		"bitstamp": "bitstamp.net",
		"kraken":   "kraken.com",
		"zb":       "zb.com",
		"bitfinex": "bitfinex.com",
		"binance":  "binance.com",
		"poloniex": "poloniex.com",
		"coinex":   "coinex.com",
		"bithumb":  "bithumb.com",
		"gate":     "gate.io",
		"bittrex":  "bittrex.com",
		"gdax":     "gdax.com",
		"wex":      "wex.nz",
		"big.one":  "big.one",
		"58coin":   "58coin.com",
		"fcoin":    "fcoin.com",
		"hitbtc":   "hitbtc.com",
	}
)

type GoExStrategy struct {
	goalgo.BaseStrategy
	Exchange  goex.API
	Exchanges []goex.API
}

func (s *GoExStrategy) Setup(params []goalgo.ExchangeParams) error {
	log.Printf("GoExStrategy Setup")
	s.Exchanges = []goex.API{}
	for _, p := range params {
		log.Print(p)
		s.Exchanges = append(s.Exchanges, s.createExchange(p))
	}
	if len(s.Exchanges) > 0 {
		s.Exchange = s.Exchanges[0]
	}
	return nil
}

func (s *GoExStrategy) createExchange(params goalgo.ExchangeParams) goex.API {
	b := builder.NewAPIBuilder()
	exName, ok := goExExchangeNameMap[params.Name]
	if !ok {
		return nil
	}
	return b.APIKey(params.AccessKey).APISecretkey(params.SecretKey).Build(exName)
}
