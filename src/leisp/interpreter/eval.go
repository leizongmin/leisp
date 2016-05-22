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

	return types.NewAtom(types.NewNull())

	// switch ast.Type {

	// case "ratio", "integer", "float", "string", "char", "null", "boolean", "symbol", "keyword":
	// 	return newAtom(ast.Value)

	// case "list":
	// 	return newAtom(ast.Children)

	// case "q-expr":
	// 	return newAtom(newSExpressionAST(ast.Children))

	// case "s-expr":
	// 	if len(ast.Children) < 1 {
	// 		return newEmptyAtom()
	// 	}
	// 	op := EvalAST(ast.Children[0])
	// 	args := ast.Children[1:]
	// 	values := make([]*types.Atom, len(args))
	// 	for i, v := range args {
	// 		values[i] = EvalAST(v)
	// 	}
	// 	fn, ok := op.Value.(string)
	// 	if !ok {
	// 		return newErrorAtom(fmt.Errorf("%s is not a function", op.ToString()))
	// 	}
	// 	// return builtin.Call(fn, values)
	// }

	// return newEmptyAtom()
}

func astListToAtomList(s *types.Scope, ast []*types.AST) []*types.Atom {
	list := make([]*types.Atom, len(ast))
	for i, a := range ast {
		list[i] = EvalAST(s, a)
	}
	return list
}
