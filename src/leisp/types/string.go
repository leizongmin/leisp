// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import "fmt"

type StringValue struct {
	Value string
}

func (t *StringValue) ToString() string {
	return fmt.Sprint(t.Value)
}

func (t *StringValue) GetType() string {
	return "string"
}

func NewString(v string) *StringValue {
	return &StringValue{Value: v}
}
