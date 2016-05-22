// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import "fmt"

type BooleanValue struct {
	Value bool
}

func (t *BooleanValue) ToString() string {
	return fmt.Sprint(t.Value)
}

func (t *BooleanValue) GetValueType() string {
	return "boolean"
}

func NewBoolean(v bool) *BooleanValue {
	return &BooleanValue{Value: v}
}
