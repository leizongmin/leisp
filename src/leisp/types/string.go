// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

type StringValue struct {
	Value string
}

func (t *StringValue) ToString() string {
	return "\"" + t.Value + "\""
}

func (t *StringValue) GetType() string {
	return "string"
}

func NewString(v string) *StringValue {
	return &StringValue{Value: v}
}
