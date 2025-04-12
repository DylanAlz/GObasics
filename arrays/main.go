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


arraywithvalues := [6]int{1, 2, 3}
fmt.Println(arraywithvalues)

array[2] = 5
fmt.Println(array) //5

fmt.Println("Size array 1: ", len(array)) 
fmt.Println("Size array 2: ", len(arraywithvalues)) 

//Slices, son como los vectores, pero los slice no tieen un tamaño fijo y este puede cambiar en tiempo de ejecucion

 var slice []int //solo se inicializa, no se le asigna un tamaño
 fmt.Printf("size: %d ,  value: %v\n",len(slice) ,slice) 

slice = append(slice, 1, 2, 3, 4, 5)
 fmt.Printf("size: %d ,  value: %v\n",len(slice) ,slice) 

mySlice := arraywithvalues[2:4] //capturará los valores de el array de la posicion 2 a la 4 sin incluir la 4
 fmt.Println(mySlice) 

 //ambos ocupan la misma direccion en memoria
fmt.Println(&arraywithvalues[2]) 
fmt.Println(&mySlice[0]) 

fmt.Println(arraywithvalues[:4]) //captura desde el inicio hasta la posicion 4 sin incluir la 4

//otra manera de instanciar un slice
mySlice2 := make([]int, 10) //crea un slice de 10 items
fmt.Println(mySlice2) 
}