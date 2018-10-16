package goalgo

// GetValue 获取一个全局变量
func GetValue(key string) (Value, error) {
	return GetClient().GetValue(key)
}

// SetValue 设置一个全局变量
func SetValue(key string, value Value) error {
	return GetClient().SetValue(key, value)
}
