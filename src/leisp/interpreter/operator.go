// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package interpreter

import (
	"fmt"
	"leisp/types"
)

var Scope = types.NewRootScope()

func RegisterBuiltinFunction(name string, handler types.BuiltinFunction) {
	Scope.Declare(name, types.NewFunctionValue(name, handler))
}

func CallBuiltinFunction(s *types.Scope, args []*types.Atom) *types.Atom {
	l := len(args)
	if l < 1 {
		return types.NewErrorMessageAtom("invalid s-expression")
	}
	switch args[0].Value.(type) {
	case *types.SymbolValue:
		return callBuiltinFunction(s, args[0], args[1:])
	case *types.KeywordValue:
		return types.NewErrorMessageAtom("does not implement keyword s-expression")
	default:
		return types.NewErrorAtom(fmt.Errorf("invalid s-expression, operator cannot be type %s", args[0].Value.GetType()))
	}
}

func callBuiltinFunction(s *types.Scope, op *types.Atom, args []*types.Atom) *types.Atom {
	symbol, _ := op.Value.(*types.SymbolValue)
	value, err := s.Get(symbol.Value)
	if err != nil {
		return types.NewErrorAtom(err)
	}
	fn, ok := value.(*types.FunctionValue)
	if !ok {
		return types.NewErrorAtom(fmt.Errorf("%s(%s) is not a function", value.ToString(), value.GetType()))
	}
	return fn.Call(s, args)
}
