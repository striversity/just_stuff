package main

import (
	"flag"
	"fmt"

	"github.com/sirupsen/logrus"
)

var (
	oper     string
	op1, op2 int
)

func main() {
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
			logrus.Error("op2 is equal to 0, can't divid by 0")
			return
		}
		res := op1 / op2
		fmt.Printf("%v / %v = %v\n", op1, op2, res)
	}
}
