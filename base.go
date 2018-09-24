package goalgo

// RobotState 机器人状态
type RobotState int

const (
	// RStateOff 无
	RStateOff RobotState = iota
	// RStateInitialized 初始化
	RStateInitialized
	// RStateRunning 运行中
	RStateRunning
	// RStateStopRequested 请求停止
	RStateStopRequested
	// RStateStopped 停止
	RStateStopped
	// RStateError 出错
	RStateError
)
