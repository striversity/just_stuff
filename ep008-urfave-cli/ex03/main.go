package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	oper     string
	op1, op2 int
)

func main() {
	app := cli.NewApp() // &cli.App{}
	app.Name = "mycalc"
	app.Usage = "The best calculator for the commandline"
	app.Description = "Looking for a neat commandline caculator? Look no further, this is the world's best commandline calculator, full stop."
	app.Authors = []*cli.Author{
		{Name: "Verrol Adams", Email: "youtube@striversity.com"},
	}

	app.Action = mainAction

	app.Run(os.Args)
}

func mainAction(ctx *cli.Context) error {
	flag.StringVar(&oper, "oper", "add", "add, sub, mul, or div operation on two int operands")
	flag.IntVar(&op1, "op1", 0, "operand 1 for operation")
	flag.IntVar(&op2, "op2", 0, "operand 2 for operation")
	flag.Parse()

	switch oper {
	case "add":
		res := op1 + op2
		fmt.Printf("%v + %v = %v\n", op1, op2, res)
	case "sub":
		res := op1 - op2
		fmt.Printf("%v - %v = %v\n", op1, op2, res)
	case "mul":
		res := op1 * op2
		fmt.Printf("%v * %v = %v\n", op1, op2, res)
	case "div":
		if op2 == 0 {
			return fmt.Errorf("op2 is equal to 0, can't divid by 0")
		}
		res := op1 / op2
		fmt.Printf("%v / %v = %v\n", op1, op2, res)
	}

	return nil
}
