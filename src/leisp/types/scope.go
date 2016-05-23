// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import "fmt"

type Scope struct {
	Variables map[string]ValueType
	Parent    *Scope
}

func NewScope(parent *Scope) *Scope {
	return &Scope{
		Parent:    parent,
		Variables: make(map[string]ValueType),
	}
}

func NewRootScope() *Scope {
	return NewScope(nil)
}

func (s *Scope) Get(name string) (val ValueType, err error) {
	if val, ok := s.Variables[name]; ok {
		return val, nil
	}
	if s.Parent != nil {
		return s.Parent.Get(name)
	}
	return nil, fmt.Errorf("%s is undefined", name)
}

func (s *Scope) Declare(name string, val ValueType) error {
	if _, ok := s.Variables[name]; ok {
		return fmt.Errorf("%s has already been declared", name)
	}
	s.Variables[name] = val
	return nil
}

func (s *Scope) Set(name string, val ValueType) error {
	if val, ok := s.Variables[name]; ok {
		s.Variables[name] = val
		return nil
	}
	if s.Parent != nil {
		s.Parent.Set(name, val)
		return nil
	}
	return fmt.Errorf("%s is undefined", name)
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

type ScopeValue struct {
	Value *Scope
}

func (v *ScopeValue) ToString() string {
	return fmt.Sprint(v.Value)
}

func (v *ScopeValue) GetType() string {
	return "scope"
}

func (v *ScopeValue) IsValue() bool {
	return true
}

func (v *ScopeValue) To(t string) (ValueType, error) {
	return nil, fmt.Errorf("cannot convert ratio to %s: does not implement yet", t)
}

func NewScopeValue(v *Scope) *ScopeValue {
	return &ScopeValue{Value: v}
}
