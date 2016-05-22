// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import (
	"errors"
	"fmt"
)

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

func NewErrorMessageAtom(err string) *Atom {
	return NewErrorAtom(errors.New(err))
}

func (a *Atom) ToString() string {
	if a.Error != nil {
		return fmt.Sprintf("<Error#\"%s\">", a.Error)
	}
	return fmt.Sprint(a.Value.ToString())
}

func (a *Atom) Print() {
	fmt.Println(a.ToString())
}

func (a *Atom) IsError() bool {
	if a.Error != nil {
		return true
	}
	return false
}

func (a *Atom) IsValue() bool {
	if a.Error == nil {
		return true
	}
	return false
}
