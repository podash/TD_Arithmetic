package lab2

import (
	"io"
	"io/ioutil"
	"fmt"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	Input io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	// TODO: Implement.
	buf, err := ioutil.ReadAll(ch.Input)
	fmt.Println(string(buf))
	return err
}
