// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package interpreter

import (
	"fmt"
	"math"

	"github.com/leizongmin/leisp/types"
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

func builtinZerop(s *types.Scope, list []*types.AST) *types.Atom {

	if len(list) != 1 {
		return types.NewErrorAtom(fmt.Errorf("wrong arguments number for zero?"))
	}

	args, errAtom := astListToAtomList(s, list)
	if errAtom != nil {
		return errAtom
	}

	number := args[0]
	if integerValue, ok := number.Value.(*types.IntegerValue); ok {
		if integerValue.Value == 0 {
			return types.NewAtom(types.NewBooleanValue(true))
		}
		return types.NewAtom(types.NewBooleanValue(false))
	}
	if floatValue, ok := number.Value.(*types.FloatValue); ok {
		if floatValue.Value == 0 {
			return types.NewAtom(types.NewBooleanValue(true))
		}
		return types.NewAtom(types.NewBooleanValue(false))
	}

	return types.NewAtom(types.NewBooleanValue(false))
}

func builtinMathAdd(s *types.Scope, list []*types.AST) *types.Atom {

	args, errAtom := astListToAtomList(s, list)
	if errAtom != nil {
		return errAtom
	}

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

func builtinMathSubtract(s *types.Scope, list []*types.AST) *types.Atom {

	args, errAtom := astListToAtomList(s, list)
	if errAtom != nil {
		return errAtom
	}

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
			ret -= v
		}
		return types.NewAtom(types.NewIntegerValue(ret))
	}

	ret := floats[0]
	for _, v := range floats[1:] {
		ret -= v
	}
	return types.NewAtom(types.NewFloatValue(ret))
}

func builtinMathMultiply(s *types.Scope, list []*types.AST) *types.Atom {

	args, errAtom := astListToAtomList(s, list)
	if errAtom != nil {
		return errAtom
	}

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
			ret *= v
		}
		return types.NewAtom(types.NewIntegerValue(ret))
	}

	ret := floats[0]
	for _, v := range floats[1:] {
		ret *= v
	}
	return types.NewAtom(types.NewFloatValue(ret))
}

func builtinMathDivide(s *types.Scope, list []*types.AST) *types.Atom {

	args, errAtom := astListToAtomList(s, list)
	if errAtom != nil {
		return errAtom
	}

	if len(args) < 1 {
		return types.NewAtom(types.NewIntegerValue(0))
	}

	values, err := getAtomListFinalValues(s, args)
	if err != nil {
		return types.NewErrorAtom(err)
	}

	_, floats, _, err := getNumberValues(values)
	if err != nil {
		return types.NewErrorAtom(err)
	}

	ret := floats[0]
	for _, v := range floats[1:] {
		if v == 0 {
			return types.NewAtom(types.NewInfinityValue())
		}
		ret /= v
	}
	return types.NewAtom(types.NewFloatValue(ret))
}

func builtinMathPow(s *types.Scope, list []*types.AST) *types.Atom {

	args, errAtom := astListToAtomList(s, list)
	if errAtom != nil {
		return errAtom
	}

	if len(args) < 1 {
		return types.NewAtom(types.NewIntegerValue(0))
	}

	values, err := getAtomListFinalValues(s, args)
	if err != nil {
		return types.NewErrorAtom(err)
	}

	_, floats, _, err := getNumberValues(values)
	if err != nil {
		return types.NewErrorAtom(err)
	}

	ret := floats[0]
	for _, v := range floats[1:] {
		ret = math.Pow(ret, v)
	}
	return types.NewAtom(types.NewFloatValue(ret))
}
