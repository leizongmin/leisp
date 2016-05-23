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

func (v *IntegerValue) To(t string) (ValueType, error) {
	return nil, fmt.Errorf("cannot convert integer to %s: does not implement yet", t)
}

func NewIntegerValue(v int64) *IntegerValue {
	return &IntegerValue{Value: v}
}
