package main

import "fmt";

func main() {
	//Estrcutura de variables
	/*
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128

	*/
	var nombre1 string = "Juan";
	var nombre2 string = "Alex";
	
	
	//Forma 2 de declarar variables
	edad := 19;

	fmt.Println(nombre1, nombre2, edad);
	
}