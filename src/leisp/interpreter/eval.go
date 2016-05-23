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

func Eval(prog string) *types.Atom {

	p := parser.NewParser(strings.NewReader(prog))
	var r *types.Atom

	for {
		if ast, err := p.Parse(); err != nil {
			pos := p.GetPosition()
			return types.NewErrorAtom(fmt.Errorf("Error: %s at line %d,%d\n", err.Error(), pos.Line, pos.Column-1))
		} else if ast.IsEOF() {
			break
		} else {
			r = EvalAST(Scope, ast)
			if r.Error != nil {
				break
			}
		}
	}

	return r
}

func EvalAST(s *types.Scope, ast *types.AST) *types.Atom {

	if ast.IsValue() {
		return types.NewAtom(ast.Value)
	}

	if ast.IsSExpression() {
		return CallBuiltinFunction(s, astListToAtomList(s, ast.Children))
	}

	if ast.IsList() {
		list := astListToAtomList(s, ast.Children)
		ret := make([]types.ValueType, len(list))
		for i, a := range list {
			if a.IsError() {
				return a
			}
			ret[i] = a.Value
		}
		return types.NewAtom(types.NewListValue(ret))
	}

	if ast.IsQExpression() {
		return types.NewAtom(types.NewExpressionValue(types.NewSExpressionAST(ast.Children)))
	}

	if ast.IsQuote() {
		return types.NewAtom(ast.Value)
	}

	return types.NewAtom(types.NewNullValue())
}

func astListToAtomList(s *types.Scope, list []*types.AST) []*types.Atom {
	ret := make([]*types.Atom, len(list))
	for i, a := range list {
		ret[i] = EvalAST(s, a)
	}
	return ret
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
