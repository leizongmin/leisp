// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import "fmt"

type RatioValue struct {
	Value string
}

func (v *RatioValue) ToString() string {
	return fmt.Sprint(v.Value)
}

func (v *RatioValue) GetType() string {
	return "ratio"
}

func (v *RatioValue) To(t string) (ValueType, error) {
	return nil, fmt.Errorf("cannot convert ratio to %s: does not implement yet", t)
}

func NewRatioValue(v string) *RatioValue {
	return &RatioValue{Value: v}
}
