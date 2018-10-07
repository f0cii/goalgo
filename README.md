# goalgo
A real-time quantitative trading platform in Golang.

这是一个严肃的基于Go的量化系统，目前正在内部使用。

目前支持基于以下两个库编写策略：
1、GoEx https://github.com/nntaoli-project/GoEx
2、coinex https://github.com/sumorf/coinex

其中coinex提供了对BitMEX交易所的支持，内部集成了Rest和WebSocket接口

欢迎交流，预使用此系统，请自行登录以下官网注册账号:

http://39.104.22.80:8000

示例策略脚本:

```go
package main

import (
	"time"

	"github.com/sumorf/goalgo"
	"github.com/sumorf/goalgo/algo"
	"github.com/sumorf/goalgo/log"
)

// DemoStrategy 演示策略
type DemoStrategy struct {
	algo.GoExStrategy
}

// Init 策略初始化方法，必须实现
func (g *DemoStrategy) Init() error {
	log.Info("Init")
	return nil
}

// Run 策略主逻辑，必须实现
func (g *DemoStrategy) Run() error {
	log.Info("Run")
	for g.IsRunning() {
		time.Sleep(3 * time.Second)
		log.Info("Run do")
	}
	return nil
}

func main() {
	s := &DemoStrategy{}
	goalgo.Serve(s)
}
```
