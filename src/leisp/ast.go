// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package leisp

type AST struct {
	Type     string
	Value    interface{}
	Children []AST
}

func NewAST(t string, v interface{}) *AST {
	return &AST{
		Type:  t,
		Value: v,
	}
}

func NewEmptyAST() *AST {
	return NewAST("empty", nil)
}

func NewNullAST() *AST {
	return NewAST("null", nil)
}

func NewBooleanAST() *AST {
	return NewAST("boolean", nil)
}

func NewCharAST(c string) *AST {
	return NewAST("char", c)
}

func NewStringAST(s string) *AST {
	return NewAST("string", s)
}

func NewIntegerAST(v int64) *AST {
	return NewAST("integer", v)
}

func NewFloatAST(v float64) *AST {
	return NewAST("float", v)
}

func NewRatioAST(s string) *AST {
	return NewAST("ratio", s)
}

func NewSymbolAST(s string) *AST {
	return NewAST("symbol", s)
}

func NewKeywordAST(s string) *AST {
	return NewAST("keyword", s)
}

func NewListAST() *AST {
	return NewAST("list", nil)
}

func NewSExpressionAST() *AST {
	return NewAST("s-expression", nil)
}

func NewQExpressionAST() *AST {
	return NewAST("q-expression", nil)
}
