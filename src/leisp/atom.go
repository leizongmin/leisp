// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package leisp

import "fmt"

type Atom struct {
	Error error
	Value interface{}
}

func newAtom(value interface{}) *Atom {
	return &Atom{
		Value: value,
	}
}

func newEmptyAtom() *Atom {
	return newAtom(nil)
}

func newErrorAtom(err error) *Atom {
	return &Atom{
		Error: err,
	}
}

func (r *Atom) ToString() string {
	if r.Error != nil {
		return fmt.Sprintf("<Error#\"%s\">", r.Error)
	}
	if arr, ok := r.Value.([]*AST); ok {
		arr2 := make([]interface{}, len(arr))
		for i, v := range arr {
			arr2[i] = v.Value
		}
		r.Value = arr2
	}
	return fmt.Sprint(r.Value)
}

func (r *Atom) Print() {
	fmt.Println(r.ToString())
}
