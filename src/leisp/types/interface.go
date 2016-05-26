// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

type ValueType interface {
	ToString() string               // Return a string representation of the value
	GetType() string                // Return the valueType (enum of all Values)
	To(t string) (ValueType, error) // Convert to specified type
	IsValue() bool                  // Return true if it's an value
	EqualTo(t ValueType) bool       // Return true if it's equal to specified value
}
