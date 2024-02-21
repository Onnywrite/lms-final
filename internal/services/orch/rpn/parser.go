package rpn

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/Onnywrite/lms-final/internal/domain/structs"
)

var (
	SymbolsErr     = errors.New("expression contains restricted characters")
	NoOperationErr = errors.New("expression must contains at least one operation")
	ParsingErr     = errors.New("error while parsing your expression")
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

var (
	validSymbols = regexp.MustCompile(`^[ .0-9()*^/+-]+$`)
	operations   = regexp.MustCompile(`[*^/+-]+`)
)

func validateSymbols(e string) error {
	if !validSymbols.MatchString(e) {
		return SymbolsErr
	}
	if !operations.MatchString(e) {
		return NoOperationErr
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
			for s.TryTop(&top) && priority(top) >= priority(op) {
				buf = append(buf, ' ', s.Pop())
			}
		}
		s.Push(op)
	}

	infixRunes := []rune(infix)
	for i := 0; i < len(infixRunes); i++ {
		r := infixRunes[i]
		switch r {
		case ')':
			for top := s.Pop(); top != '('; top = s.Pop() {
				buf = append(buf, ' ', top)
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
			buf = append(buf, ' ')
			buf = append(buf, fetchNumber(infixRunes, &i)...)
			buf = append(buf, ' ')
		}
	}

	var op rune
	for s.TryPop(&op) {
		if op == '(' {
			return "", ParsingErr
		}
		buf = append(buf, ' ', op)
	}

	return strings.ReplaceAll(strings.TrimSpace(string(buf)), "  ", " "), nil
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

func fetchNumber(runes []rune, i *int) []rune {
	buf := make([]rune, 0, 32)

	for ; *i < len(runes); *i++ {
		if unicode.IsDigit(runes[*i]) || runes[*i] == '.' {
			buf = append(buf, runes[*i])
		} else {
			break
		}
	}
	*i--

	return buf
}
