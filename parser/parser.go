// leisp
// Copyright 2016 Zongmin Lei <leizongmin@gmail.com>. All rights reserved.
// Under the MIT License

package parser

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/leizongmin/leisp/types"
)

type Parser struct {
	s   *Scanner
	buf struct {
		tok token
		lit string
		n   int
	}
}

func NewParser(r io.Reader) *Parser {
	return &Parser{
		s: NewScanner(r),
	}
}

func ParseAll(prog string) (list []*types.AST, err error) {

	p := NewParser(strings.NewReader(prog))

	for {
		if ast, err := p.Parse(); err != nil {
			pos := p.GetPosition()
			return nil, fmt.Errorf("Error: %s at line %d,%d\n", err.Error(), pos.Line, pos.Column-1)
		} else if ast.IsEOF() {
			break
		} else {
			list = append(list, ast)
		}
	}

	return list, nil
}

func (p *Parser) scan() (tok token, lit string) {

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

func (p *Parser) scanIgnoreWhitespaceOrComment() (tok token, lit string) {
	for {
		tok, lit = p.scan()
		if tok != tokenWhitespace && tok != tokenComment {
			break
		}
	}
	return tok, lit
}

func (p *Parser) GetPosition() Position {
	return p.s.Position
}

func (p *Parser) Parse() (*types.AST, error) {

	tok, lit := p.scanIgnoreWhitespaceOrComment()
	switch tok {

	case tokenString:
		return types.NewValueAST(types.NewStringValue(lit)), nil

	case tokenEOF:
		return types.NewEOFAST(), nil

	case tokenNumber:
		return p.parseNumber(lit)

	case tokenKeyword:
		return types.NewValueAST(types.NewKeywordValue(lit)), nil

	case tokenSymbol:
		LIT := strings.ToLower(lit)
		if LIT == "null" {
			return types.NewValueAST(types.NewNullValue()), nil
		} else if LIT == "true" {
			return types.NewValueAST(types.NewBooleanValue(true)), nil
		} else if LIT == "false" {
			return types.NewValueAST(types.NewBooleanValue(false)), nil
		} else {
			return types.NewValueAST(types.NewSymbolValue(lit)), nil
		}

	case tokenQuote:
		return p.parseQuote(lit)

	case tokenPunctuation:
		return p.parsePunctuation(lit)

	default:
		return types.NewEOFAST(), fmt.Errorf("illegal token %s", lit)
	}
}

func (p *Parser) parseNumber(lit string) (*types.AST, error) {

	if i := strings.IndexAny(lit, "/"); i != -1 {
		return types.NewValueAST(types.NewRatioValue(lit)), nil
	}
	if i := strings.IndexAny(lit, "."); i != -1 {
		if val, err := strconv.ParseFloat(lit, 64); err != nil {
			return types.NewEOFAST(), fmt.Errorf("invalid float number %s", lit)
		} else {
			return types.NewValueAST(types.NewFloatValue(val)), nil
		}
	}
	if val, err := strconv.ParseInt(lit, 10, 64); err != nil {
		return types.NewEOFAST(), fmt.Errorf("invalid integer number %s", lit)
	} else {
		return types.NewValueAST(types.NewIntegerValue(val)), nil
	}
}

func (p *Parser) parsePunctuation(lit string) (*types.AST, error) {

	switch lit {

	case "(":
		return p.parseSExpression(lit)

	case "{":
		return p.parseQExpression(lit)

	case "[":
		return p.parseArray(lit)

	case ")", "}", "]":
		return types.NewEOFAST(), fmt.Errorf("mismatching %s", lit)

	default:
		return types.NewEOFAST(), fmt.Errorf("illegal token %s", lit)
	}
}

func (p *Parser) parseSExpression(lit string) (*types.AST, error) {

	var children []*types.AST

	for {
		tok, lit := p.scanIgnoreWhitespaceOrComment()
		if tok == tokenEOF {
			return nil, fmt.Errorf("unexpected end of s-expression")
		} else if tok == tokenPunctuation && lit == ")" {
			break
		} else {
			p.unscan()
			if ast, err := p.Parse(); err != nil {
				return types.NewEOFAST(), err
			} else {
				children = append(children, ast)
			}
		}
	}

	return types.NewSExpressionAST(children), nil
}

func (p *Parser) parseQExpression(lit string) (*types.AST, error) {

	var children []*types.AST

	for {
		tok, lit := p.scanIgnoreWhitespaceOrComment()
		if tok == tokenEOF {
			return nil, fmt.Errorf("unexpected end of q-expression")
		} else if tok == tokenPunctuation && lit == "}" {
			break
		} else {
			p.unscan()
			if ast, err := p.Parse(); err != nil {
				return types.NewEOFAST(), err
			} else {
				children = append(children, ast)
			}
		}
	}

	return types.NewQExpressionAST(children), nil
}

func (p *Parser) parseQuote(lit string) (*types.AST, error) {

	ast, err := p.Parse()
	if err != nil {
		return nil, err
	}

	if ast.IsEOF() {
		return ast, nil
	}
	if ast.IsValue() {
		switch ast.Value.(type) {
		case *types.KeywordValue, *types.SymbolValue:
			return types.NewQuoteAST(ast.Value), nil
		default:
			return ast, nil
		}
	}

	return types.NewQExpressionAST(ast.Children), nil
}

func (p *Parser) parseArray(lit string) (*types.AST, error) {

	var children []*types.AST

	for {
		tok, lit := p.scanIgnoreWhitespaceOrComment()
		if tok == tokenEOF {
			return nil, fmt.Errorf("unexpected end of list")
		} else if tok == tokenPunctuation && lit == "]" {
			break
		} else {
			p.unscan()
			if ast, err := p.Parse(); err != nil {
				return types.NewEOFAST(), err
			} else {
				children = append(children, ast)
			}
		}
	}

	return types.NewArrayAST(children), nil
}
