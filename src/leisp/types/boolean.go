// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

type BooleanValue struct {
	Value bool
}

func (t *BooleanValue) ToString() string {
	if t.Value {
		return "T"
	}
	return "NIL"
}

func (t *BooleanValue) GetType() string {
	return "boolean"
}

func NewBoolean(v bool) *BooleanValue {
	return &BooleanValue{Value: v}
}
