// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package interpreter

import (
	"fmt"
	"leisp/parser"
	"leisp/types"
	"strings"
)

func Eval(s *types.Scope, prog string) *types.Atom {

	p := parser.NewParser(strings.NewReader(prog))
	r := types.NewEmptyAtom()

	if s == nil {
		s = Scope
	}

	for {
		if ast, err := p.Parse(); err != nil {
			pos := p.GetPosition()
			return types.NewErrorAtom(fmt.Errorf("Error: %s at line %d,%d\n", err.Error(), pos.Line, pos.Column-1))
		} else if ast.IsEOF() {
			break
		} else {
			r = EvalAST(s, ast)
			if r.IsError() {
				break
			}
		}
	}

	return r
}

func EvalASTList(s *types.Scope, list []*types.AST) *types.Atom {

	var r *types.Atom
	for _, a := range list {
		r = EvalAST(s, a)
		if r.IsError() {
			break
		}
	}

	return r
}

func EvalAST(s *types.Scope, a *types.AST) *types.Atom {

	if a.IsValue() {
		return types.NewAtom(a.Value)
	}

	if a.IsSExpression() {
		return evalSExpression(s, a.Children)
	}

	if a.IsList() {
		atoms, err := astListToAtomList(s, a.Children)
		if err != nil {
			return err
		}
		list := make([]types.ValueType, len(atoms))
		for i, a := range atoms {
			list[i] = a.Value
		}
		return types.NewAtom(types.NewListValue(list))
	}

	if a.IsQExpression() {
		return types.NewAtom(types.NewExpressionValue(types.NewSExpressionAST(a.Children)))
	}

	if a.IsQuote() {
		return types.NewAtom(a.Value)
	}

	return types.NewAtom(types.NewNullValue())
}

func evalSExpression(s *types.Scope, list []*types.AST) *types.Atom {

	if len(list) < 1 {
		return types.NewErrorMessageAtom("invalid s-expression")
	}

	first := list[0]
	if !first.IsValue() {
		return types.NewErrorMessageAtom("invalid s-expression")
	}

	var op string
	if sym, ok := first.Value.(*types.SymbolValue); ok {

		op = sym.Value

	} else if _, ok := first.Value.(*types.KeywordValue); ok {

		return types.NewErrorMessageAtom("keyword s-expression does not implement")

	} else {

		return types.NewErrorMessageAtom("invalid s-expression, operator must be symbol")

	}

	return callOperator(s, op, list[1:])
}

func astListToAtomList(s *types.Scope, list []*types.AST) (ret []*types.Atom, err *types.Atom) {
	ret = make([]*types.Atom, len(list))
	for i, a := range list {
		ret[i] = EvalAST(s, a)
		if ret[i].IsError() {
			return nil, ret[i]
		}
	}
	return ret, nil
}

func getAtomFinalValue(s *types.Scope, a *types.Atom) (types.ValueType, error) {
	if a.IsError() {
		return nil, a.Error
	}
	return getFinalValue(s, a.Value)
}

func getAtomListFinalValues(s *types.Scope, list []*types.Atom) ([]types.ValueType, error) {
	ret := make([]types.ValueType, len(list))
	for i, a := range list {
		if a.IsError() {
			return nil, a.Error
		}
		v, err := getAtomFinalValue(s, a)
		if err != nil {
			return nil, err
		}
		ret[i] = v
	}
	return ret, nil
}

func getFinalValue(s *types.Scope, v types.ValueType) (types.ValueType, error) {
	if s == nil {
		return nil, fmt.Errorf("invalid scope: cannot be nil")
	}
	if v.IsValue() {
		return v, nil
	}
	if sym, ok := v.(*types.SymbolValue); ok {
		v2, err := s.Get(sym.Value)
		if err != nil {
			return nil, err
		}
		return getFinalValue(s, v2)
	}
	return v, nil
}
