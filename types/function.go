// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import "fmt"

type BuiltinFunction func(s *Scope, args []*AST) *Atom

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

func (v *FunctionValue) IsValue() bool {
	return false
}

func (v *FunctionValue) ConvertTo(t string) (ValueType, error) {
	return nil, fmt.Errorf("cannot convert function to %s: does not implement yet", t)
}

func (v *FunctionValue) EqualTo(t ValueType) bool {
	if v2, ok := t.(*FunctionValue); ok {
		if &v2.Value == &v.Value {
			return true
		}
	}
	return false
}

func (t *FunctionValue) Call(s *Scope, args []*AST) *Atom {
	return t.Value(s, args)
}

func NewFunctionValue(n string, v BuiltinFunction) *FunctionValue {
	return &FunctionValue{Name: n, Value: v}
}
