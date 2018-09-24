package goalgo

import (
	"log"
	"sync"

	"github.com/hashicorp/go-plugin"
)

// BaseStrategy 策略基础类
type BaseStrategy struct {
	self  interface{}
	mutex sync.RWMutex
	state RobotState
}

// SetSelf 设置 self 对象
func (s *BaseStrategy) SetSelf(self Strategy) {
	s.self = self.(interface{})
}

// GetState 获取策略状态
func (s *BaseStrategy) GetState() RobotState {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.state
}

// IsRunning 是否运行中
func (s *BaseStrategy) IsRunning() bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.state == RStateRunning
}

// Start 启动
func (s *BaseStrategy) Start() plugin.BasicError {
	go s.run()
	return plugin.BasicError{}
}

func (s *BaseStrategy) run() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Run error: %v", err)
			s.state = RStateStopped
		}
	}()
	log.Printf("Start")
	if s.self == nil {
		log.Printf("The strategy this is nil")
		s.state = RStateStopped
		return
	}
	strategy, ok := s.self.(Strategy)
	if !ok {
		log.Printf("The strategy does not implement Strategy")
		s.state = RStateStopped
		return
	}
	s.state = RStateRunning
	strategy.Init()
	strategy.Setup(nil)
	s.state = RStateInitialized
	strategy.Run()
	log.Printf("Run done")
	s.state = RStateStopped
}

// Stop 停止
func (s *BaseStrategy) Stop() plugin.BasicError {
	log.Printf("Stop")
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.state == RStateStopped {
		return plugin.BasicError{}
	}
	if s.state != RStateRunning {
		return plugin.BasicError{Message: "State error"}
	}
	s.state = RStateStopRequested
	return plugin.BasicError{}
}

// Pause 暂停
func (s *BaseStrategy) Pause() plugin.BasicError {
	return plugin.BasicError{}
}
