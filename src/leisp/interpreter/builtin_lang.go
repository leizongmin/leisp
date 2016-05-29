// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package interpreter

import (
	"fmt"
	"leisp/types"
	"os"
)

func builtinTypeOf(s *types.Scope, list []*types.AST) *types.Atom {

	args, errAtom := astListToAtomList(s, list)
	if errAtom != nil {
		return errAtom
	}

	if len(args) != 1 {
		return types.NewErrorMessageAtom(`wrong arguments number for type-of`)
	}
	a := args[0]
	if a.IsError() {
		return types.NewAtom(types.NewStringValue("error"))
	}
	return types.NewAtom(types.NewStringValue(a.Value.GetType()))
}

func builtinDefvar(s *types.Scope, list []*types.AST) *types.Atom {

	if len(list) != 2 {
		return types.NewErrorMessageAtom(`wrong arguments number for defvar`)
	}

	first := list[0]
	if !first.IsValue() {
		return types.NewErrorAtom(fmt.Errorf("function name must be symbol: %s", first.ToString()))
	}
	name, ok := first.Value.(*types.SymbolValue)
	if !ok {
		return types.NewErrorAtom(fmt.Errorf("function name must be symbol: %s", first.ToString()))
	}

	value := EvalAST(s, list[1])
	if value.IsError() {
		return value
	}

	if err := s.Declare(name.Value, value.Value); err != nil {
		return types.NewErrorAtom(err)
	}

	return value
}

func builtinNewScope(s *types.Scope, list []*types.AST) *types.Atom {

	args, errAtom := astListToAtomList(s, list)
	if errAtom != nil {
		return errAtom
	}

	argc := len(args)
	if argc > 1 {
		return types.NewErrorMessageAtom(`wrong arguments number for new-scope`)
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
		return types.NewErrorMessageAtom("wrong arguments number for lambda")
	}

	first := list[0]
	if !first.IsArray() {
		return types.NewErrorAtom(fmt.Errorf("lambda arguments must be a list: %s", first.ToString()))
	}

	var names []string
	var varNames []string
	beginRestName := false
	for _, v := range first.Children {
		if !v.IsValue() {
			return types.NewErrorAtom(fmt.Errorf("lambda argument must be symbol: %s", v.ToString()))
		}
		if n, ok := v.Value.(*types.SymbolValue); ok {
			if beginRestName {
				varNames = append(varNames, n.Value)
			} else {
				if n.Value == "&" {
					beginRestName = true
				} else {
					names = append(names, n.Value)
				}
			}
		} else {
			return types.NewErrorAtom(fmt.Errorf("invalid arguments type: %s", v.ToString()))
		}
	}

	restNameCount := len(varNames)
	if restNameCount > 1 {
		return types.NewErrorAtom(fmt.Errorf("only support 1 variable arugment, actually got %d", restNameCount))
	}

	body := make([]*types.ExpressionValue, argc-1)
	for i, v := range list[1:] {
		body[i] = types.NewExpressionValue(v)
	}
	lam := types.NewLambdaValue(types.NewLambdaValueInfo(s, names, varNames, body, "no source"))

	return types.NewAtom(lam)
}

