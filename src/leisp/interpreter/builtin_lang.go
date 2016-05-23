// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package interpreter

import (
	"fmt"
	"leisp/types"
)

func builtinTypeOf(s *types.Scope, list []*types.AST) *types.Atom {

	args, errAtom := astListToAtomList(s, list)
	if errAtom != nil {
		return errAtom
	}

	if len(args) != 1 {
		return types.NewErrorMessageAtom(`invalid arguments number for type-of`)
	}
	a := args[0]
	if a.IsError() {
		return types.NewAtom(types.NewStringValue("error"))
	}
	return types.NewAtom(types.NewStringValue(a.Value.GetType()))
}

func builtinDefvar(s *types.Scope, list []*types.AST) *types.Atom {

	args, errAtom := astListToAtomList(s, list)
	if errAtom != nil {
		return errAtom
	}

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

func builtinNewScope(s *types.Scope, list []*types.AST) *types.Atom {

	args, errAtom := astListToAtomList(s, list)
	if errAtom != nil {
		return errAtom
	}

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

func builtinLambda(s *types.Scope, list []*types.AST) *types.Atom {

	argc := len(list)
	if argc < 2 {
		return types.NewErrorMessageAtom("invalid arguments number for lambda")
	}

	first := list[0]
	if !first.IsList() {
		return types.NewErrorAtom(fmt.Errorf("lambda arguments must be a list: %s", first.ToString()))
	}

	names := make([]string, len(first.Children))
	for i, v := range first.Children {
		if !v.IsValue() {
			return types.NewErrorAtom(fmt.Errorf("lambda argument must be symbol: %s", v.ToString()))
		}
		if n, ok := v.Value.(*types.SymbolValue); ok {
			names[i] = n.Value
		} else {
			return types.NewErrorAtom(fmt.Errorf("invalid arguments type: %s", v.ToString()))
		}
	}

	body := make([]*types.ExpressionValue, argc-1)
	for i, v := range list[1:] {
		body[i] = types.NewExpressionValue(v)
	}
	lam := types.NewLambdaValue(types.NewLambdaValueInfo(s, names, body, "no source"))

	return types.NewAtom(lam)
}

func builtinDefn(s *types.Scope, list []*types.AST) *types.Atom {

	argc := len(list)
	if argc < 3 {
		return types.NewErrorMessageAtom("invalid arguments number for defn")
	}

	first := list[0]
	if !first.IsValue() {
		return types.NewErrorAtom(fmt.Errorf("function name must be symbol: %s", first.ToString()))
	}
	name, ok := first.Value.(*types.SymbolValue)
	if !ok {
		return types.NewErrorAtom(fmt.Errorf("function name must be symbol: %s", name.ToString()))
	}

	lam := builtinLambda(s, list[1:])
	if lam.IsError() {
		return lam
	}

	if err := s.Declare(name.Value, lam.Value); err != nil {
		return types.NewErrorAtom(err)
	}

	return lam
}

func init() {

	RegisterBuiltinFunction("type-of", builtinTypeOf)
	RegisterBuiltinFunction("defvar", builtinDefvar)
	RegisterBuiltinFunction("lambda", builtinLambda)
	RegisterBuiltinFunction("defn", builtinDefn)
	RegisterBuiltinFunction("new-scope", builtinNewScope)

}
