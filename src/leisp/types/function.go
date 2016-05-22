// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import "fmt"

type Function func(s *Scope, args []*Atom) *Atom

type FunctionValue struct {
	Value Function
}

func (t *FunctionValue) ToString() string {
	return fmt.Sprint(t.Value)
}

func (t *FunctionValue) GetType() string {
	return "function"
}

func (t *FunctionValue) Call(s *Scope, args []*Atom) *Atom {
	return t.Value(s, args)
}

func NewFunction(v Function) *FunctionValue {
	return &FunctionValue{Value: v}
}
