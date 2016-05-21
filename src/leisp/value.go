// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package leisp

type valueType int

const (
	valueKeyword valueType = iota
	valueSymbol
	valueNull
	valueNumber
	valueString
	valueChar
	valueBoolean
	valueInteger
	valueFloat
	valueRatio
)

type Value interface {
	ToString() string            // Returns a string representation of the value.
	GetValueType() valueType     // Return the valueType (enum of all Values).
	To(valueType) (Value, error) // Convert to a different Value type.
	OfType(string) bool          // Check if a given string is of this type.
	NewValue(string) Value       // Create a new Value of this type.
}

type IntType interface {
	ToString() string
	GetValueType() valueType
	To(valueType) (Value, error)
	OfType(string) bool
	NewValue(string) Value
}
