// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import "fmt"

type FloatValue struct {
	Value float64
}

func (v *FloatValue) ToString() string {
	return fmt.Sprint(v.Value)
}

func (v *FloatValue) GetType() string {
	return "float"
}

func (v *FloatValue) IsValue() bool {
	return true
}

func (v *FloatValue) To(t string) (ValueType, error) {
	return nil, fmt.Errorf("cannot convert float to %s: does not implement yet", t)
}

func NewFloatValue(v float64) *FloatValue {
	return &FloatValue{Value: v}
}
