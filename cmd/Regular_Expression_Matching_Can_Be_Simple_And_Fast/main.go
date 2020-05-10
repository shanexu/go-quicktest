package main

import "fmt"

func main() {
	fmt.Println("hello world")
}

const Split rune = 256
const Match rune = 257

type State struct {
	c        rune
	out      *State
	out1     *State
	lastList int
}

func state(c rune, out, out1 *State) *State {
	return &State{c: c, out: out, out1: out1}
}

type Frag struct {
	start *State
	out   []*State
}

func frag(s *State, out []*State) *Frag {
	return &Frag{s, out}
}

func patch(l []*State, s *State) {
	for _, i := range l {
		i.out = s
	}
}

func str2nfa(str string) *State {
	stack := make(*Frag, 1000)
	push := func(s *Frag) {
		stack = append(stack, s)
	}

	pop := func(stack []*Frag) *Frag {
		f := stack[len(stack)-1]
		stack = stack[:len(stack)]
		return f
	}

	for r := range rs {
		switch r {
		case '*':
			e := pop()
			s := state(Split, e.start, nil)
		default:
			s := state(r, nil, nil)
			push(frag(s, nil))
		}
	}
}
