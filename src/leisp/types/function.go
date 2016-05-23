// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import "fmt"

type BuiltinFunction func(s *Scope, args []*Atom) *Atom

type FunctionValue struct {
	Value BuiltinFunction
	Name  string
}

func (v *FunctionValue) ToString() string {
	n := v.Name
	if n == "" {
		n = "anonymous"
	}
	return fmt.Sprintf("<Function#%s>", n)
}

func (v *FunctionValue) GetType() string {
	return "function"
}

func (v *FunctionValue) To(t string) (ValueType, error) {
	return nil, fmt.Errorf("cannot convert function to %s: does not implement yet", t)
}

func (t *FunctionValue) Call(s *Scope, args []*Atom) *Atom {
	return t.Value(s, args)
}

func NewFunctionValue(n string, v BuiltinFunction) *FunctionValue {
	return &FunctionValue{Name: n, Value: v}
}