func builtinDefn(s *types.Scope, list []*types.AST) *types.Atom {

	argc := len(list)
	if argc < 3 {
		return types.NewErrorMessageAtom("wrong arguments number for defn")
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

func builtinEval(s *types.Scope, list []*types.AST) *types.Atom {

	args, errAtom := astListToAtomList(s, list)
	if errAtom != nil {
		return errAtom
	}

	r := types.NewEmptyAtom()

	for _, a := range args {
		if expr, ok := a.Value.(*types.ExpressionValue); ok {
			r = EvalAST(s, expr.Value)
			if r.IsError() {
				return r
			}
			continue
		}
		r = a
	}

	return r
}

func builtinExit(s *types.Scope, list []*types.AST) *types.Atom {

	args, errAtom := astListToAtomList(s, list)
	if errAtom != nil {
		return errAtom
	}

	exitCode := 0
	if len(args) > 0 {
		first := args[0]
		code, ok := first.Value.(*types.IntegerValue)
		if !ok {
			return types.NewErrorAtom(fmt.Errorf("exit code must be type integer: actually type is %s", first.Value.GetType()))
		}
		exitCode = int(code.Value)
	}
	os.Exit(exitCode)

	return types.NewEmptyAtom()
}

func builtinValue(s *types.Scope, list []*types.AST) *types.Atom {

	if len(list) != 1 {
		return types.NewErrorMessageAtom("wrong arguments number for value")
	}

	first := list[0]
	if first.IsValue() {
		if sym, ok := first.Value.(*types.SymbolValue); ok {
			val, err := s.Get(sym.Value)
			if err != nil {
				return types.NewErrorAtom(err)
			}
			return types.NewAtom(val)
		}
		return types.NewAtom(first.Value)
	}

	return types.NewAtom(types.NewExpressionValue(first))
}

func builtinValueByName(s *types.Scope, list []*types.AST) *types.Atom {

	if len(list) != 1 {
		return types.NewErrorMessageAtom("wrong arguments number for value")
	}

	args, errAtom := astListToAtomList(s, list)
	if errAtom != nil {
		return errAtom
	}

	first := args[0]
	str, ok := first.Value.(*types.StringValue)
	if !ok {
		return types.NewErrorAtom(fmt.Errorf("invalid arguments type for symbol-value, expected string actually %s", first.Value.GetType()))
	}

	val, err := s.Get(str.Value)
	if err != nil {
		return types.NewErrorAtom(err)
	}

	return types.NewAtom(val)
}

func builtinEqual(s *types.Scope, list []*types.AST) *types.Atom {

	args, errAtom := astListToAtomList(s, list)
	if errAtom != nil {
		return errAtom
	}

	if len(args) < 2 {
		return types.NewErrorMessageAtom("wrong arguments number for value")
	}

	ok := true
	for i, right := range args[1:] {
		left := args[i]
		if left.Value.GetType() != right.Value.GetType() {
			ok = false
			break
		}
		ok = left.Value.EqualTo(right.Value)
		if !ok {
			break
		}
	}

	return types.NewAtom(types.NewBooleanValue(ok))
}

func builtinAnd(s *types.Scope, list []*types.AST) *types.Atom {

	if len(list) < 2 {
		return types.NewErrorMessageAtom("wrong arguments number for and")
	}

	for _, a := range list {

		r := EvalAST(s, a)
		if r.IsError() {
			return r
		}

		if _, ok := r.Value.(*types.NullValue); ok {
			return types.NewAtom(types.NewBooleanValue(false))
		}
		if boolValue, ok := r.Value.(*types.BooleanValue); ok {
			if !boolValue.Value {
				return types.NewAtom(types.NewBooleanValue(false))
			}
		}
	}

	return types.NewAtom(types.NewBooleanValue(true))
}

func builtinOr(s *types.Scope, list []*types.AST) *types.Atom {

	if len(list) < 2 {
		return types.NewErrorMessageAtom("wrong arguments number for or")
	}

	for _, a := range list {

		r := EvalAST(s, a)
		if r.IsError() {
			return r
		}

		if _, ok := r.Value.(*types.NullValue); ok {
			continue
		}
		if boolValue, ok := r.Value.(*types.BooleanValue); ok {
			if !boolValue.Value {
				continue
			}
		}

		return types.NewAtom(types.NewBooleanValue(true))
	}

	return types.NewAtom(types.NewBooleanValue(false))
}

func builtinNot(s *types.Scope, list []*types.AST) *types.Atom {

	if len(list) != 1 {
		return types.NewErrorMessageAtom("wrong arguments number for not")
	}

	r := EvalAST(s, list[0])
	if r.IsError() {
		return r
	}

	if _, ok := r.Value.(*types.NullValue); ok {
		return types.NewAtom(types.NewBooleanValue(true))
	}
	if boolValue, ok := r.Value.(*types.BooleanValue); ok {
		if !boolValue.Value {
			return types.NewAtom(types.NewBooleanValue(true))
		}
	}

	return types.NewAtom(types.NewBooleanValue(false))

}

func builtinToString(s *types.Scope, list []*types.AST) *types.Atom {

	if len(list) != 1 {
		return types.NewErrorMessageAtom("wrong arguments number for to-string")
	}

	r := EvalAST(s, list[0])
	if r.IsError() {
		return r
	}

	if strValue, ok := r.Value.(*types.StringValue); ok {
		return types.NewAtom(strValue)
	}

	return types.NewAtom(types.NewStringValue(r.ToString()))

}
