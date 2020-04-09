package calculator

import (
	"fmt"
	"io"
)

type mock struct {
	io.Writer
}

func NewMock(w io.Writer) mock {
	return mock{w}
}

func (m mock) Num(n int) {
	fmt.Fprintf(m, "Num %v\n", n)
}

func (m mock) Add() {
	fmt.Fprintln(m, "Add")
}
func (m mock) Sub() {
	fmt.Fprintln(m, "Sub")
}
func (m mock) Mul() {
	fmt.Fprintln(m, "Mul")
}
func (m mock) Div() {
	fmt.Fprintln(m, "Div")
}

func (m mock) SigInv() {
	fmt.Fprintln(m, "SigInv")
}

func (m mock) Calc() {
	fmt.Fprintln(m, "Calc")
}
