package goalgo

// RobotStatus 机器人状态
type RobotStatus int8

const (
	// RobotStatusDisabled 禁用
	RobotStatusDisabled RobotStatus = iota
	// RobotStatusStopped 停止
	RobotStatusStopped
	// RobotStatusStarting 启动中
	RobotStatusStarting
	// RobotStatusRunning 运行中
	RobotStatusRunning
	// RobotStatusRequested RobotStatusStopping 停止中
	RobotStatusRequested
	// RobotStatusError 出错
	RobotStatusError
)

// Command 表示一个交互命令结构
type Command struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
