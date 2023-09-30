package main

import (
	"fmt"
)

/*
ASSIGNMENT
One of our clients likes us to send text messages reminding users of life events coming up.

Fix the bug by using named return values in the function signature so the code will compile and run as intended.
*/

func yearsUntilEvents(age int) (yearsUntilAdult int, yearsUntilDrinking int, yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	};
	return;
}

// don't edit below this line

func test(age int) {
	fmt.Println("Age:", age)
	yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental := yearsUntilEvents(age)
	fmt.Println("You are an adult in", yearsUntilAdult, "years")
	fmt.Println("You can drink in", yearsUntilDrinking, "years")
	fmt.Println("You can rent a car in", yearsUntilCarRental, "years")
	fmt.Println("====================================")
}

func main() {
	// continueProgram := "";
	test(4)
	test(10)
	test(22)
	test(35)
	// fmt.Println("Press enter to continue...");
	// fmt.Scanf("%s", continueProgram);
}
