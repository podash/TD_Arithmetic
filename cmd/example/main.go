package main

import (
	"flag"
	"fmt"
	"io"
	"strings"
	lab2 "github.com/Trafle/TD_Arithmetic"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile = flag.String("f", "", "File to read input from")
	outputFile = flag.String("o", "", "File to write output into")
)

func main() {
	flag.Parse()

	in := strings.NewReader(*inputExpression)
	var out io.Writer
	handler := &lab2.ComputeHandler {
		Input: in,
		Output: out,
	}
	err := handler.Compute()

	res, _ := &lab2.PostfixToPrefix("+ 2 2")
	fmt.Println(err)
	fmt.Println(res)
}
