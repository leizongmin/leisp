// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package interpreter

import (
	"fmt"
	"leisp/types"
)

func builtinArrayIndex(s *types.Scope, list []*types.AST) *types.Atom {

	argc := len(list)
	if argc != 2 {
		return types.NewErrorMessageAtom("invalid arguments number for array-index")
	}

	array := EvalAST(s, list[0])
	if array.IsError() {
		return array
	}
	arr, ok := array.Value.(*types.ArrayValue)
	if !ok {
		return types.NewErrorAtom(fmt.Errorf("first argument must be array, actually is %s", array.Value.GetType()))
	}

	index := list[1]
	if !index.IsValue() {
		return types.NewErrorAtom(fmt.Errorf("second argument be integer: %s", index.ToString()))
	}
	i, ok := index.Value.(*types.IntegerValue)
	if !ok {
		return types.NewErrorAtom(fmt.Errorf("second argument be integer: %s", index.ToString()))
	}

	if int(i.Value) >= len(arr.Value) {
		return types.NewAtom(types.NewNullValue())
	}

	return types.NewAtom(arr.Value[i.Value])
}

func init() {

	RegisterBuiltinFunction("array-index", builtinArrayIndex)

}
