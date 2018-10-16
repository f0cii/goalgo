package goalgo

import (
	"bytes"
	"testing"

	"github.com/vmihailenco/msgpack"
)

func TestValueBase(t *testing.T) {
	v := Value{}
	v.SetValue("a")

	if v.ToString() != "a" {
		t.Errorf("%v", v)
		return
	}

	v.SetValue("b")
	if v.ToString() != "b" {
		t.Errorf("%v", v)
		return
	}

	v.SetValue(int(1))
	if v.ToInt() != 1 {
		t.Errorf("%v", v)
		return
	}

	v.SetValue(int32(1))
	if v.ToInt() != 1 {
		t.Errorf("%v", v)
		return
	}

	v.SetValue(float32(1.5))
	if v.ToFloat64() != 1.5 {
		t.Errorf("%v", v)
		return
	}

	v.SetValue(float64(3.5))
	if v.ToFloat64() != 3.5 {
		t.Errorf("%v", v)
		return
	}
}

func TestValue(t *testing.T) {
	v := Float64Value(1.2)
	writer := bytes.NewBuffer(nil)
	enc := msgpack.NewEncoder(writer)
	var err error
	err = enc.Encode(&v)
	if err != nil {
		t.Error(err)
		return
	}

	data := writer.Bytes()

	r := bytes.NewBuffer(data)
	dec := msgpack.NewDecoder(r)
	var v1 Value
	err = dec.Decode(&v1)
	if err != nil {
		t.Error(err)
		return
	}

	s := v1.ToString()
	if s != "1.2" {
		t.Errorf("%v", s)
		return
	}
	f1 := v1.ToFloat64()
	if f1 != 1.2 {
		t.Errorf("%v", f1)
		return
	}
}
