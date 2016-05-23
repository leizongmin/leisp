// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package interpreter

import (
	"fmt"
	"leisp/types"
)

func getNumberValues(list []types.ValueType) (integers []int64, floats []float64, isInteger bool, err error) {
	isInteger = true
	size := len(list)
	integers = make([]int64, size)
	floats = make([]float64, size)
	for i, v := range list {
		if vi, ok := v.(*types.IntegerValue); ok {
			integers[i] = vi.Value
			floats[i] = float64(vi.Value)
			continue
		}
		if vf, ok := v.(*types.FloatValue); ok {
			integers[i] = int64(vf.Value)
			floats[i] = vf.Value
			isInteger = false
			continue
		}
		return nil, nil, isInteger, fmt.Errorf("type %s is not a number: %s", v.GetType(), v.ToString())
	}
	return integers, floats, isInteger, nil
}

func builtinMathAdd(s *types.Scope, args []*types.Atom) *types.Atom {

	if len(args) < 1 {
		return types.NewAtom(types.NewIntegerValue(0))
	}

	values, err := getAtomListFinalValues(s, args)
	if err != nil {
		return types.NewErrorAtom(err)
	}

	integers, floats, isInteger, err := getNumberValues(values)
	if err != nil {
		return types.NewErrorAtom(err)
	}

	if isInteger {
		ret := integers[0]
		for _, v := range integers[1:] {
			ret += v
		}
		return types.NewAtom(types.NewIntegerValue(ret))
	}

	ret := floats[0]
	for _, v := range floats[1:] {
		ret += v
	}
	return types.NewAtom(types.NewFloatValue(ret))
}

func init() {

	RegisterBuiltinFunction("+", builtinMathAdd)

}
