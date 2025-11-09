package main

import (
    "fmt"
)

func main() {
    fmt.Println("***** Variables ******")
    fmt.Println("Variable numerico positivo y negativo ")
    var myIntVar int  = 12
    fmt.Println("My Variable Int Is ",myIntVar)
    fmt.Println("My direccion de memoria ",&myIntVar)

    fmt.Println("Variable nuerico positivo ")
    var myUintVar uint = 12
    fmt.Println("Variable Uint is", myUintVar)

    fmt.Println("Varibale String")
    var myStringVar string = "My string var"
    fmt.Println("Variable String ", myStringVar)

    fmt.Println("Valiable Boleana")
    var myBoolVar bool = true
    fmt.Println("Variable Bool ",myBoolVar)

    myIntVar2 := 30
    fmt.Println("Variable seteada tipo de dato", myIntVar2)
}
