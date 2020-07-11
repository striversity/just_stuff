package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
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

	app.Commands = []*cli.Command{
		addCommand(),
		subCommand(),
		mulCommand(),
		divCommand(),
	}

	err := app.Run(os.Args)
	if err != nil {
		logrus.Error(err)
	}
}

func mulCommand() *cli.Command {
	return &cli.Command{
		Name:    "mul",
		Usage:   "used to multiply several integers",
		Aliases: []string{"m"},
		Action: func(ctx *cli.Context) error {
			n := ctx.NArg()
			if n == 0 {
				return fmt.Errorf("no argument provided for mul operation")
			}

			a := ctx.Args().Get(0)
			res, _ := strconv.Atoi(a)
			fmt.Print(res)

			for i := 1; i < n; i++ {
				a = ctx.Args().Get(i)
				op, _ := strconv.Atoi(a)
				res *= op
				fmt.Printf(" * %v", op)
			}

			fmt.Printf(" = %v\n", res)
			return nil
		},
	}
}

func divCommand() *cli.Command {
	return &cli.Command{
		Name:    "div",
		Usage:   "used to divide several integers",
		Aliases: []string{"d"},
		Action: func(ctx *cli.Context) error {
			n := ctx.NArg()
			if n == 0 {
				return fmt.Errorf("no argument provided for div operation")
			}

			a := ctx.Args().Get(0)
			res, _ := strconv.Atoi(a)
			fmt.Print(res)

			for i := 1; i < n; i++ {
				a = ctx.Args().Get(i)
				op, _ := strconv.Atoi(a)

				if op == 0 {
					return fmt.Errorf("an operand is equal to 0, can't divid by 0")
				}

				res /= op
				fmt.Printf(" / %v", op)
			}

			fmt.Printf(" = %v\n", res)
			return nil
		},
	}
}

func subCommand() *cli.Command {
	return &cli.Command{
		Name:    "sub",
		Usage:   "used to subtract several integers",
		Aliases: []string{"s"},
		Action: func(ctx *cli.Context) error {
			n := ctx.NArg()
			if n == 0 {
				return fmt.Errorf("no argument provided for sub operation")
			}

			a := ctx.Args().Get(0)
			res, _ := strconv.Atoi(a)
			fmt.Print(res)

			for i := 1; i < n; i++ {
				a = ctx.Args().Get(i)
				op, _ := strconv.Atoi(a)
				res -= op
				fmt.Printf(" - %v", op)
			}

			fmt.Printf(" = %v\n", res)
			return nil
		},
	}
}

func addCommand() *cli.Command {
	return &cli.Command{
		Name:    "add",
		Aliases: []string{"a"},
		Action: func(ctx *cli.Context) error {
			n := ctx.NArg()
			if n == 0 {
				return fmt.Errorf("no argument provided for add operation")
			}

			a := ctx.Args().Get(0)
			res, _ := strconv.Atoi(a)
			fmt.Print(res)

			for i := 1; i < n; i++ {
				a = ctx.Args().Get(i)
				op, _ := strconv.Atoi(a)
				res += op
				fmt.Printf(" + %v", op)
			}

			fmt.Printf(" = %v\n", res)
			return nil
		},
	}
}

func mainAction(ctx *cli.Context) error {
	ctx.App.Command("help").Run(ctx)

	return nil
}
