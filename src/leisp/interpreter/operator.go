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
	switch op := args[0].Value.(type) {
	case types.SymbolValue:
		return types.NewAtom(types.NewString("ok"))
	case types.KeywordValue:
		return types.NewErrorMessageAtom("keyword s-expression is not implementation")
	default:
		return types.NewErrorAtom(fmt.Errorf("invalid s-expression, operator cannot be type %s", op.Value.GetValueType()))
	}
}
