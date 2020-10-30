package lab2

import (
	"io"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	Input io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	// TODO: Implement.
	buf := make([]byte, 1024)
	for {
		n, err := ch.Input.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 { 
			break 
		}
		_, err = ch.Output.Write(buf[:n])
		if err != nil {
			return err
		}
	}
	return nil
}
