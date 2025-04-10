package main

import (
	"fmt"
)

func main() {
	yearsOld := 32
	fmt.Println("Operators")
	fmt.Println(yearsOld > 30)  //retorna true
	fmt.Println(yearsOld < 30)  //retorna false
	fmt.Println(yearsOld >= 40) //retorna false
	fmt.Println(yearsOld <= 32) //retorna true
	fmt.Println(yearsOld == 32) //retorna true

	//OR
	fmt.Println("OR")
	fmt.Println(yearsOld >= 32 || yearsOld <= 30) //retorna true
	fmt.Println(yearsOld >= 32 || yearsOld <= 30) //retorna true

	//AND
	fmt.Println("AND")
	fmt.Println(yearsOld >= 32 && yearsOld <= 30) //retorna false
	fmt.Println(yearsOld >= 32 && yearsOld == 32) //retorna true

	//NOT
	fmt.Println("NOT")
	fmt.Println(true)                                //retorna true
	fmt.Println(!true)                               //retorna false(true)
	fmt.Println(!(yearsOld >= 32 && yearsOld == 32)) //retorna true

	//IF ELSE
	varBool := true
	if varBool {
		fmt.Println("is True")
	} else {
		fmt.Println("is False")
	}

	if value := true; value == true {
		fmt.Println("is True")
	} else if value == false {
		fmt.Println("is False")
	} else {
		fmt.Println("is not True or False")
	}

//Switch
 varNum := 6

 switch varNum {
 case 1:
	fmt.Println("the value is one")
 case 2:
	fmt.Println("the value is two")
 case 6:
	fmt.Println("the value is six")
 default:
	fmt.Println("value not recognized")
 }

//Tarea 1
/*Dentro de nuestro codigo vamos a tener 2 variables definidas, que van a ser:
license: si la persona tiene licencia
age: la edad de la persona

    -Si la persona tiene mas de 15 años y tiene licencia, debemos imprimir un mensaje que diga "Puede seguir avanzando"
    -Si la persona tiene 15 años o menos, o no tiene licencia, debemos imprimir un mensaje que diga "No puede seguir circulando" */

licence := true
age := 12

if age>15 && licence {
	fmt.Println("Puede mantener en circulación")
}else if age <=15 || !licence {
	fmt.Println("No puede seguir circulando")
}



}
