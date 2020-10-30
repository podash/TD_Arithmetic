package lab2

import (
	"regexp"
	"bytes"
)
// TODO: document this function.
// PostfixToPrefix converts
func StringToArray(inputS string) []string {

	var array []string
	
	var readingOperand bool = false

	for _, sym := range inputS {
		operand, _ := regexp.MatchString("[0-9]|[a-zA-Z]", string(sym))
		operator, _ := regexp.MatchString(`[+*^/-]`, string(sym))
		if operand {
			if readingOperand {
				var buffer bytes.Buffer
				buffer.WriteString(array[len(array) - 1]);
				buffer.WriteString(string(sym))
				array[len(array) - 1] = buffer.String()
			} else {
				array = append(array, string(sym))
				readingOperand = true
			}
		}	else if string(sym) == " " {
			readingOperand = false;
		}	else if operator {
			readingOperand = false;
			array = append(array, string(sym))
		}
	}

	return array
}


func PostfixToPrefix(inputStr string) ([]string, error) {

	var input = StringToArray(inputStr)
	
	stack := make([][]string, 0)

	for _, oper := range input {
		operator, _ := regexp.MatchString(`[+*^/-]`, oper) 
		if !operator {
			stack = append(stack, []string{oper})
		} else {
			op1 := stack[len(stack) - 1]
			op2 := stack[len(stack) - 2]
			stack = stack[:len(stack) - 2]
			newLastElement := append([]string{oper}, op2...)
			newLastElement = append(newLastElement, op1...)
			stack = append(stack, newLastElement)
		}
	}

	// return "TODO", fmt.Errorf("TODO")
	return stack[0], nil;
}
