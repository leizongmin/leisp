// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import "fmt"

type Function func(s *Scope, args []*Atom) *Atom

type FunctionValue struct {
	Value Function
	Name  string
}

func (t *FunctionValue) ToString() string {
	n := t.Name
	if n == "" {
		n = "anonymous"
	}
	return fmt.Sprintf("<Function#%s>", n)
}

func (t *FunctionValue) GetType() string {
	return "function"
}

func (t *FunctionValue) Call(s *Scope, args []*Atom) *Atom {
	return t.Value(s, args)
}

func NewFunction(n string, v Function) *FunctionValue {
	return &FunctionValue{Name: n, Value: v}
}
