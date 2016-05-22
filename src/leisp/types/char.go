// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import "fmt"

type CharValue struct {
	Value rune
}

func (t *CharValue) ToString() string {
	return fmt.Sprint(t.Value)
}

func (t *CharValue) GetValueType() string {
	return "char"
}

func NewChar(v string) *CharValue {
	return &CharValue{Value: v}
}
