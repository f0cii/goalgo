package goalgo

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/sumorf/goalgo/log"

	"runtime/debug"

	"github.com/Workiva/go-datastructures/queue"
	"github.com/hashicorp/go-plugin"
)

const (
	// OptionTag 选项Tag
	OptionTag = "option"
)

// OptionInfo 参数信息
type OptionInfo struct {
	Name         string      `json:"name"`
	Description  string      `json:"description"`
	Type         string      `json:"type"`
	Value        interface{} `json:"value"`
	DefaultValue interface{} `json:"default_value"`
}

// AfterOptionsChanged 策略参数改变接口，实现此接口的策略，在动态改变参数后触发此方法
type AfterOptionsChanged interface {
	AfterOptionsChanged()
}

// BaseStrategy 策略基础类
type BaseStrategy struct {
	self         interface{}
	mutex        sync.RWMutex
	commandQueue queue.Queue
	status       RobotStatus

	proxy string
}

// SetSelf 设置 self 对象
func (s *BaseStrategy) SetSelf(self Strategy) {
	s.self = self.(interface{})
}

// SetProxy 设置访问网络的代理(主要方便本地测试，如：http://127.0.0.1:1080)
func (s *BaseStrategy) SetProxy(proxy string) {
	s.proxy = proxy
}

