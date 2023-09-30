package main

import "fmt"

//Anonymous structs in Go lang

type car struct {
	Make   string
	Model  string
	Height int
	Width  int
	// Wheel is a field containing an anonymous struct
	Wheel struct {
		Radius   int
		Material string
	}
}

func main() {

	myCar := struct {
		Make  string
		Model string
	}{
		Make:  "tesla",
		Model: "model 3",
	}

	notMyCar := car{
		Make:   "tesla X",
		Model:  "model X",
		Height: 100,
		Width:  50,
		Wheel: struct {
			Radius   int
			Material string
		}{
			Radius:   20,
			Material: "rubber",
		},
	}

	fmt.Println(myCar)
	fmt.Println(notMyCar)

}
