package main

import (
	"fmt"
	"os"

	calc "github.com/wasabi315/calc/calculator"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	m := calc.NewMock(os.Stdout)
	i := calc.NewInterpreter()
	c := calc.NewComposer([]calc.Calculator{m, i})

	calc.InterpretString(os.Args[1], c)

	fmt.Printf("%v => %v\n", os.Args[1], i.GetResult())
}
