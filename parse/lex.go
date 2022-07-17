package parse

import (
	"fmt"
	"strconv"
	"unicode"
)

var (
	opMap = map[byte]interface{}{
		'+': nil,
		'-': nil,
		'*': nil,
		'/': nil,
	}
	num = map[byte]interface{}{
		'0': nil,
		'1': nil,
		'2': nil,
		'3': nil,
		'4': nil,
		'5': nil,
		'6': nil,
		'7': nil,
		'8': nil,
		'9': nil,
		'.': nil,
	}
)

type lex struct {
	b      []byte
	idx    int
	result Cal
	err    error
}

func (l *lex) setResult(r Cal) {
	l.result = r
}
func (l *lex) readFloat() float64 {
	i := l.idx
	for l.idx < len(l.b) {
		b := l.b[l.idx]
		if _, ok := num[b]; !ok {
			break
		}
		l.idx += 1
	}

	r, _ := strconv.ParseFloat(string(l.b[i:l.idx]), 10)
	return r
}
func (l *lex) peekByte() byte {
	b := l.b[l.idx]
	return b
}

func (l *lex) Lex(lval *yySymType) int {
	for l.idx < len(l.b) {
		b := l.peekByte()
		if unicode.IsSpace(rune(b)) {
			l.idx++
			continue
		}
		if b == '(' {
			l.idx++
			return '('
		} else if b == ')' {
			l.idx++
			return ')'
		}
		if _, ok := num[b]; ok {
			lval.val = l.readFloat()
			return NUM
		}
		if _, ok := opMap[b]; ok {
			lval.Op = string([]byte{b})
			l.idx++
			return int(b)
		}
		l.err = fmt.Errorf("unknown token %s", string([]byte{b}))
		return 0
	}
	return 0
}

func (l *lex) Error(s string) {
	if l.err == nil {
		l.err = fmt.Errorf(s)
	}
}

func Parse(s string) (*Expr, error) {
	l := &lex{
		b: []byte(s),
	}
	yyParse(l)
	if l.err != nil {
		return nil, l.err
	}
	return l.result.(*Expr), nil
}
