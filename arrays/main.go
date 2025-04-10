package main

import (
	"fmt"
)

func main() {
var intVar int
fmt.Printf("the value of intvar is: %v\n", intVar) 

//vector
var array [6]int
fmt.Println(array)


arraywithvalues := [3]int{1, 2, 3}
fmt.Println(arraywithvalues)

array[2] = 5
fmt.Println(array) //5

fmt.Println("Size array 1: ", len(array)) 
fmt.Println("Size array 2: ", len(arraywithvalues)) 

//Slices, son como los vectores, pero los slice no tieen un tamaño fijo y este puede cambiar en tiempo de ejecucion

 var slice []int //solo se inicializa, no se le asigna un tamaño
 fmt.Printf("size: %d ,  value: %v\n",len(slice) ,slice) //el valor de slice es: []

}