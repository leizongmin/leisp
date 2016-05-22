// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package interpreter

import (
	"fmt"
	"leisp/types"
)

var Scope = types.NewRootScope()

func RegisterBuiltinFunction(name string, handler types.Function) {
	Scope.Declare(name, types.NewFunction(handler))
}

func CallFunction(s *types.Scope, args []*types.Atom) *types.Atom {
	l := len(args)
	if l < 1 {
		return types.NewErrorMessageAtom("invalid s-expression")
	}
	op := args[0]
	switch op.Value.GetType() {

	case "symbol":
		return callFunction(s, op, args[1:])

	case "keyword":
		return types.NewErrorMessageAtom("keyword s-expression is not implementation")

	default:
		return types.NewErrorAtom(fmt.Errorf("invalid s-expression, operator cannot be type %s", op))
	}
}

func callFunction(s *types.Scope, op *types.Atom, args []*types.Atom) *types.Atom {
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
