// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

type ValueType interface {
	ToString() string     // Returns a string representation of the value.
	GetValueType() string // Return the valueType (enum of all Values).
}