// GetProxy 获取当前代理
func (s *BaseStrategy) GetProxy() string {
	return s.proxy
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

// GetOptions 获取参数
func (s *BaseStrategy) GetOptions() (optionMap map[string]*OptionInfo) {
	//log.Info("GetOptions")
	optionMap = map[string]*OptionInfo{}

	if s.self == nil {
		return
	}

	val := reflect.ValueOf(s.self)

	// If it's an interface or a pointer, unwrap it.
	if val.Kind() == reflect.Ptr && val.Elem().Kind() == reflect.Struct {
		val = val.Elem()
	} else {
		return
	}

	valNumFields := val.NumField()

	for i := 0; i < valNumFields; i++ {
		field := val.Field(i)
		fieldKind := field.Kind()
		if fieldKind == reflect.Ptr && field.Elem().Kind() == reflect.Struct {
			continue
		}

		typeField := val.Type().Field(i)
		fieldName := typeField.Name
		tag := typeField.Tag

		if !field.CanInterface() {
			continue
		}

		option := tag.Get(OptionTag)

		if option == "" {
			continue
		}

		var description string
		var defaultValueString string
		index := strings.Index(option, ",")
		//fmt.Printf("tag: %v i: %v\n", option, index)
		if index != -1 {
			description = option[0:index]
			defaultValueString = option[index+1:]
		} else {
			description = option
		}
		value := field.Interface()
		defaultValue := s.getDefaultValue(fieldKind, defaultValueString)

		optionMap[fieldName] = &OptionInfo{
			Name:         fieldName,
			Description:  description,
			Type:         typeField.Type.String(),
			Value:        value,
			DefaultValue: defaultValue,
		}
		//log.Infof("F: %v V: %v", fieldName, value)
	}

	return
}

func (s *BaseStrategy) getDefaultValue(kind reflect.Kind, value string) interface{} {
	switch kind {
	case reflect.Bool:
		return ToBool(value)
	case reflect.String:
		return value
	case reflect.Int:
		return ToInt(value)
	case reflect.Int8:
		return int8(ToInt(value))
	case reflect.Int16:
		return int16(ToInt(value))
	case reflect.Int32:
		return int32(ToInt(value))
	case reflect.Int64:
		return ToInt64(value)
	case reflect.Uint:
		return uint(ToInt(value))
	case reflect.Uint8:
		return uint8(ToInt(value))
	case reflect.Uint16:
		return uint16(ToInt(value))
	case reflect.Uint32:
		return uint32(ToInt(value))
	case reflect.Uint64:
		return uint64(ToInt64(value))
	case reflect.Float32:
		return ToFloat32(value)
	case reflect.Float64:
		return ToFloat(value)
	}
	return 0
}

// SetOptions 设置参数
func (s *BaseStrategy) SetOptions(options map[string]interface{}) plugin.BasicError {
	log.Info("SetOptions")

	if len(options) == 0 {
		return plugin.BasicError{}
	}

	rawOptions := s.GetOptions()

	// 反射成员变量
	val := reflect.ValueOf(s.self)

	// If it's an interface or a pointer, unwrap it.
	if val.Kind() == reflect.Ptr && val.Elem().Kind() == reflect.Struct {
		val = val.Elem()
	} else {
		return plugin.BasicError{}
	}

	for name, value := range options {
		var fieldName string

		if ipi, ok := rawOptions[name]; !ok {
			continue
		} else {
			fieldName = ipi.Name
		}

		//fmt.Println(fieldName)

		v := val.FieldByName(fieldName)
		if !v.IsValid() {
			continue
		}

		switch v.Kind() {
		default:
			fmt.Printf("Error Kind: %v\n", v.Kind())
		case reflect.Bool:
			v.SetBool(ToBool(value))
		case reflect.String:
			v.SetString(value.(string))
		case reflect.Int:
			v.SetInt(ToInt64(value))
		case reflect.Int8:
			v.SetInt(ToInt64(value))
		case reflect.Int16:
			v.SetInt(ToInt64(value))
		case reflect.Int32:
			v.SetInt(ToInt64(value))
		case reflect.Int64:
			v.SetInt(ToInt64(value))
		case reflect.Uint:
			v.SetUint(ToUint64(value))
		case reflect.Uint8:
			v.SetUint(ToUint64(value))
		case reflect.Uint16:
			v.SetUint(ToUint64(value))
		case reflect.Uint32:
			v.SetUint(ToUint64(value))
		case reflect.Uint64:
			v.SetUint(ToUint64(value))
		case reflect.Float32:
			v.SetFloat(ToFloat(value))
		case reflect.Float64:
			v.SetFloat(ToFloat(value))
			// case reflect.Struct:
			// 	v.Set(reflect.ValueOf(value))
		}
	}

	// 触发事件
	if v, ok := s.self.(AfterOptionsChanged); ok {
		v.AfterOptionsChanged()
	}

	return plugin.BasicError{}
}

// QueueCommand 命令入队列
func (s *BaseStrategy) QueueCommand(command string) plugin.BasicError {
	cmd := Command{}
	err := json.Unmarshal([]byte(command), &cmd)
	if err != nil {
		return plugin.BasicError{}
	}
	_ = s.commandQueue.Put(&cmd)
	return plugin.BasicError{}
}

// GetCommand 获取一个命令结构
func (s *BaseStrategy) GetCommand() *Command {
	_, err := s.commandQueue.Peek()
	if err != nil {
		return nil
	}
	result, err := s.commandQueue.Get(1)
	if err != nil {
		return nil
	}
	if len(result) != 1 {
		return nil
	}
	return result[0].(*Command)
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
		log.Error("The strategy this is nil")
		s.status = RobotStatusStopped
		s.updateStatus(s.status)
		return
	}

	strategy, ok := s.self.(Strategy)
	if !ok {
		log.Error("The strategy does not implement Strategy")
		s.status = RobotStatusStopped
		s.updateStatus(s.status)
		return
	}

	var rError error
	var pErr plugin.BasicError

	// Set options
	func() {
		defer func() {
			if r := recover(); r != nil {
				rError = fmt.Errorf("%v", r)
			}
		}()
		client := GetClient()
		options, err := client.GetRobotOptions("", id)
		if err != nil {
			log.Errorf("GetRobotOptions error: %v", err)
		}

		mOptions := map[string]interface{}{}
		for _, v := range options {
			mOptions[v.Key] = v.Value
		}

		pErr = strategy.SetOptions(mOptions)
	}()

	if rError != nil {
		log.Errorf("Setup error: %v", rError)
		s.status = RobotStatusError
		s.updateStatus(s.status)
		return
	}

	if pErr.Error() != "" {
		log.Errorf("SetOptions error: %v", pErr.Error())
		s.status = RobotStatusError
		s.updateStatus(s.status)
		return
	}

	// Setup
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
			var params []ExchangeParams
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
		log.Errorf("Setup error: %v", rError)
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
		log.Errorf("Init error: %v", rError)
		s.status = RobotStatusError
		s.updateStatus(s.status)
		return
	}

	s.status = RobotStatusRunning

	func() {
		defer func() {
			if r := recover(); r != nil {
				rError = fmt.Errorf("%v", r)
				log.Errorf("Run error: stack=%v", string(debug.Stack()))
			}
		}()
		rError = strategy.Run()
	}()

	//log.Info("Run done")

	if rError != nil {
		s.status = RobotStatusError
		log.Errorf("Run error: %v", rError)
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

// GetValue 获取一个全局变量
func (s *BaseStrategy) GetValue(key string) (Value, error) {
	return GetValue(key)
}

// SetValue 设置一个全局变量
func (s *BaseStrategy) SetValue(key string, value Value) error {
	return SetValue(key, value)
}

// UpdateStat 更新统计数据
func (s *BaseStrategy) UpdateStat(name string, value []byte) error {
	return GetClient().UpdateStat(name, value)
}
