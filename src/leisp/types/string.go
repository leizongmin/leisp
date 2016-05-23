// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import "fmt"

type StringValue struct {
	Value string
}

func (v *StringValue) ToString() string {
	return "\"" + v.Value + "\""
}

func (v *StringValue) GetType() string {
	return "string"
}

func (v *StringValue) To(t string) (ValueType, error) {
	return nil, fmt.Errorf("cannot convert string to %s: does not implement yet", t)
}

func NewString(v string) *StringValue {
	return &StringValue{Value: v}
}
