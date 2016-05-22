// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package interpreter

import "leisp/types"

func init() {

	RegisterBuiltinFunction("type-of", func(s *types.Scope, args []*types.Atom) *types.Atom {
		if len(args) != 1 {
			return types.NewErrorMessageAtom(`invalid arguments number for "type-of"`)
		}
		a := args[0]
		if a.IsError() {
			return types.NewAtom(types.NewString("error"))
		}
		return types.NewAtom(types.NewString(a.Value.GetValueType()))
	})

}
