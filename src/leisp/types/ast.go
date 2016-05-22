// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import "fmt"

type AST struct {
	Type     string
	Value    ValueType
	Children []*AST
}

func NewValueAST(v ValueType) *AST {
	return &AST{
		Type:     "value",
		Value:    v,
		Children: nil,
	}
}

func NewEOFAST() *AST {
	return &AST{
		Type:     "eof",
		Value:    nil,
		Children: nil,
	}
}

func NewSExpressionAST(ch []*AST) *AST {
	return &AST{
		Type:     "s-expr",
		Value:    nil,
		Children: ch,
	}
}

func NewQExpressionAST(ch []*AST) *AST {
	return &AST{
		Type:     "q-expr",
		Value:    nil,
		Children: ch,
	}
}

func NewArrayAST(ch []*AST) *AST {
	return &AST{
		Type:     "array",
		Value:    nil,
		Children: ch,
	}
}

// Dump prints the AST structs
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

// IsType returns true if it's specified type
func (a *AST) IsType(t string) bool {
	if a.Type == t {
		return true
	}
	return false
}

// IsEOF returns true if this is EOF AST
func (a *AST) IsEOF() bool {
	return a.IsType("eof")
}

// IsSExpression returns true if this is s-expr AST
func (a *AST) IsSExpression() bool {
	return a.IsType("s-expr")
}

// IsQExpression returns true if this is s-expr AST
func (a *AST) IsQExpression() bool {
	return a.IsType("q-expr")
}

// IsArray returns true if this is s-expr AST
func (a *AST) IsArray() bool {
	return a.IsType("array")
}

// IsValue returns true if this is value AST
func (a *AST) IsValue() bool {
	return a.IsType("value")
}
