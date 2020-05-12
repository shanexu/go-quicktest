package main

import "fmt"

func main() {
	s := post2nfa([]rune{'s', '*'})
	l1 := make(List, 0, nstate)
	l2 := make(List, 0, nstate)
	m := match(s, []rune("si"), l1, l2)
	fmt.Println(m)
}

const (
	Split rune = iota + 256
	Cat
	AnyChar
	Match
)

var matchstate = State{Match, nil, nil, 0}
var nstate int
var listid int

type State struct {
	c        rune
	out      *State
	out1     *State
	lastList int
}

func state(c rune, out, out1 *State) *State {
	nstate++
	return &State{c: c, out: out, out1: out1}
}

type Frag struct {
	start *State
	out   *Ptrlist
}

func frag(s *State, out *Ptrlist) *Frag {
	return &Frag{s, out}
}

type Ptrlist struct {
	next *Ptrlist
	s    **State
}

func list1(outp **State) *Ptrlist {
	p := &Ptrlist{s: outp}
	return p
}

func patch(l *Ptrlist, s *State) {
	var next *Ptrlist
	for ; l != nil; l = next {
		next = l.next
		*l.s = s
	}
}

func (l1 *Ptrlist) append(l2 *Ptrlist) *Ptrlist {
	oldl1 := l1
	for l1.next != nil {
		l1 = l1.next
	}
	l1.next = l2
	return oldl1
}

func post2nfa(postfix []rune) *State {
	if len(postfix) == 0 {
		return nil
	}

	stack := make([]*Frag, 0, 1024)
	push := func(s *Frag) {
		stack = append(stack, s)
	}

	pop := func() *Frag {
		f := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return f
	}
	for _, p := range postfix {
		switch p {
		case Cat:
			e2 := pop()
			e1 := pop()
			patch(e1.out, e2.start)
			push(frag(e1.start, e2.out))
		case '|':
			e2 := pop()
			e1 := pop()
			s := state(Split, e1.start, e2.start)
			push(frag(s, e1.out.append(e2.out)))
		case '?':
			e := pop()
			s := state(Split, e.start, nil)
			push(frag(s, e.out.append(list1(&s.out1))))
		case '*':
			e := pop()
			s := state(Split, e.start, nil)
			patch(e.out, s)
			push(frag(s, list1(&s.out1)))
		case '.':
			s := state(AnyChar, nil, nil)
			push(frag(s, list1(&s.out)))
		default:
			s := state(p, nil, nil)
			push(frag(s, list1(&s.out)))
		}
	}
	e := pop()
	if len(stack) != 0 {
		return nil
	}
	patch(e.out, &matchstate)
	return e.start
}

type List []*State

func startlist(start *State, l *List) *List {
	*l = (*l)[:0]
	listid++
	addState(l, start)
	return l
}

func ismatch(l *List) bool {
	for i := range *l {
		if (*l)[i] == &matchstate {
			return true
		}
	}
	return false
}

func addState(l *List, s *State) {
	if s == nil || s.lastList == listid {
		return
	}
	s.lastList = listid
	if s.c == Split {
		addState(l, s.out)
		addState(l, s.out1)
		return
	}
	*l = append(*l, s)
}

func step(clist *List, c rune, nlist *List) {
	listid++
	*nlist = (*nlist)[:0]
	for i := range *clist {
		s := (*clist)[i]
		if s.c == c || s.c == AnyChar {
			addState(nlist, s.out)
		}
	}
}

func match(start *State, s []rune, l1 List, l2 List) bool {
	clist := startlist(start, &l1)
	nlist := &l2
	for _, c := range s {
		step(clist, c, nlist)
		clist, nlist = nlist, clist
	}
	return ismatch(clist)
}
