// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package interpreter

import (
	"fmt"
	"leisp/types"
	"strings"
)

func builtinPrint(s *types.Scope, list []*types.AST) *types.Atom {

	args, errAtom := astListToAtomList(s, list)
	if errAtom != nil {
		return errAtom
	}

	if len(args) > 0 {
		list := make([]string, len(args))
		for i, a := range args {
			if a.IsValue() {
				if sym, ok := a.Value.(*types.SymbolValue); ok {
					if v, err := sym.GetFinalValue(s); err != nil {
						return types.NewErrorAtom(err)
					} else {
						list[i] = v.ToString()
					}
					continue
				}
			}
			list[i] = a.ToString()
		}
		fmt.Print(strings.Join(list, " "))
	}
	return types.NewAtom(types.NewNullValue())
}

func builtinPrintln(s *types.Scope, list []*types.AST) *types.Atom {

	a := builtinPrint(s, list)
	fmt.Println("")
	return a
}

func init() {

	RegisterBuiltinFunction("print", builtinPrint)
	RegisterBuiltinFunction("println", builtinPrintln)

}
