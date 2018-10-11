package goalgo

import (
	"fmt"
	"sync"

	"github.com/sumorf/goalgo/log"

	"github.com/hashicorp/go-plugin"
)

// BaseStrategy 策略基础类
type BaseStrategy struct {
	self   interface{}
	mutex  sync.RWMutex
	status RobotStatus
}

// SetSelf 设置 self 对象
func (s *BaseStrategy) SetSelf(self Strategy) {
	s.self = self.(interface{})
}

// GetState 获取策略状态
func (s *BaseStrategy) GetState() RobotStatus {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.status
}

// IsRunning 是否运行中
func (s *BaseStrategy) IsRunning() bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.status == RobotStatusRunning
}

// Start 启动
func (s *BaseStrategy) Start() plugin.BasicError {
	go s.run()
	return plugin.BasicError{}
}

func (s *BaseStrategy) run() {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("Run error: %v", err)
			s.status = RobotStatusStopped
		}
	}()

	//log.Info("Start")

	if s.self == nil {
		log.Errorf("The strategy this is nil")
		s.status = RobotStatusStopped
		s.updateStatus(s.status)
		return
	}

	strategy, ok := s.self.(Strategy)
	if !ok {
		log.Errorf("The strategy does not implement Strategy")
		s.status = RobotStatusStopped
		s.updateStatus(s.status)
		return
	}

	var rError error

	func() {
		defer func() {
			if r := recover(); r != nil {
				rError = fmt.Errorf("%v", r)
			}
		}()
		client := GetClient()
		exchanges, err := client.GetRobotExchangeInfo("", id)
		if err != nil {
			log.Errorf("GetRobotExchangeInfo error: %v", err)
		} else {
			params := []ExchangeParams{}
			for _, ex := range exchanges {
				params = append(params, ExchangeParams{
					Label:     ex.Label,
					Name:      ex.Name,
					AccessKey: ex.AccessKey,
					SecretKey: ex.SecretKey,
				})
			}
			//log.Info("Setup...")
			rError = strategy.Setup(params)
			//log.Info("Setup.")
		}
	}()

	if rError != nil {
		log.Errorf("%v", rError)
		s.status = RobotStatusError
		s.updateStatus(s.status)
		return
	}

	func() {
		defer func() {
			if r := recover(); r != nil {
				rError = fmt.Errorf("%v", r)
			}
		}()
		rError = strategy.Init()
	}()

	if rError != nil {
		log.Errorf("%v", rError)
		s.status = RobotStatusError
		s.updateStatus(s.status)
		return
	}

	s.status = RobotStatusRunning

	func() {
		defer func() {
			if r := recover(); r != nil {
				rError = fmt.Errorf("%v", r)
			}
		}()
		rError = strategy.Run()
	}()

	//log.Info("Run done")

	if rError != nil {
		s.status = RobotStatusError
		log.Errorf("%v", rError)
	} else {
		s.status = RobotStatusStopped
		log.Info("Stopped")
	}

	// 同步状态
	s.updateStatus(s.status)
}

func (s *BaseStrategy) updateStatus(status RobotStatus) {
	client := GetClient()
	client.UpdateStatus(id, status)
}

// Stop 停止
func (s *BaseStrategy) Stop() plugin.BasicError {
	log.Info("OnStop")
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.status == RobotStatusStopped {
		return plugin.BasicError{}
	}
	if s.status != RobotStatusRunning {
		return plugin.BasicError{Message: "State error"}
	}
	s.status = RobotStatusRequested
	return plugin.BasicError{}
}

// Pause 暂停
func (s *BaseStrategy) Pause() plugin.BasicError {
	return plugin.BasicError{}
}
