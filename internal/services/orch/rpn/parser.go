package rpn

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/Onnywrite/lms-final/internal/domain/structs"
)

var (
	SymbolsErr = errors.New("expression contains restricted characters")
	ParsingErr = errors.New("error while parsing your expression")
)

func FromInfix(infixExpression string) (string, error) {
	if err := validateSymbols(infixExpression); err != nil {
		return "", err
	}
	if err := validateParentheses(infixExpression); err != nil {
		return "", err
	}
	infixExpression = removeSpaces(infixExpression)
	if rpn, err := parseRPN(infixExpression); err != nil {
		return "", err
	} else {
		return rpn, nil
	}
}

const (
	pow = '^'
	mul = '*'
	div = '/'
	sub = '-'
	add = '+'
)

func validateSymbols(e string) error {
	if !regexp.MustCompile(`^[0-9()*/+-]+$`).MatchString(e) {
		return SymbolsErr
	}
	return nil
}

func validateParentheses(e string) error {
	c := 0
	for i, r := range e {
		switch r {
		case '(':
			c++
		case ')':
			c--
		default:
			if c < 0 {
				return errors.New(fmt.Sprintf("expected ( to open the parenthesis at %d", i+1))
			}
		}
	}
	if c != 0 {
		return errors.New("not all parenthesis are closed")
	}
	return nil
}

func removeSpaces(expr string) string {
	return strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(expr, "\t", ""), " ", ""))
}

func parseRPN(infix string) (string, error) {
	s := structs.NewStack[rune]()
	buf := make([]rune, 0, 1024)

	pushIfNotLess := func(op rune) {
		var top rune
		if s.TryTop(&top) {
			for ; priority(top) >= priority(op); s.TryTop(&top) {
				buf = append(buf, s.Pop(), ' ')
			}
		}
		s.Push(op)
	}

	for _, r := range infix {
		switch r {
		case ')':
			for top := s.Pop(); top != '('; top = s.Pop() {
				buf = append(buf, top, ' ')
			}
		case pow:
			pushIfNotLess(r)
		case mul:
			pushIfNotLess(r)
		case div:
			pushIfNotLess(r)
		case add:
			pushIfNotLess(r)
		case sub:
			pushIfNotLess(r)
		case '(':
			s.Push(r)
		default:
			buf = append(buf, r)
		}
	}

	var op rune
	for s.TryPop(&op) {
		if op == '(' {
			return "", ParsingErr
		}
		buf = append(buf, op)
	}

	return string(buf), nil
}

func priority(op rune) int {
	switch op {
	case pow:
		return 2
	case mul:
		return 1
	case div:
		return 1
	case add:
		return 0
	case sub:
		return 0
	default:
		return -1
	}
}
