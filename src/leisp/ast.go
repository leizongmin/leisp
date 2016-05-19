// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package leisp

type AST struct {
	Type     string
	Value    interface{}
	Children []*AST
}

func newAST(t string, v interface{}, ch []*AST) *AST {
	return &AST{
		Type:     t,
		Value:    v,
		Children: ch,
	}
}

func newEmptyAST() *AST {
	return newAST("empty", nil, nil)
}

func newNullAST() *AST {
	return newAST("null", nil, nil)
}

func newBooleanAST() *AST {
	return newAST("boolean", nil, nil)
}

func newCharAST(c string) *AST {
	return newAST("char", c, nil)
}

func newStringAST(s string) *AST {
	return newAST("string", s, nil)
}

func newIntegerAST(v int64) *AST {
	return newAST("integer", v, nil)
}

func newFloatAST(v float64) *AST {
	return newAST("float", v, nil)
}

func newRatioAST(s string) *AST {
	return newAST("ratio", s, nil)
}

func newSymbolAST(s string) *AST {
	return newAST("symbol", s, nil)
}

func newKeywordAST(s string) *AST {
	return newAST("keyword", s, nil)
}

func newListAST(ch []*AST) *AST {
	return newAST("list", nil, ch)
}

func newSExpressionAST(ch []*AST) *AST {
	return newAST("s-expression", nil, ch)
}

func newQExpressionAST(ch []*AST) *AST {
	return newAST("q-expression", nil, ch)
}
