package brackets

import "fmt"

type TBracketClosable uint8

const (
	TBracketOpen  TBracketClosable = 0x01
	TBracketClose TBracketClosable = 0x02
)

type TBracketForm uint8

const (
	TBracketSquare    TBracketForm = 0x01
	TBracketFigure    TBracketForm = 0x02
	TBracketRound     TBracketForm = 0x03
	TBracketMultiPlus TBracketForm = 0x04
)

type TBracketProp struct {
	OpenT TBracketClosable
	FormT TBracketForm
}

var bkt = map[rune]TBracketProp{
	'[': {TBracketOpen, TBracketSquare},
	'{': {TBracketOpen, TBracketFigure},
	'(': {TBracketOpen, TBracketRound},
	'*': {TBracketOpen, TBracketMultiPlus},
	']': {TBracketClose, TBracketSquare},
	'}': {TBracketClose, TBracketFigure},
	')': {TBracketClose, TBracketRound},
	'+': {TBracketClose, TBracketMultiPlus},
}

type StEl struct {
	v byte
	c int
}

type Stack struct {
	el []StEl
}

func (s *Stack) Push(e rune, c int) {
	s.el = append(s.el, StEl{
		v: byte(e),
		c: c,
	})
}

func (s *Stack) Pop() (StEl, error) {
	if s.IsEmpty() {
		return StEl{}, fmt.Errorf("stack is empty")
	} else {
		index := len(s.el) - 1
		element := s.el[index]
		s.el = s.el[:index]

		return element, nil
	}
}

func (s *Stack) IsEmpty() bool {
	return len(s.el) == 0
}

type IStack interface {
	Push(e rune, c int)
	Pop() (StEl, error)
	IsEmpty() bool
}

func MustNewBracket() *Brackets {
	return &Brackets{
		s: &Stack{},
	}
}

type Brackets struct {
	s      IStack
	isFail bool
}

func (b *Brackets) add(e rune, c int) {
	b.s.Push(e, c)
}

func (b *Brackets) revoke(e rune) {
	n, err := b.s.Pop()
	if err != nil {
		b.isFail = true
		return
	}

	if v, ok := bkt[rune(n.v)]; ok {
		if v2, ok := bkt[e]; ok {
			if v2.FormT != v.FormT {
				b.isFail = true
				return
			}
		}
	}
}

func (b *Brackets) isValid() bool {
	return !b.isFail
}

func (b *Brackets) Validate(s string) int {
	for k, v := range s {
		if val, ok := bkt[v]; ok {
			if val.OpenT == TBracketOpen {
				b.add(v, k+1)
			} else {
				b.revoke(v)
			}
			if !b.isValid() {
				return k + 1
			}
		}
	}

	if !b.s.IsEmpty() {
		e, _ := b.s.Pop()
		return e.c
	}

	return 0
}
