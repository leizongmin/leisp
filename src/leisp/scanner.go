// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package leisp

import (
	"bufio"
	"bytes"
	"io"
)

const eof = rune(0)

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{
		r: bufio.NewReader(r),
	}
}

func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func (s *Scanner) unread() {
	_ = s.r.UnreadRune()
}

func (s *Scanner) Scan() (tok Token, lit string) {

	ch := s.read()

	if isWhitesapce(ch) {
		s.unread()
		return s.scanWhitespace()
	}
	if isDigit(ch) {
		s.unread()
		return s.scanNumber()
	}

	switch ch {
	case eof:
		return TokenEOF, ""
	case '"':
		s.unread()
		return s.scanString()
	case '\'':
		s.unread()
		return s.scanChar()
	case ';':
		s.unread()
		return s.scanComment()
	case ':':
		s.unread()
		return s.scanKeyword()
	}

	if isPunctuation(ch) {
		return TokenPunctuation, string(ch)
	}

	s.unread()
	return s.scanSymbol()
}

func (s *Scanner) scanWhitespace() (tok Token, lit string) {

	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitesapce(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return TokenWhitespace, buf.String()
}

func (s *Scanner) scanString() (tok Token, lit string) {

	var buf bytes.Buffer
	s.read()

	for {
		if ch := s.read(); ch == eof {
			return TokenIllegal, buf.String()
		} else if ch == '\\' {
			buf.WriteRune(ch)
			ch = s.read()
			buf.WriteRune(ch)
		} else if ch == '"' {
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return TokenString, buf.String()
}

func (s *Scanner) scanNumber() (tok Token, lit string) {

	var buf bytes.Buffer
	buf.WriteRune(s.read())

	isSlash := false
	isDot := false
	isE := false

	for {
		if ch := s.read(); ch == eof {
			break
		} else if isDigit(ch) {
			buf.WriteRune(ch)
		} else if ch == '/' {
			if isSlash {
				return TokenIllegal, string(ch)
			} else {
				buf.WriteRune(ch)
				isSlash = true
			}
		} else if ch == '.' {
			if isDot {
				return TokenIllegal, string(ch)
			} else {
				buf.WriteRune(ch)
				isDot = true
			}
		} else if ch == 'e' || ch == 'E' {
			if isE {
				return TokenIllegal, string(ch)
			} else {
				buf.WriteRune(ch)
				isE = true
			}
		} else if isWhitesapce(ch) || isPunctuation(ch) {
			s.unread()
			break
		} else {
			return TokenIllegal, string(ch)
		}
	}

	return TokenNumber, buf.String()
}

func (s *Scanner) scanChar() (tok Token, lit string) {

	var buf bytes.Buffer
	s.read()

	if ch := s.read(); ch == '\\' {
		buf.WriteRune(s.read())
	} else if ch == '\'' {
		return TokenIllegal, string(ch)
	} else {
		buf.WriteRune(ch)
	}

	ch := s.read()
	if ch != '\'' {
		return TokenIllegal, string(ch)
	}

	return TokenChar, buf.String()
}

func (s *Scanner) scanSymbol() (tok Token, lit string) {

	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if isWhitesapce(ch) {
			s.unread()
			break
		} else if isPunctuation(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return TokenSymbol, buf.String()
}

func (s *Scanner) scanKeyword() (tok Token, lit string) {

	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if isWhitesapce(ch) {
			s.unread()
			break
		} else if ch == ':' {
			return TokenIllegal, string(ch)
		} else if isPunctuation(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return TokenKeyword, buf.String()
}

func (s *Scanner) scanComment() (tok Token, lit string) {

	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if ch == '\n' {
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return TokenComment, buf.String()
}
