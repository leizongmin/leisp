// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import "fmt"

type FloatValue struct {
	Value float64
}

func (t *FloatValue) ToString() string {
	return fmt.Sprint(t.Value)
}

func (t *FloatValue) GetValueType() string {
	return "float"
}

func NewFloat(v float64) *FloatValue {
	return &FloatValue{Value: v}
}
