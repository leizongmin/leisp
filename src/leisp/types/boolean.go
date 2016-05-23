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
		return "T"
	}
	return "NIL"
}

func (v *BooleanValue) GetType() string {
	return "boolean"
}

func (v *BooleanValue) To(t string) (ValueType, error) {
	return nil, fmt.Errorf("cannot convert boolean to %s: does not implement yet", t)
}

func NewBoolean(v bool) *BooleanValue {
	return &BooleanValue{Value: v}
}
