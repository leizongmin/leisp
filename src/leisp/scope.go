// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package leisp

import "fmt"

type Scope struct {
	Variables map[string]interface{}
	Parent    *Scope
}

func NewScope(parent *Scope) *Scope {
	return &Scope{
		Parent: parent,
	}
}

func NewRootScope() *Scope {
	return &Scope{}
}

func (s *Scope) Get(name string) (val interface{}, err error) {
	if val, ok := s.Variables[name]; ok {
		return val, nil
	}
	if s.Parent != nil {
		return s.Parent.Get(name)
	}
	return nil, fmt.Errorf("'%s' is undefined", name)
}

func (s *Scope) Declare(name string, val interface{}) error {
	if _, ok := s.Variables[name]; ok {
		return fmt.Errorf("'%s' has already been declared", name)
	}
	s.Variables[name] = val
	return nil
}

func (s *Scope) Set(name string, val interface{}) error {
	if val, ok := s.Variables[name]; ok {
		s.Variables[name] = val
		return nil
	}
	if s.Parent != nil {
		s.Parent.Set(name, val)
		return nil
	}
	return fmt.Errorf("'%s' is undefined", name)
}

func (s *Scope) Delete(name string) error {
	if _, ok := s.Variables[name]; ok {
		delete(s.Variables, name)
		return nil
	}
	if s.Parent != nil {
		s.Parent.Delete(name)
		return nil
	}
	return nil
}
