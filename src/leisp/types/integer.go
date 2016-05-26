// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import "fmt"

type IntegerValue struct {
	Value int64
}

func (v *IntegerValue) ToString() string {
	return fmt.Sprint(v.Value)
}

func (v *IntegerValue) GetType() string {
	return "integer"
}

func (v *IntegerValue) IsValue() bool {
	return true
}

func (v *IntegerValue) To(t string) (ValueType, error) {
	return nil, fmt.Errorf("cannot convert integer to %s: does not implement yet", t)
}

func (v *IntegerValue) EqualTo(t ValueType) bool {
	if v2, ok := t.(*IntegerValue); ok {
		if v2.Value == v.Value {
			return true
		}
	}
	return false
}

func NewIntegerValue(v int64) *IntegerValue {
	return &IntegerValue{Value: v}
}
