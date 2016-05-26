// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import "fmt"

type AnyValue struct {
	Value interface{}
}

func (v *AnyValue) ToString() string {
	return fmt.Sprint(v.Value)
}

func (v *AnyValue) GetType() string {
	return "any"
}

func (v *AnyValue) IsValue() bool {
	return true
}

func (v *AnyValue) To(t string) (ValueType, error) {
	return nil, fmt.Errorf("cannot convert any to %s: does not implement yet", t)
}

func NewAnyValue(v interface{}) *AnyValue {
	return &AnyValue{Value: v}
}
