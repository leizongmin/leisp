// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import "fmt"

type ArrayValue struct {
	Value []*ValueType
}

func (t *ArrayValue) ToString() string {
	return fmt.Sprint(t.Value)
}

func (t *ArrayValue) GetValueType() string {
	return "array"
}

func NewArray(v []*ValueType) *ArrayValue {
	return &ArrayValue{Value: v}
}
