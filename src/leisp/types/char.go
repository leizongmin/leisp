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

func (t *CharValue) GetType() string {
	return "char"
}

func getRuneByIndex(s string, i int) rune {
	r := []rune(s)
	return r[i]
}

func NewChar(v string) *CharValue {
	return &CharValue{Value: getRuneByIndex(v, 0)}
}
