package goalgo

import (
	"errors"
)

var (
	// ErrNotFound is returned when an item or index is not in the database.
	ErrNotFound = errors.New("not found")
)

// GetValue 获取一个全局变量
func GetValue(key string) (Value, error) {
	value, err := GetClient().GetValue(key)
	if err != nil && err == ErrNotFound {
		return Value{}, nil
	}
	return value, err
}

// SetValue 设置一个全局变量
func SetValue(key string, value Value) error {
	return GetClient().SetValue(key, value)
}
