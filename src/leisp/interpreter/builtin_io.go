// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package interpreter

import (
	"fmt"
	"leisp/types"
	"strings"
)

func builtinPrint(s *types.Scope, args []*types.Atom) *types.Atom {
	if len(args) > 0 {
		list := make([]string, len(args))
		for i, a := range args {
			list[i] = a.ToString()
		}
		fmt.Print(strings.Join(list, " "))
	}
	return types.NewAtom(types.NewNull())
}

func builtinPrintln(s *types.Scope, args []*types.Atom) *types.Atom {
	a := builtinPrint(s, args)
	fmt.Println("")
	return a
}

func init() {

	RegisterBuiltinFunction("print", builtinPrint)
	RegisterBuiltinFunction("println", builtinPrintln)

}
