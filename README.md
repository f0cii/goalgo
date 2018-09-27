# goalgo
A real-time quantitative trading platform in Golang.

请登录官方网站:

http://39.104.22.80:8000

示例策略脚本:

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
