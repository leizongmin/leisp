// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package leisp

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Parser struct {
	s   *Scanner
	buf struct {
		tok Token
		lit string
		n   int
	}
}

func NewParser(r io.Reader) *Parser {
	return &Parser{
		s: NewScanner(r),
	}
}

func (p *Parser) scan() (tok Token, lit string) {

	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	tok, lit = p.s.Scan()
	p.buf.tok, p.buf.lit = tok, lit

	return tok, lit
}

func (p *Parser) unscan() {
	p.buf.n = 1
}

func (p *Parser) scanIgnoreWhitespaceOrComment() (tok Token, lit string) {
	for {
		tok, lit = p.scan()
		if tok != TokenWhitespace && tok != TokenComment {
			break
		}
	}
	return tok, lit
}

func (p *Parser) Parse() (*AST, error) {

	tok, lit := p.scanIgnoreWhitespaceOrComment()
	switch tok {
	case TokenChar:
		return NewCharAST(lit), nil
	case TokenString:
		return NewStringAST(lit), nil
	case TokenEOF:
		return NewEmptyAST(), nil
	case TokenNumber:
		return p.parseNumber(lit)
	case TokenKeyword:
		return NewKeywordAST(lit), nil
	case TokenSymbol:
		return NewSymbolAST(lit), nil
	case TokenPunctuation:
		return p.parsePunctuation(lit)
	default:
		return NewEmptyAST(), fmt.Errorf("illegal token %s", lit)
	}
}

func (p *Parser) parseNumber(lit string) (*AST, error) {
	if i := strings.IndexAny(lit, "/"); i != -1 {
		return NewRatioAST(lit), nil
	}
	if i := strings.IndexAny(lit, "."); i != -1 {
		if val, err := strconv.ParseFloat(lit, 64); err != nil {
			return NewFloatAST(val), nil
		} else {
			return NewEmptyAST(), fmt.Errorf("invalid float number %s", lit)
		}
	}
	if val, err := strconv.ParseInt(lit, 10, 64); err != nil {
		return NewIntegerAST(val), nil
	} else {
		return NewEmptyAST(), fmt.Errorf("invalid integer number %s", lit)
	}
}

func (p *Parser) parsePunctuation(lit string) (*AST, error) {

	switch lit {
	case "(":
		return p.parseSExpression(lit)
	case "{":
		return p.parseQExpression(lit)
	case "[":
		return p.parseList(lit)
	case ")", "}", "]":
		return NewEmptyAST(), fmt.Errorf("mismatching %s", lit)
	default:
		return NewEmptyAST(), fmt.Errorf("illegal token %s", lit)
	}
}

func (p *Parser) parseSExpression(lit string) (*AST, error) {

	var children []*AST
	_, _ = p.scan()

	for {
		tok, lit := p.scanIgnoreWhitespaceOrComment()
		if tok == TokenPunctuation && lit == ")" {
			break
		} else {
			if ast, err := p.Parse(); err != nil {
				children = append(children, ast)
			} else {
				return NewSExpressionAST(children), err
			}
		}
	}

	return NewSExpressionAST(children), nil
}

func (p *Parser) parseQExpression(lit string) (*AST, error) {
	return NewEmptyAST(), fmt.Errorf("empty q-expression %s", lit)
}

func (p *Parser) parseList(lit string) (*AST, error) {
	return NewEmptyAST(), fmt.Errorf("empty list %s", lit)
}
