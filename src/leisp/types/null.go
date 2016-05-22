// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import "fmt"

type NullValue struct {
	Value bool
}

func (t *NullValue) ToString() string {
	return fmt.Sprint(t.Value)
}

func (t *NullValue) GetValueType() string {
	return "null"
}

func NewNull() *NullValue {
	return &NullValue{Value: false}
}
