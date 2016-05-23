// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package interpreter

import (
	"leisp/types"
	"strings"
)

func builtinStr(s *types.Scope, list []*types.AST) *types.Atom {

	args, errAtom := astListToAtomList(s, list)
	if errAtom != nil {
		return errAtom
	}

	values, err := getAtomListFinalValues(s, args)
	if err != nil {
		return types.NewErrorAtom(err)
	}

	strs := make([]string, len(values))
	for i, v := range values {
		if s, ok := v.(*types.StringValue); ok {
			strs[i] = s.Value
		} else {
			strs[i] = v.ToString()
		}
	}

	return types.NewAtom(types.NewStringValue(strings.Join(strs, "")))
}

func init() {

	RegisterBuiltinFunction("str", builtinStr)

}
