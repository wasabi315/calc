package calculator

// Calculator methods represent keys of a calculator.
type Calculator interface {
	Num(n int) // Num keys (0-9)
	Add()      // +
	Sub()      // -
	Mul()      // *
	Div()      // /
	SigInv()   // ±
	Calc()     // =
}

func InterpretRune(tok rune, c Calculator) {
	switch {
	case '0' <= tok && tok <= '9':
		c.Num(int(tok - '0'))
	case tok == '+':
		c.Add()
	case tok == '-':
		c.Sub()
	case tok == '*':
		c.Mul()
	case tok == '/':
		c.Div()
	case tok == 'S':
		c.SigInv()
	case tok == '=':
		c.Calc()
	}
}

func InterpretString(s string, c Calculator) {
	for _, r := range []rune(s) {
		InterpretRune(r, c)
	}
}
