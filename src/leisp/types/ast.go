// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import (
	"fmt"
	"strings"
)

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

func NewListAST(ch []*AST) *AST {
	return &AST{
		Type:     "list",
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

	if a.Value == nil {
		fmt.Printf("%s %s\n", prefix, a.Type)
	} else {
		fmt.Printf("%s %s(%s):  %s\n", prefix, a.Type, a.Value.GetType(), a.Value.ToString())
	}

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

// IsList returns true if this is s-expr AST
func (a *AST) IsList() bool {
	return a.IsType("list")
}

// IsValue returns true if this is value AST
func (a *AST) IsValue() bool {
	return a.IsType("value")
}

func (a *AST) ToString() string {
	if a.IsValue() {
		return a.Value.ToString()
	}
	if a.IsEOF() {
		return "EOF"
	}
	if a.IsList() {
		return "[" + a.ChildrenToString() + "]"
	}
	if a.IsQExpression() {
		return "{" + a.ChildrenToString() + "}"
	}
	if a.IsSExpression() {
		return "(" + a.ChildrenToString() + ")"
	}
	return ""
}

func (a *AST) ChildrenToString() string {
	if len(a.Children) > 0 {
		list := make([]string, len(a.Children))
		for i, c := range a.Children {
			list[i] = c.ToString()
		}
		return strings.Join(list, " ")
	}
	return ""
}
