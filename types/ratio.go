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

func (v *RatioValue) IsValue() bool {
	return true
}

func (v *RatioValue) ConvertTo(t string) (ValueType, error) {
	return nil, fmt.Errorf("cannot convert ratio to %s: does not implement yet", t)
}

func (v *RatioValue) EqualTo(t ValueType) bool {
	if v2, ok := t.(*RatioValue); ok {
		if v2.Value == v.Value {
			return true
		}
	}
	return false
}

func NewRatioValue(v string) *RatioValue {
	return &RatioValue{Value: v}
}
