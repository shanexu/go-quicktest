package main

import (
	"fmt"
)

func main() {
	str := "aaaaaaaaaaaaab"
	pat := "a*a*a*a*a*a*a*a*a*a*c"
	fmt.Println(isMatch(str, pat))
}

func isMatch(s, p string) bool {
	pattern := compile(p)
	return match(State{
		rs: []rune(s),
		p:  pattern,
	})
}

func match(state State) bool {
	ss := state.p.Match(state.rs)
	for _, s := range ss {
		if s.p == nil {
			if len(s.rs) == 0 {
				return true
			}
			continue
		}
		if match(s) {
			return true
		}
	}
	return false
}

type Pattern interface {
	Match(runes []rune) []State
	SetNext(next Pattern)
}

type State struct {
	rs []rune
	p  Pattern
}

type Begin struct {
	next Pattern
}

func (b *Begin) SetNext(next Pattern) {
	b.next = next
}

func (b *Begin) Match(source []rune) []State {
	return []State{{source, b.next}}
}

type Constant struct {
	rs   []rune
	next Pattern
}

func (c *Constant) SetNext(next Pattern) {
	c.next = next
}

func (c *Constant) Match(source []rune) []State {
	if len(source) < len(c.rs) {
		return nil
	}
	for i := 0; i < len(c.rs); i++ {
		if c.rs[i] != source[i] {
			return nil
		}
	}
	return []State{{source[len(c.rs):], c.next}}
}

type OneAny struct {
	next Pattern
}

func (a *OneAny) SetNext(next Pattern) {
	a.next = next
}

func (a *OneAny) Match(source []rune) []State {
	if len(source) < 1 {
		return nil
	}
	return []State{{source[1:], a.next}}
}

type Many struct {
	r    rune
	next Pattern
}

func (m *Many) SetNext(next Pattern) {
	m.next = next
}

func (m *Many) Match(source []rune) []State {
	if len(source) == 0 {
		return []State{{source, m.next}}
	}
	if source[0] == m.r {
		return []State{
			{source, m.next},
			{source[1:], m.next},
			{source[1:], m},
		}
	}
	return []State {
		{source, m.next},
	}
}

type ManyAny struct {
	next Pattern
}

func (m *ManyAny) SetNext(next Pattern) {
	m.next = next
}

func (m *ManyAny) Match(source []rune) []State {
	if len(source) == 0 {
		return []State{{source, m.next}}
	}
	return []State{
		{source, m.next},
		{source[1:], m.next},
		{source[1:], m},
	}
}

func compile(str string) Pattern {
	rs := []rune(str)
	var p Pattern = &Begin{}
	begin := p
	start := -1
	for i := 0; i < len(rs); {
		r0 := rs[i]
		if i+1 < len(rs) {
			r1 := rs[i+1]
			switch r0 {
			case '.':
				if start != -1 {
					c := &Constant{rs[start:i], nil}
					p.SetNext(c)
					p = c
					start = -1
				}
				if r1 == '*' {
					m := &ManyAny{}
					p.SetNext(m)
					p = m
					i+=2
				} else {
					o := &OneAny{}
					p.SetNext(o)
					p = o
					i+=1
				}
			default:
				if start == -1 {
					start = i
				}
				if r1 == '*' {
					if start != i {
						c := &Constant{rs[start:i], nil}
						p.SetNext(c)
						p = c
					}
					start = -1
					ma := &Many{r: r0}
					p.SetNext(ma)
					p = ma
					i+=2
				} else {
					i+=1
				}
			}
		} else {
			switch r0 {
			case '.':
				if start != -1 {
					c := &Constant{rs[start:i], nil}
					p.SetNext(c)
					p = c
					start = -1
				}
				o := &OneAny{}
				p.SetNext(o)
				p = o
			default:
				if start == -1 {
					start = i
				}
				c := &Constant{rs: rs[start:i+1]}
				p.SetNext(c)
				p = c
			}
			i+=1
		}

	}
	return begin
}
