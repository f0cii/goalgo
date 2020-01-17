# goalgo

A real-time quantitative trading platform in Golang.

这是一个严肃的基于 Go 的量化系统，目前正在内部使用。

目前支持基于以下两个库编写策略

GoEx https://github.com/nntaoli-project/GoEx

bitmex-api https://github.com/frankrap/bitmex-api

其中 bitmexwrap 提供了对 BitMEX 交易所的支持，内部集成了 Rest 和 WebSocket 接口

欢迎交流，欲使用此系统，请自行登录以下官网注册账号:

https://quant.coinrust.com (演示站点)

如有问题，请联系 QQ: 529808348

1.简单的基于 GoEx 策略脚本:

```go
package main

import (
	"time"

	"github.com/nntaoli-project/GoEx"
	"github.com/frankrap/goalgo"
	"github.com/frankrap/goalgo/algo"
	"github.com/frankrap/goalgo/log"
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

2.简单的基于 bitmexwrap(BitMEX 期货市场)策略脚本:

```go
package main

import (
	"time"

	"github.com/frankrap/goalgo"
	"github.com/frankrap/goalgo/algo"
	"github.com/frankrap/goalgo/log"
)

// BitMEXDemoStrategy 示例策略(BitMEX)
type BitMEXDemoStrategy struct {
	algo.BitMEXStrategy
}

// Init 策略初始化方法，必须实现
func (s *BitMEXDemoStrategy) Init() error {
	log.Info("Init")
	balance, err := s.Exchange.Balance()
	if err != nil {
		log.Errorf("Balance error: %v", err)
		return err
	}
	log.Infof("balance: %v", balance)

	return nil
}

// Run 策略主逻辑，必须实现
func (s *BitMEXDemoStrategy) Run() error {
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
	s := &BitMEXDemoStrategy{}
	goalgo.Serve(s)
}
```

### Thanks
> This project uses the following external projects or products
1. [RabbitMQ] https://www.rabbitmq.com/
2. [PostgreSQL] https://www.postgresql.org/
3. [ClickHouse] https://clickhouse.yandex/
4. [Docker] https://www.docker.com/
5. [GoEx] https://github.com/nntaoli-project/GoEx
6. [coinex] https://github.com/SuperGod/coinex
7. [porthos-go] https://github.com/porthos-rpc/porthos-go
8. [go-plugin] https://github.com/hashicorp/go-plugin
9. [gin] https://github.com/gin-gonic/gin
10. [vuejs] https://vuejs.org/
11. ...

### 捐赠

> BTC:17XauetDWtwAZKHRyEw3DpYcRqzhKZzPHq

> ETH:0x5e065711852fda9b75c4e490b1e35eea89b9aaa5

> XRP:rshdB4kyxXkuPCtCgZw8DrcuSo7jEWfqck
