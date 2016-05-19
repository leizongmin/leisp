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
	case "integer", "float", "string", "char":
		return newResult(ast.Value)
	}
	return newResult(2)
}
