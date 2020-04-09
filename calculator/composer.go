package calculator

type composer struct {
	cs []Calculator
}

func NewComposer(cs []Calculator) *composer {
	return &composer{cs}
}

func (c *composer) foreach(fn func(Calculator)) {
	for i := range c.cs {
		cal := c.cs[i]
		fn(cal)
	}
}

func (c *composer) Num(n int) {
	c.foreach(func(cal Calculator) {
		cal.Num(n)
	})
}

func (c *composer) Add() {
	c.foreach(func(cal Calculator) {
		cal.Add()
	})
}
func (c *composer) Sub() {
	c.foreach(func(cal Calculator) {
		cal.Sub()
	})
}
func (c *composer) Mul() {
	c.foreach(func(cal Calculator) {
		cal.Mul()
	})
}
func (c *composer) Div() {
	c.foreach(func(cal Calculator) {
		cal.Div()
	})
}

func (c *composer) SigInv() {
	c.foreach(func(cal Calculator) {
		cal.SigInv()
	})
}

func (c *composer) Calc() {
	c.foreach(func(cal Calculator) {
		cal.Calc()
	})
}
