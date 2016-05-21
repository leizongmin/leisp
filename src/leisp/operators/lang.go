// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package builtin

func init() {

	Register("setf", Function(func(args []*Atom) *Atom {
		return newAtom("setf")
	}))

}

// type-of, is-int32?
// int32, int64, int, float32, float64, string ...
