package main

type piece interface {
	String() string
	moveIsValid(p Posn) bool
}
