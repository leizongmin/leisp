// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package leisp

import "fmt"

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
	return newAST("boolean", true, nil)
}

func newCharAST(s string) *AST {
	return newAST("char", []rune(s)[0], nil)
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
	return newAST("s-expr", nil, ch)
}

func newQExpressionAST(ch []*AST) *AST {
	return newAST("q-expr", nil, ch)
}

func (a *AST) Dump() {
	a.dump(0)
}

func makeIndentString(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += "  "
	}
	return s
}

func (a *AST) dump(indent int) {

	prefix := makeIndentString(indent) + "--"

	val := ""
	if a.Value != nil {
		val = fmt.Sprint(a.Value)
	}

	fmt.Printf("%s %s:\t%s\n", prefix, a.Type, val)
	if len(a.Children) > 0 {
		for _, c := range a.Children {
			c.dump(indent + 1)
		}
	} else if a.Type == "list" || a.Type == "q-expr" || a.Type == "s-expr" {
		fmt.Printf("%s-- empty %s\n", makeIndentString(indent+1), a.Type)
	}
}

func (a *AST) IsEmpty() bool {
	if a.Type == "empty" {
		return true
	}
	return false
}
