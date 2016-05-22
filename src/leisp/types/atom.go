// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import "fmt"

type Atom struct {
	Error error
	Value ValueType
}

func NewAtom(v ValueType) *Atom {
	return &Atom{
		Value: v,
	}
}

func NewEmptyAtom() *Atom {
	return NewAtom(NewNull())
}

func NewErrorAtom(err error) *Atom {
	return &Atom{
		Error: err,
	}
}

func (r *Atom) ToString() string {
	if r.Error != nil {
		return fmt.Sprintf("<Error#\"%s\">", r.Error)
	}
	return fmt.Sprint(r.Value)
}

func (r *Atom) Print() {
	fmt.Println(r.ToString())
}
