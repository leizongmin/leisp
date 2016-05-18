// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package leisp

type Token int

const (
	TokenIllegal Token = iota
	TokenEOF
	TokenWhitespace

	TokenSymbol
	TokenNumber
	TokenString
	TokenChar

	TokenPunctuation
)
