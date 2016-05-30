// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import "fmt"

type NullValue struct {
	Value bool
}

func (v *NullValue) ToString() string {
	return "null"
}

func (v *NullValue) GetType() string {
	return "null"
}

func (v *NullValue) IsValue() bool {
	return true
}

func (v *NullValue) To(t string) (ValueType, error) {
	return nil, fmt.Errorf("cannot convert null to %s: does not implement yet", t)
}

func (v *NullValue) EqualTo(t ValueType) bool {
	if v2, ok := t.(*NullValue); ok {
		if v2.Value == v.Value {
			return true
		}
	}
	return false
}

func NewNullValue() *NullValue {
	return &NullValue{Value: false}
}
