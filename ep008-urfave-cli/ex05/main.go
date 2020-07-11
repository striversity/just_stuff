package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
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

	app.Flags = []cli.Flag{
		&cli.StringFlag{Destination: &oper, Name: "oper", Value: "add", Usage: "add, sub, mul, or div operation on two int operands"},
		&cli.IntFlag{Destination: &op1, Name: "op1", Usage: "operand 1 for operation"},
		&cli.IntFlag{Destination: &op2, Name: "op2", Usage: "operand 2 for operation"},
	}

	app.Action = mainAction

	app.Commands = []*cli.Command{
		addCommand(),
	}

	err := app.Run(os.Args)
	if err != nil {
		logrus.Error(err)
	}
}

func addCommand() *cli.Command {
	return &cli.Command{
		Name:    "add",
		Aliases: []string{"a"},
		Action: func(ctx *cli.Context) error {
			fmt.Printf("action/implementation for the add command\n")
			return nil
		},
	}
}

func mainAction(ctx *cli.Context) error {
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
