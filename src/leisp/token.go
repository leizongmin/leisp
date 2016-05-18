// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package leisp

type Token int

const (
	TokenIllegal Token = itoa
	TokenEOF
	TokenWhitespace

	TokenSymbol
	TokenNumber
	TokenString

	TokenColon // :
	TokenComma // ,

	TokenRoundBacketLeft  // (
	TokenRoundBacketRight // )

	TokenSquareBacketLeft  // [
	TokenSquareBacketRight // ]

	TokenCurlyBacketLeft  // {
	TokenCurlyBacketRight // {

	TokenDoubleQuotation // "
	TokenSingleQuotation // '
)
