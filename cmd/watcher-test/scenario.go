package main

type state map[string]string

type scenario struct {
	name string
	run  func(state state) (state, error)
}
