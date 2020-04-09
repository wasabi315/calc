package calculator

type calcState int

const (
	calcOp calcState = iota
	calcNum
)

type operator int

const (
	add operator = iota
	sub
	mul
	div
)

type interpreter struct {
	inp, res int
	st       calcState
	op       operator
}

func NewInterpreter() *interpreter {
	return &interpreter{0, 0, calcOp, add}
}

func (i *interpreter) GetResult() int {
	return i.res
}

func (i *interpreter) Num(n int) {
	if i.st == calcOp {
		i.st = calcNum
		i.inp = n
		return
	}
	if i.inp < 0 {
		i.inp = i.inp*10 - n
	} else {
		i.inp = i.inp*10 + n
	}
}

func (i *interpreter) execOp() {
	switch i.op {
	case add:
		i.res += i.inp
	case sub:
		i.res -= i.inp
	case mul:
		i.res *= i.inp
	case div:
		i.res /= i.inp
	}
}

func (i *interpreter) Add() {
	if i.st == calcNum {
		i.execOp()
		i.st = calcOp
	}
	i.op = add
}

func (i *interpreter) Sub() {
	if i.st == calcNum {
		i.execOp()
		i.st = calcOp
	}
	i.op = sub
}

func (i *interpreter) Mul() {
	if i.st == calcNum {
		i.execOp()
		i.st = calcOp
	}
	i.op = mul
}

func (i *interpreter) Div() {
	if i.st == calcNum {
		i.execOp()
		i.st = calcOp
	}
	i.op = div
}

func (i *interpreter) SigInv() {
	if i.st == calcNum {
		i.inp = -i.inp
	}
}

func (i *interpreter) Calc() {
	if i.st == calcNum {
		i.execOp()
	}
	i.inp = 0
	i.op = add
	i.st = calcOp
}
