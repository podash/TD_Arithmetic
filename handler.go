package lab2

import (
	"io"
	"io/ioutil"
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
	if err != nil && err != io.EOF {
		return err
	}
	//res, err := PostfixToPrefix(string(buf))
	//if err != nil {
	//	return err
	//}
	_, err = ch.Output.Write(buf)
	if err != nil {
		return err
	}
	return nil
}
