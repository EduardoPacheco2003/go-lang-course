package main
import "fmt"

/*
MULTIPLE PARAMETERS
When multiple arguments are of the same type, and are next to each other in the function signature, the type only needs to be declared after the last argument.

Here are some examples:

func addToDatabase(hp, damage int) {
  // ...
}
Copy icon
func addToDatabase(hp, damage int, name string) {
  // ?
}
Copy icon
func addToDatabase(hp, damage int, name string, level int) {
  // ?
*/

func main(){
	addToDatabase(100, 20);
	addToDatabase2(100, 20, "Goblin");
	addToDatabase3(100, 20, "Goblin", 1);
}

func addToDatabase(hp, damage int) {
	fmt.Printf("\nAdded hp: %d, damage: %d", hp, damage);
  return;
}

func addToDatabase2(hp, damage int, name string) {
  fmt.Printf("\nAdded hp: %d, damage: %d, name: %s", hp, damage, name)
}

func addToDatabase3(hp, damage int, name string, level int) {
  fmt.Printf("\nAdded hp: %d, damage: %d, name: %s, level: %d", hp, damage, name, level);
}
