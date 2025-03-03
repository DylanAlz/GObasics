package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var myIntVar int
	myIntVar = 42
	fmt.Println("My variable is: ",  myIntVar)

	var myUnitVar uint 
	myUnitVar = 42
	fmt.Println("My variable is: ",  myUnitVar)

	var myStringVar string
	myStringVar = "Hello, World!, this is my first go program"
	fmt.Println("My variable value is: ",  myStringVar)

	var myBoolVar bool
	myBoolVar = true
	fmt.Println("My variable boolean is: ",  myBoolVar)

	//para buscar una direccion en memoria de mi variable

	fmt.Println("My variable address is: ",  &myIntVar)

	//podemos asignar un valor a una variable sin haberla declarado
	myNotDeclaratedIntVar := 21 //esto es igual a var myNotDeclaratedIntVar int = 21
	fmt.Println("My variable is: ",  myNotDeclaratedIntVar)

	//constante
	const myConstVar int = 21
	fmt.Println("My constant is: ",  myConstVar)

	//podemos asignar espacios de memora a variables
	var int8Var int8 //int8 es un entero de 8 bits, hay de 16, 32, 64
	int8Var = 127
	fmt.Println("My variable is: ",  int8Var)

	//imprimir todos los valores
	fmt.Printf("int default value: %v\n", myIntVar) //%v es el valor de la variable y \n es un salto de linea, ademas está \t  para tabular

	fmt.Printf("type %T, bytes: %d, bits: %d\n", int8Var, unsafe.Sizeof(int8Var), unsafe.Sizeof(int8Var)*8)

	var float32Var float32
	float32Var = 3.14
	fmt.Printf("type %T, bytes: %d, bits: %d\n", float32Var, unsafe.Sizeof(float32Var), unsafe.Sizeof(float32Var)*8)

	var newString string
	newString = "Hello, NEW!"
	fmt.Printf("String: %s\n", newString)
	} 