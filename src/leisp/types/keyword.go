// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import "fmt"

type KeywordValue struct {
	Value string
}

func (v *KeywordValue) ToString() string {
	return fmt.Sprint(v.Value)
}

func (v *KeywordValue) GetType() string {
	return "keyword"
}

func (v *KeywordValue) IsValue() bool {
	return false
}

func (v *KeywordValue) To(t string) (ValueType, error) {
	return nil, fmt.Errorf("cannot convert keyword to %s: does not implement yet", t)
}

func NewKeywordValue(v string) *KeywordValue {
	return &KeywordValue{Value: v}
}
