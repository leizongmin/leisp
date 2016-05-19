// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package leisp

import (
	"fmt"
	"strings"
)

func Eval(prog string) *Result {

	p := NewParser(strings.NewReader(prog))
	var r *Result

	for {
		if ast, err := p.Parse(); err != nil {
			pos := p.GetPosition()
			return newErrorResult(fmt.Errorf("Error: %s at line %d,%d\n", err.Error(), pos.Line, pos.Column))
		} else if ast.IsEmpty() {
			break
		} else {
			r = EvalAST(ast)
		}
	}

	return r
}

func EvalAST(ast *AST) *Result {

	switch ast.Type {
	case "ratio", "integer", "float", "string", "char", "null", "boolean", "symbol":
		return newResult(ast.Value)
	case "list":
		return newResult(ast.Children)
	case "q-expr":
		return newResult(ast.Children)
	case "s-expr":
		if len(ast.Children) < 1 {
			return newEmptyResult()
		}
		op := EvalAST(ast.Children[0])
		args := ast.Children[1:]
		values := make([]*Result, len(args))
		for i, v := range args {
			values[i] = EvalAST(v)
		}
		fn, ok := op.Value.(string)
		if !ok {
			return newErrorResult(fmt.Errorf("%s is not a function", op.ToString()))
		}
		return callBuiltinFunction(fn, values)
	}

	return newEmptyResult()
}
