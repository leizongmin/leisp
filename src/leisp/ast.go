// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package leisp

type AST struct {
	Type     string
	Value    interface{}
	Children []*AST
}

func NewAST(t string, v interface{}, ch []*AST) *AST {
	return &AST{
		Type:     t,
		Value:    v,
		Children: ch,
	}
}

func NewEmptyAST() *AST {
	return NewAST("empty", nil, nil)
}

func NewNullAST() *AST {
	return NewAST("null", nil, nil)
}

func NewBooleanAST() *AST {
	return NewAST("boolean", nil, nil)
}

func NewCharAST(c string) *AST {
	return NewAST("char", c, nil)
}

func NewStringAST(s string) *AST {
	return NewAST("string", s, nil)
}

func NewIntegerAST(v int64) *AST {
	return NewAST("integer", v, nil)
}

func NewFloatAST(v float64) *AST {
	return NewAST("float", v, nil)
}

func NewRatioAST(s string) *AST {
	return NewAST("ratio", s, nil)
}

func NewSymbolAST(s string) *AST {
	return NewAST("symbol", s, nil)
}

func NewKeywordAST(s string) *AST {
	return NewAST("keyword", s, nil)
}

func NewListAST(ch []*AST) *AST {
	return NewAST("list", nil, ch)
}

func NewSExpressionAST(ch []*AST) *AST {
	return NewAST("s-expression", nil, ch)
}

func NewQExpressionAST(ch []*AST) *AST {
	return NewAST("q-expression", nil, ch)
}
