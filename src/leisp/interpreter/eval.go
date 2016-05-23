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
			return types.NewErrorAtom(fmt.Errorf("Error: %s at line %d,%d\n", err.Error(), pos.Line, pos.Column))
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
		return CallFunction(s, astListToAtomList(s, ast.Children))
	}

	if ast.IsList() {
		return types.NewErrorMessageAtom("does not implement list AST yet")
	}

	if ast.IsQExpression() {
		return types.NewAtom(types.NewExpression(types.NewSExpressionAST(ast.Children)))
	}

	return types.NewAtom(types.NewNull())
}

func astListToAtomList(s *types.Scope, ast []*types.AST) []*types.Atom {
	list := make([]*types.Atom, len(ast))
	for i, a := range ast {
		list[i] = EvalAST(s, a)
	}
	return list
}
