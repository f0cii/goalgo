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
