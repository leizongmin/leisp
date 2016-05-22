// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

type NullValue struct {
	Value bool
}

func (t *NullValue) ToString() string {
	return "NIL"
}

func (t *NullValue) GetType() string {
	return "null"
}

func NewNull() *NullValue {
	return &NullValue{Value: false}
}
