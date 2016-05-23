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

func CallBuiltinFunction(s *types.Scope, name string, args []*types.AST) *types.Atom {

	op, err := s.Get(name)
	if err != nil {
		return types.NewErrorAtom(err)
	}
	fn, ok := op.(*types.FunctionValue)
	if !ok {
		return types.NewErrorAtom(fmt.Errorf("invalid s-expression, operator cannot be type %s", op.GetType()))
	}
	return fn.Call(s, args)
}
