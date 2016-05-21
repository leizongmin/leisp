// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package types

type ValueType interface {
	ToString() string            // Returns a string representation of the value.
	GetValueType() valueType     // Return the valueType (enum of all Values).
	To(valueType) (Value, error) // Convert to a different Value type.
	OfType(string) bool          // Check if a given string is of this type.
	NewValue(string) Value       // Create a new Value of this type.
}
