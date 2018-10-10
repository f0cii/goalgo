# goalgo

A real-time quantitative trading platform in Golang.

这是一个严肃的基于 Go 的量化系统，目前正在内部使用。

目前支持基于以下两个库编写策略

GoEx https://github.com/nntaoli-project/GoEx

coinex https://github.com/SuperGod/coinex

其中 coinex 提供了对 BitMEX 交易所的支持，内部集成了 Rest 和 WebSocket 接口

欢迎交流，欲使用此系统，请自行登录以下官网注册账号:

http://39.104.22.80:8000

1.简单的基于GoEx策略脚本:

```go
package main

import (
	"time"

	"github.com/nntaoli-project/GoEx"
	"github.com/sumorf/goalgo"
	"github.com/sumorf/goalgo/algo"
	"github.com/sumorf/goalgo/log"
)

// SimpleGoExStrategy 简单的GoEx策略
type SimpleGoExStrategy struct {
	algo.GoExStrategy
}

// Init 策略初始化方法，必须实现
func (s *SimpleGoExStrategy) Init() error {
	log.Info("Init")
	return nil
}

// Run 策略主逻辑，必须实现
func (s *SimpleGoExStrategy) Run() error {
	log.Info("Run")
	for s.IsRunning() {
		time.Sleep(5 * time.Second)
		// 获取交易所 Ticker 数据
		ticker, err := s.Exchange.GetTicker(goex.BTC_USDT)
		if err != nil {
			log.Errorf("%v", err)
			continue
		}
		log.Infof("Ticker %#v", ticker)
	}
	return nil
}

func main() {
	s := &SimpleGoExStrategy{}
	goalgo.Serve(s)
}
```

2.简单的基于coinex(BitMEX期货市场)策略脚本:

```go
package main

import (
	"time"

	"github.com/sumorf/goalgo"
	"github.com/sumorf/goalgo/algo"
	"github.com/sumorf/goalgo/log"
)

// CoinEXDemoStrategy 示例策略(BitMEX)
type CoinEXDemoStrategy struct {
	algo.CoinEXStrategy
}

// Init 策略初始化方法，必须实现
func (s *CoinEXDemoStrategy) Init() error {
	log.Info("Init")
	banlance, err := s.Exchange.ContractBalances()
	if err != nil {
		log.Errorf("ContractBalances error: %v", err)
		return err
	}
	log.Infof("banlance: %v", banlance)

	return nil
}

// Run 策略主逻辑，必须实现
func (s *CoinEXDemoStrategy) Run() error {
	log.Info("Run")
	for s.IsRunning() {
		time.Sleep(3 * time.Second)
		ticker, err := s.Exchange.Ticker()
		if err != nil {
			log.Errorf("%v", err)
			continue
		}
		log.Infof("ticker: %v", ticker)
	}
	return nil
}

func main() {
	s := &CoinEXDemoStrategy{}
	goalgo.Serve(s)
}
```
