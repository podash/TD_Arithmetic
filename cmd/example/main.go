package main

import (
	"flag"
	"fmt"
	"io"
	"os"
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
	var in io.Reader
	var out io.Writer

	if *inputExpression != "" {
		in = strings.NewReader(*inputExpression)
	}	else if *inputFile != "" {
		file, e := os.Open(*inputFile)
		if e != nil {
        fmt.Println(e)
        return
    }
		in = file
	} else { return }

	if *outputFile != "" {
		file, e := os.Create(*outputFile)
		if e != nil {
        fmt.Println(e)
        return
    }
		out = file
	}
	handler := &lab2.ComputeHandler {
		Input: in,
		Output: out,
	}
	err := handler.Compute()

	if err != nil {
		fmt.Println(err)
	}
}
