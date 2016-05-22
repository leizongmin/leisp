// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package interpreter

import (
	"fmt"
	"leisp/types"
)

func builtinTypeOf(s *types.Scope, args []*types.Atom) *types.Atom {
	if len(args) != 1 {
		return types.NewErrorMessageAtom(`invalid arguments number for type-of`)
	}
	a := args[0]
	if a.IsError() {
		return types.NewAtom(types.NewString("error"))
	}
	return types.NewAtom(types.NewString(a.Value.GetType()))
}

func builtinDef(s *types.Scope, args []*types.Atom) *types.Atom {
	if len(args) != 2 {
		return types.NewErrorMessageAtom(`invalid arguments number for def`)
	}
	n := args[0]
	v := args[1]
	if !n.IsValue() {
		return n
	}
	if !v.IsValue() {
		return v
	}
	sym, ok := n.Value.(*types.SymbolValue)
	if !ok {
		return types.NewErrorAtom(fmt.Errorf("invalid type of variable name: %s", n.ToString()))
	}

	if err := s.Declare(sym.ToString(), v.Value); err != nil {
		return types.NewErrorAtom(err)
	}

	return v
}

func init() {

	RegisterBuiltinFunction("type-of", builtinTypeOf)
	RegisterBuiltinFunction("def", builtinDef)

}
