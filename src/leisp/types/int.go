// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import "fmt"

type IntegerValue struct {
	Value int64
}

func (t *IntegerValue) ToString() string {
	return fmt.Sprint(t.Value)
}

func (t *IntegerValue) GetType() string {
	return "integer"
}

func NewInteger(v int64) *IntegerValue {
	return &IntegerValue{Value: v}
}
