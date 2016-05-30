// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package interpreter

import (
	"fmt"
	"leisp/types"
)

func builtinStdOutPrint(s *types.Scope, list []*types.AST) *types.Atom {

	if len(list) != 1 {
		return types.NewErrorMessageAtom("wrong arguments number for to-string")
	}

	args, errAtom := astListToAtomList(s, list)
	if errAtom != nil {
		return errAtom
	}

	a := args[0]
	if v, ok := a.Value.(*types.StringValue); ok {
		fmt.Print(v.Value)
	} else {
		fmt.Print(a.Value)
	}

	return types.NewAtom(types.NewNullValue())
}
