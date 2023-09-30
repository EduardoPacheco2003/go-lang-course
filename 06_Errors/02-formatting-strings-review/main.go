package main

import (
	"fmt"
)

/*
FORMATTING STRINGS REVIEW
A convenient way to format strings in Go is by using the standard library's fmt.Sprintf() function. It's a string interpolation function, similar to JavaScript's built-in template literals. The %v substring uses the type's default formatting, which is often what you want.

DEFAULT VALUES
EXAMPLE: --------------------------------------
		const name = "Kim"
		const age = 22
		s := fmt.Sprintf("%v is %v years old.", name, age)
		// s = "Kim is 22 years old."

The equivalent JavaScript code:
		const name = 'Kim'
		const age = 22
		s = `${name} is ${age} years old.`
		// s = "Kim is 22 years old."
*/

/*
ASSIGNMENT
We need better error logs for our backend developers to help them debug their code.

Complete the getSMSErrorString() function. It should return a string with this format:

SMS that costs $COST to be sent to 'RECIPIENT' can not be sent
\
COST is the cost of the SMS, always showing the price formatted to 2 decimal places.
RECIPIENT is the stringified representation of the recipient's phone number
Be sure to include the $ symbol and the single quotes
*/

func getSMSErrorString(cost float64, recipient string) string {
	return fmt.Sprintf("SMS that costs $%.2f to be sent to '%s' can not be sent", cost, recipient)
}

// don't edit below this line

func test(cost float64, recipient string) {
	s := getSMSErrorString(cost, recipient)
	fmt.Println(s)
	fmt.Println("====================================")
}

func main() {
	test(1.4, "+1 (435) 555 0923")
	test(2.1, "+2 (702) 555 3452")
	test(32.1, "+1 (801) 555 7456")
	test(14.4, "+1 (234) 555 6545")
}
