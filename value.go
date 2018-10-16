package goalgo

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/vmihailenco/msgpack"
)

type ValueKind uint32

const (
	ValueNil ValueKind = iota
	ValueBoolean
	ValueString
	ValueInt32
	ValueInt64
	ValueUint32
	ValueUint64
	ValueFloat32
	ValueFloat64
)

type Value struct {
	Kind  ValueKind
	Value interface{}
}

var (
	_ msgpack.CustomEncoder = &Value{}
	_ msgpack.CustomDecoder = &Value{}
)

func (r *Value) EncodeMsgpack(enc *msgpack.Encoder) error {
	return enc.Encode(r.Kind, r.Value)
}

func (r *Value) DecodeMsgpack(enc *msgpack.Decoder) error {
	return enc.Decode(&r.Kind, &r.Value)
}

func (value *Value) ToBool() bool {
	return ToBool(value.Value)
}

func (value *Value) ToString() string {
	if value.Kind == ValueString {
		return value.Value.(string)
	}
	if v, ok := value.Value.(string); ok {
		return v
	}
	return fmt.Sprintf("%v", value.Value)
}

func (value *Value) ToInt() int {
	return ToInt(value.Value)
}

func (value *Value) ToFloat64() float64 {
	return ToFloat(value.Value)
}

func (value *Value) SetValue(val interface{}) {
	switch v := val.(type) {
	case bool:
		value.Kind = ValueBoolean
		value.Value = v
	case string:
		log.Printf("string %#v", val)
		value.Kind = ValueString
		value.Value = v
	case int:
		value.Kind = ValueInt32
		value.Value = int32(v)
	case int32:
		value.Kind = ValueInt32
		value.Value = v
	case int64:
		value.Kind = ValueInt64
		value.Value = v
	case float32:
		value.Kind = ValueFloat32
		value.Value = v
	case float64:
		value.Kind = ValueFloat64
		value.Value = v
	default:
		log.Printf("%#v", val)
		value.Value = ""
	}
}

func BoolValue(value bool) Value {
	return Value{Value: value, Kind: ValueBoolean}
}

func StringValue(value string) Value {
	return Value{Value: value, Kind: ValueString}
}

func Int32Value(value int32) Value {
	return Value{Value: value, Kind: ValueInt32}
}

func Int64Value(value int64) Value {
	return Value{Value: value, Kind: ValueInt64}
}

func Float32Value(value float32) Value {
	return Value{Value: value, Kind: ValueFloat32}
}

func Float64Value(value float64) Value {
	return Value{Value: value, Kind: ValueFloat64}
}

// ToBool convert any type to boolean
func ToBool(value interface{}) bool {
	switch value := value.(type) {
	case bool:
		return value
	case *bool:
		return *value
	case string:
		switch value {
		case "", "false":
			return false
		}
		return true
	case *string:
		return ToBool(*value)
	case float64:
		if value != 0 {
			return true
		}
		return false
	case *float64:
		return ToBool(*value)
	case float32:
		if value != 0 {
			return true
		}
		return false
	case *float32:
		return ToBool(*value)
	case int:
		if value != 0 {
			return true
		}
		return false
	case *int:
		return ToBool(*value)
	}
	return false
}

// ToInt convert any type to int
func ToInt(value interface{}) int {
	switch value := value.(type) {
	case bool:
		if value == true {
			return 1
		}
		return 0
	case int:
		if value < int(math.MinInt32) || value > int(math.MaxInt32) {
			return 0 // nil
		}
		return value
	case *int:
		return ToInt(*value)
	case int8:
		return int(value)
	case *int8:
		return int(*value)
	case int16:
		return int(value)
	case *int16:
		return int(*value)
	case int32:
		return int(value)
	case *int32:
		return int(*value)
	case int64:
		if value < int64(math.MinInt32) || value > int64(math.MaxInt32) {
			return 0 // nil
		}
		return int(value)
	case *int64:
		return ToInt(*value)
	case uint:
		if value > math.MaxInt32 {
			return 0 // nil
		}
		return int(value)
	case *uint:
		return ToInt(*value)
	case uint8:
		return int(value)
	case *uint8:
		return int(*value)
	case uint16:
		return int(value)
	case *uint16:
		return int(*value)
	case uint32:
		if value > uint32(math.MaxInt32) {
			return 0 // nil
		}
		return int(value)
	case *uint32:
		return ToInt(*value)
	case uint64:
		if value > uint64(math.MaxInt32) {
			return 0 // nil
		}
		return int(value)
	case *uint64:
		return ToInt(*value)
	case float32:
		if value < float32(math.MinInt32) || value > float32(math.MaxInt32) {
			return 0 // nil
		}
		return int(value)
	case *float32:
		return ToInt(*value)
	case float64:
		if value < float64(math.MinInt32) || value > float64(math.MaxInt32) {
			return 0 // nil
		}
		return int(value)
	case *float64:
		return ToInt(*value)
	case string:
		val, err := strconv.ParseFloat(value, 0)
		if err != nil {
			return 0 // nil
		}
		return ToInt(val)
	case *string:
		return ToInt(*value)
	}

	// If the value cannot be transformed into an int, return nil instead of '0'
	// to denote 'no integer found'
	return 0 // nil
}

// ToFloat convert any type to float
func ToFloat(value interface{}) float64 {
	switch value := value.(type) {
	case bool:
		if value == true {
			return 1.0
		}
		return 0.0
	case *bool:
		return ToFloat(*value)
	case int:
		return float64(value)
	case *int32:
		return ToFloat(*value)
	case float32:
		return ToFloat(fmt.Sprintf("%v", value))
	case *float32:
		return ToFloat(*value)
	case float64:
		return value
	case *float64:
		return ToFloat(*value)
	case string:
		val, err := strconv.ParseFloat(value, 0)
		if err != nil {
			return 0
		}
		return val
	case *string:
		return ToFloat(*value)
	}
	return 0.0
}
