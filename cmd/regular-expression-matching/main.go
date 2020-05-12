package main

import "fmt"

func main() {
	fmt.Println(isMatch("a", "."))
	fmt.Println(isMatch("aa", ".*"))
	fmt.Println(isMatch("ab", "a*"))
}

func isMatch(s, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	pattern := compile([]rune(p))
	return pattern.match([]rune(s))
}

const (
	Split rune = iota + 256
	AnyChar
	Match
)

type Regexp struct {
	nstate     int
	l1         List
	l2         List
	start      *State
	listid     int
	matchState *State
}

func (re *Regexp) state(c rune, out, out1 *State) *State {
	re.nstate++
	return &State{c: c, out: out, out1: out1}
}

func (re *Regexp) match(s []rune) bool {
	re.l1 = make(List, 0, re.nstate)
	re.l2 = make(List, 0, re.nstate)
	clist := re.startList(re.start, &re.l1)
	nlist := &re.l2
	for _, c := range s {
		re.step(clist, c, nlist)
		clist, nlist = nlist, clist
	}
	return re.ismatch(clist)
}

type State struct {
	c        rune
	out      *State
	out1     *State
	lastList int
}

type Frag struct {
	start *State
	out   *PtrList
}

func frag(s *State, out *PtrList) *Frag {
	return &Frag{s, out}
}

type PtrList struct {
	next *PtrList
	s    **State
}

func list1(outp **State) *PtrList {
	p := &PtrList{s: outp}
	return p
}

func patch(l *PtrList, s *State) {
	var next *PtrList
	for ; l != nil; l = next {
		next = l.next
		*l.s = s
	}
}

func (l1 *PtrList) append(l2 *PtrList) *PtrList {
	oldl1 := l1
	for l1.next != nil {
		l1 = l1.next
	}
	l1.next = l2
	return oldl1
}

type Stack []*Frag

func stack(size int) Stack {
	return make(Stack, 0, size)
}

func (s *Stack) push(f *Frag) {
	*s = append(*s, f)
}

func (s *Stack) pop() *Frag {
	f := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return f
}

func compile(rs []rune) *Regexp {
	re := &Regexp{}
	st := stack(128)
	for _, p := range rs {
		switch p {
		case '.':
			s := re.state(AnyChar, nil, nil)
			st.push(frag(s, list1(&s.out)))
		case '*':
			e := st.pop()
			s := re.state(Split, e.start, nil)
			patch(e.out, s)
			st.push(frag(s, list1(&s.out1)))
		default:
			s := re.state(p, nil, nil)
			st.push(frag(s, list1(&s.out)))
		}
	}
	for ; len(st) > 1; {
		e2 := st.pop()
		e1 := st.pop()
		patch(e1.out, e2.start)
		st.push(frag(e1.start, e2.out))
	}
	e := st.pop()
	re.matchState = &State{Match, nil, nil, 0}
	patch(e.out, re.matchState)
	re.start = e.start
	return re
}

type List []*State

func (re *Regexp) startList(start *State, l *List) *List {
	*l = (*l)[:0]
	re.listid++
	re.addState(l, start)
	return l
}

func (re *Regexp) ismatch(l *List) bool {
	for i := range *l {
		if (*l)[i] == re.matchState {
			return true
		}
	}
	return false
}

func (re *Regexp) addState(l *List, s *State) {
	if s == nil || s.lastList == re.listid {
		return
	}
	s.lastList = re.listid
	if s.c == Split {
		re.addState(l, s.out)
		re.addState(l, s.out1)
		return
	}
	*l = append(*l, s)
}

func (re *Regexp) step(clist *List, c rune, nlist *List) {
	re.listid++
	*nlist = (*nlist)[:0]
	for i := range *clist {
		s := (*clist)[i]
		if s.c == c || s.c == AnyChar {
			re.addState(nlist, s.out)
		}
	}
}
