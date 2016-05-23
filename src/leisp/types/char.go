// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import "fmt"

type CharValue struct {
	Value rune
}

func (v *CharValue) ToString() string {
	return fmt.Sprint(v.Value)
}

func (v *CharValue) GetType() string {
	return "char"
}

func (v *CharValue) To(t string) (ValueType, error) {
	return nil, fmt.Errorf("cannot convert char to %s: does not implement yet", t)
}

func getRuneByIndex(s string, i int) rune {
	r := []rune(s)
	return r[i]
}

func NewChar(v string) *CharValue {
	return &CharValue{Value: getRuneByIndex(v, 0)}
}
