package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	if len(strings.Fields(input)) == 0 {
		return "", fmt.Errorf("input string is empty(contains not character, but whitespace): %w", errorEmptyInput)
	}

	str :=
		// 8. remove leading "+" in order not to make unnecessary splitting
		strings.TrimLeft(
			// 7. replace "-" with "+-" to further split the input string by "+" sign
			strings.ReplaceAll(
				// 6. replace duplicate "++" with "+"
				strings.ReplaceAll(
					// 5. replace "+-" with "-"
					strings.ReplaceAll(
						// 4. replace "-+" with "-"
						strings.ReplaceAll(
							// 3. replace duplicate "--" with "+"
							strings.ReplaceAll(
								// 2. remove all remaining spaces inside the input string
								strings.ReplaceAll(
									// 1. remove all leading and trailing white space
									strings.TrimSpace(input),
									" ", ""),
								"--", "+"),
							"-+", "-"),
						"+-", "-"),
					"++", "+"),
				"-", "+-"),
			"+")

	operands := strings.Split(str, "+")

	if len(operands) == 1 || len(operands) > 2 {
		return "", fmt.Errorf("an expression contains one or greater than two operands: %w", errorNotTwoOperands)
	}

	res := 0
	for i := 0; i < len(operands); i++ {
		d, e := strconv.Atoi(operands[i])
		if e != nil {
			return "", fmt.Errorf("the input expression is not valid(contains characters, that are not numbers, +, - or whitespace): %w", e)
		}
		res += d
	}
	return fmt.Sprint(res), nil
}
