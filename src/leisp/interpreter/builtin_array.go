// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package interpreter

import (
	"fmt"
	"leisp/types"
)

func builtinArrayLength(s *types.Scope, list []*types.AST) *types.Atom {

	argc := len(list)
	if argc != 1 {
		return types.NewErrorMessageAtom("invalid arguments number for array-length")
	}

	array := EvalAST(s, list[0])
	if array.IsError() {
		return array
	}
	arr, ok := array.Value.(*types.ArrayValue)
	if !ok {
		return types.NewErrorAtom(fmt.Errorf("first argument must be array, actually is %s", array.Value.GetType()))
	}

	return types.NewAtom(types.NewIntegerValue(int64(len(arr.Value))))
}

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

func builtinArraySlice(s *types.Scope, list []*types.AST) *types.Atom {

	argc := len(list)
	if argc != 3 {
		return types.NewErrorMessageAtom("invalid arguments number for array-slice")
	}

	array := EvalAST(s, list[0])
	if array.IsError() {
		return array
	}
	arr, ok := array.Value.(*types.ArrayValue)
	if !ok {
		return types.NewErrorAtom(fmt.Errorf("first argument must be array, actually is %s", array.Value.GetType()))
	}

	start := list[1]
	if !start.IsValue() {
		return types.NewErrorAtom(fmt.Errorf("second argument must be integer, actually is: %s", start.ToString()))
	}
	b, ok := start.Value.(*types.IntegerValue)
	if !ok {
		return types.NewErrorAtom(fmt.Errorf("second argument must be integer, actually is: %s", start.Value.GetType()))
	}

	end := list[2]
	if !end.IsValue() {
		return types.NewErrorAtom(fmt.Errorf("third argument must be integer, actually is: %s", end.ToString()))
	}
	e, ok := end.Value.(*types.IntegerValue)
	if !ok {
		return types.NewErrorAtom(fmt.Errorf("third argument must be integer, actually is: %s", end.Value.GetType()))
	}

	if b.Value < 0 {
		b.Value = 0
	}
	if int(e.Value) >= len(arr.Value) {
		e.Value = int64(len(arr.Value))
	}

	return types.NewAtom(types.NewArrayValue(arr.Value[b.Value:e.Value]))
}

func init() {

	RegisterBuiltinFunction("array-length", builtinArrayLength)
	RegisterBuiltinFunction("array-index", builtinArrayIndex)
	RegisterBuiltinFunction("array-slice", builtinArraySlice)

}
