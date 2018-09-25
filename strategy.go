package goalgo

// ExchangeParams 初始化交易所所需参数
type ExchangeParams struct {
	Label     string
	Name      string
	AccessKey string
	SecretKey string
}

// Strategy 策略接口
type Strategy interface {
	StrategyCtl
	SetSelf(self Strategy)
	Setup(params []ExchangeParams) error
	Init() error
	Run() error
}
