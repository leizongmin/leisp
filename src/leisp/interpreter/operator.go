// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package interpreter

import (
	"fmt"
	"leisp/types"
	"strings"
)

var Scope = types.NewRootScope()

func RegisterBuiltinFunction(name string, handler types.BuiltinFunction) {
	Scope.Declare(name, types.NewFunctionValue(name, handler))
}

func callOperator(s *types.Scope, name string, args []*types.AST) *types.Atom {

	op, err := s.Get(name)
	if err != nil {
		return types.NewErrorAtom(err)
	}

	if fn, ok := op.(*types.FunctionValue); ok {

		return fn.Call(s, args)

	} else if lam, ok := op.(*types.LambdaValue); ok {

		ns := types.NewScope(lam.Value.Scope)

		values, errAtom := astListToAtomList(s, args)
		if errAtom != nil {
			return errAtom
		}

		if len(values) < len(lam.Value.Names) {
			return types.NewErrorAtom(fmt.Errorf("invalid arguments number for %s, expected %d actually %d", name, len(lam.Value.Names), len(values)))
		}
		for i, n := range lam.Value.Names {
			if strings.Index(n, "...") == 0 {
				return types.NewErrorMessageAtom("variable arguments does not implement yet")
			}
			if err := ns.Declare(n, values[i].Value); err != nil {
				return types.NewErrorAtom(err)
			}
		}

		var r *types.Atom
		for _, a := range lam.Value.Body {
			r = EvalAST(ns, a.Value)
			if r.IsError() {
				break
			}
		}

		return r

	}

	return types.NewErrorAtom(fmt.Errorf("invalid s-expression, operator cannot be type %s", op.GetType()))
}
