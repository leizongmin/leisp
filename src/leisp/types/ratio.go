// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import "fmt"

type RatioValue struct {
	Value string
}

func (t *RatioValue) ToString() string {
	return fmt.Sprint(t.Value)
}

func (t *RatioValue) GetValueType() string {
	return "ratio"
}

func NewRatio(v string) *RatioValue {
	return &RatioValue{Value: v}
}
