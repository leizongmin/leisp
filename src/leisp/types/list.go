// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import (
	"fmt"
	"strings"
)

type ListValue struct {
	Value []ValueType
}

func (v *ListValue) ToString() string {
	list := make([]string, len(v.Value))
	for i, v2 := range v.Value {
		list[i] = v2.ToString()
	}
	return "[" + strings.Join(list, " ") + "]"
}

func (v *ListValue) GetType() string {
	return "list"
}

func (v *ListValue) To(t string) (ValueType, error) {
	return nil, fmt.Errorf("cannot convert list to %s: does not implement yet", t)
}

func NewListValue(v []ValueType) *ListValue {
	return &ListValue{Value: v}
}
