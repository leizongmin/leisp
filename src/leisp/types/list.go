// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import "strings"

type ListValue struct {
	Value []ValueType
}

func (t *ListValue) ToString() string {
	list := make([]string, len(t.Value))
	for i, v := range t.Value {
		list[i] = v.ToString()
	}
	return "[" + strings.Join(list, " ") + "]"
}

func (t *ListValue) GetType() string {
	return "list"
}

func NewList(v []ValueType) *ListValue {
	return &ListValue{Value: v}
}
