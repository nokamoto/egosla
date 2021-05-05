package main

import "log"

type state map[string]string

type scenario struct {
	name string
	run  func(state state) (state, error)
}

type scenarios []scenario

func (xs scenarios) run() {
	st := make(state)
	for _, x := range xs {
		log.Printf("%s:", x.name)
		s, err := x.run(st)
		if err != nil {
			log.Fatalf("%s: state=%v: %v", x.name, st, err)
		}
		st = s
		log.Printf("%s: ok", x.name)
	}
}
