package main

import "fmt"

/*
ASSIGNMENT
At Textio, a "user" struct represents an account holder, and a "sender" is just a "user" with some "sender" specific data. A "sender" is a user that has a rateLimit field that tells us how many messages they are allowed to send.

Fix the system by using an embedded struct as expected by the test code.
*/

type sender struct {
	user
	rateLimit int
}

type user struct {
	name   string
	number int
}

// don't edit below this line

func test(s sender) {
	fmt.Println("Sender name:", s.name)
	fmt.Println("Sender number:", s.number)
	fmt.Println("Sender rateLimit:", s.rateLimit)
	fmt.Println("====================================")
}

func main() {
	test(sender{
		rateLimit: 10000,
		user: user{
			name:   "Deborah",
			number: 18055558790,
		},
	})
	test(sender{
		rateLimit: 5000,
		user: user{
			name:   "Sarah",
			number: 19055558790,
		},
	})
	test(sender{
		rateLimit: 1000,
		user: user{
			name:   "Sally",
			number: 19055558790,
		},
	})
}
