// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import "fmt"

type ListValue struct {
	Value []*ValueType
}

func (t *ListValue) ToString() string {
	return fmt.Sprint(t.Value)
}

func (t *ListValue) GetType() string {
	return "list"
}

func NewList(v []*ValueType) *ListValue {
	return &ListValue{Value: v}
}
