package lab2

import (
	"fmt"
	"testing"
	. "gopkg.in/check.v1"
)

func TestImplementation (t *testing.T) { TestingT(t) }
type TestSuite struct{}

func (s *TestSuite) TestPostfixToPrefix (c *C) {
	examples := map[string]string {
		"2 45 + 6 + 24 +": "+ + + 2 45 6 24",
		"60000321 764 +": "+ 60000321 764",
		"26 89 13 +": "too many operands",
		"98 4 * 234 12 * +": "+ * 98 4 * 234 12",
		"hello expression": "invalid input expression",
		"9.007 765.9999994 + 56 ^ /": "too many operators",
		"": "invalid input expression",
		"4 2 - 3 * 5+": "+ 5 * - 4 2 3",
	}

	for postfix, expected := range examples {
		res, err := PostfixToPrefix(postfix)
		if err != nil {
			c.Assert(err, ErrorMatches, expected)
		} else {
			c.Assert(res, Equals, expected)
		}
	}
}

func ExamplePostfixToPrefix() {
	res, err := PostfixToPrefix("34 5 +")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(res)
	}

	// Output:
	// + 34 5
}
