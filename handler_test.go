package lab2

import (
	"testing"
	"strings"
	"bytes"
	"fmt"
)

func TestCompute(t *testing.T) {
	input := strings.NewReader("98 4 * 234 12 * +")
	var output bytes.Buffer
	handler := ComputeHandler {
		input,
		&output,
	}
	if err := handler.Compute(); err != nil {
		t.Fatalf("Error in Compute: %s", err)
	}
	res := output.String()
	if res == "" {
		t.Errorf("Error: Result not written into output")
	}
	expected := "+ * 98 4 * 234 12"
	if res != expected {
		t.Errorf("Incorrect result: got %q, expected %q", res, expected)
	}
	fmt.Println("Computed result: ", res)
}

func TestComputeIncorrect(t *testing.T) {
	input := strings.NewReader("98 hello * 234 12 * +")
	var output bytes.Buffer
	handler := ComputeHandler {
		input,
		&output,
	}
	err := handler.Compute()
	if err == nil {
		t.Errorf("Error mustn't be nil for incorrect input")
	}
	fmt.Println("Returned error message: ", err.Error())
}