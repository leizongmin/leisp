// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import "fmt"

type KeywordValue struct {
	Value string
}

func (t *KeywordValue) ToString() string {
	return fmt.Sprint(t.Value)
}

func (t *KeywordValue) GetType() string {
	return "keyword"
}

func NewKeyword(v string) *KeywordValue {
	return &KeywordValue{Value: v}
}
