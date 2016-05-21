// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import "fmt"

type IntType struct {
	Value int64
}

func (t *IntType) ToString() string {
	return fmt.Sprint(t.Value)
}

func (t *IntType) GetValueType() string {
	return "int"
}

func NewInt(v int64) *IntType {
	return &IntType{Value: v}
}
