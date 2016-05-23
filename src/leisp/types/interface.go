// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

type ValueType interface {
	ToString() string               // Returns a string representation of the value
	GetType() string                // Returns the valueType (enum of all Values)
	To(t string) (ValueType, error) // Convert to specified type
	IsValue() bool                  // Returns true if it's an value
}
