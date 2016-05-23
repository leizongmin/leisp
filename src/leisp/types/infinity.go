// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

import "fmt"

type InfinityValue struct {
	Value int64
}

func (v *InfinityValue) ToString() string {
	return "Infinity"
}

func (v *InfinityValue) GetType() string {
	return "infinity"
}

func (v *InfinityValue) IsValue() bool {
	return true
}

func (v *InfinityValue) To(t string) (ValueType, error) {
	return nil, fmt.Errorf("cannot convert infinity to %s: does not implement yet", t)
}

func NewInfinityValue() *InfinityValue {
	return &InfinityValue{}
}
