// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import "fmt"

type BooleanValue struct {
	Value bool
}

func (v *BooleanValue) ToString() string {
	if v.Value {
		return "true"
	}
	return "false"
}

func (v *BooleanValue) GetType() string {
	return "boolean"
}

func (v *BooleanValue) IsValue() bool {
	return true
}

func (v *BooleanValue) To(t string) (ValueType, error) {
	return nil, fmt.Errorf("cannot convert boolean to %s: does not implement yet", t)
}

func (v *BooleanValue) EqualTo(t ValueType) bool {
	if v2, ok := t.(*BooleanValue); ok {
		if v2.Value == v.Value {
			return true
		}
	}
	return false
}

func NewBooleanValue(v bool) *BooleanValue {
	return &BooleanValue{Value: v}
}
