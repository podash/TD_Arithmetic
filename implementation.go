package lab2

import (
	"regexp"
	"strings"
)
// TODO: document this function.
// PostfixToPrefix converts
func StringToArray(inputS string) []string {
	re := regexp.MustCompile("[a-zA-Z]*[0-9]+|[+*^/-]")
	return re.FindAllString(inputS, -1)
}

func PostfixToPrefix(inputStr string) (string, error) {
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
	return strings.Join(stack[0], " "), nil;
}
