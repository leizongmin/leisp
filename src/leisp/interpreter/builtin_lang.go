// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package interpreter

import (
	"fmt"
	"leisp/types"
)

func builtinTypeOf(s *types.Scope, args []*types.Atom) *types.Atom {

	if len(args) != 1 {
		return types.NewErrorMessageAtom(`invalid arguments number for type-of`)
	}
	a := args[0]
	if a.IsError() {
		return types.NewAtom(types.NewStringValue("error"))
	}
	return types.NewAtom(types.NewStringValue(a.Value.GetType()))
}

func builtinDef(s *types.Scope, args []*types.Atom) *types.Atom {

	if len(args) != 2 {
		return types.NewErrorMessageAtom(`invalid arguments number for def`)
	}
	n := args[0]
	v := args[1]
	if !n.IsValue() {
		return n
	}
	if !v.IsValue() {
		return v
	}
	sym, ok := n.Value.(*types.SymbolValue)
	if !ok {
		return types.NewErrorAtom(fmt.Errorf("invalid type of variable name: %s", n.ToString()))
	}

	if err := s.Declare(sym.ToString(), v.Value); err != nil {
		return types.NewErrorAtom(err)
	}

	return v
}

func builtinNewScope(s *types.Scope, args []*types.Atom) *types.Atom {

	argc := len(args)
	if argc > 1 {
		return types.NewErrorMessageAtom(`invalid arguments number for new-scope`)
	}
	if argc == 0 {
		return types.NewAtom(types.NewScopeValue(types.NewScope(nil)))
	}
	a, err := getAtomFinalValue(s, args[0])
	if err != nil {
		return types.NewErrorAtom(err)
	}
	if s2, ok := a.(*types.ScopeValue); ok {
		return types.NewAtom(types.NewScopeValue(types.NewScope(s2.Value)))
	}
	return types.NewErrorAtom(fmt.Errorf("%s is not a scope: %s", a.GetType(), a.ToString()))
}

func builtinLambda(s *types.Scope, args []*types.Atom) *types.Atom {

	argc := len(args)
	if argc < 2 {
		return types.NewErrorMessageAtom(`invalid arguments number for lambda`)
	}

	first := args[0]
	if first.IsError() {
		return first
	}
	arg, ok := first.Value.(*types.ListValue)
	if !ok {
		return types.NewErrorAtom(fmt.Errorf("%s is not a list: %s", first.Value.GetType(), first.Value.ToString()))
	}
	names := make([]string, len(arg.Value))
	for i, v := range arg.Value {
		if n, ok := v.(*types.SymbolValue); ok {
			names[i] = n.Value
		} else {
			return types.NewErrorAtom(fmt.Errorf("invalid arguments type: %s", v.ToString()))
		}
	}

	fmt.Println(*args[1], *args[2])

	return types.NewAtom(types.NewStringValue("haha"))
}

func init() {

	RegisterBuiltinFunction("type-of", builtinTypeOf)
	RegisterBuiltinFunction("def", builtinDef)
	RegisterBuiltinFunction("lambda", builtinLambda)
	RegisterBuiltinFunction("new-scope", builtinNewScope)

}
