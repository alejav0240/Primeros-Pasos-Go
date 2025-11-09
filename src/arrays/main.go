package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// =========================================================
	// 1. Declaración e Inicialización de Arrays
	// =========================================================
	fmt.Println("=== 1. Declaración y Características ===")

	// Array de tamaño fijo: [5]int. Contendrá 5 enteros.
	var numeros [5]int
	
	// Array inicializado con valores
	frutas := [3]string{"Manzana", "Pera", "Naranja"}

	// Array con inicialización implícita de tamaño (...)
	// El compilador cuenta los elementos: [4]float64
	precios := [...]float64{1.99, 2.50, 4.00, 0.75}

	fmt.Printf("Array 'numeros': %v (Tipo: %T)\n", numeros, numeros)
	fmt.Printf("Array 'frutas': %v (Tipo: %T)\n", frutas, frutas)
	fmt.Printf("Array 'precios': %v (Tipo: %T)\n", precios, precios)
    
	// ---
	fmt.Println("\n=== 2. Tamaño Fijo y Bytes ===")
    
    // El tamaño del array es fijo y se obtiene con len()
	fmt.Printf("Tamaño de 'numeros': %d elementos\n", len(numeros))
	fmt.Printf("Tamaño de 'frutas': %d elementos\n", len(frutas))

    // Tamaño en bytes usando unsafe.Sizeof()
    // [5]int * (8 bytes/int en arquitecturas de 64 bits) = 40 bytes
	fmt.Printf("Tamaño en bytes de [5]int: %d bytes\n", unsafe.Sizeof(numeros)) 
    // [3]string * (16 bytes/string descriptor) = 48 bytes (solo el descriptor)
	fmt.Printf("Tamaño en bytes de [3]string: %d bytes (Descriptor)\n", unsafe.Sizeof(frutas))

	// =========================================================
	// 3. Funciones Esenciales (Acceso y Iteración)
	// =========================================================
	fmt.Println("\n=== 3. Acceso y Modificación ===")

	// Acceso a elementos (índice 0-basado)
	fmt.Printf("Primer elemento de 'frutas': %s\n", frutas[0]) // Manzana
	fmt.Printf("Último elemento de 'frutas': %s\n", frutas[len(frutas)-1]) // Naranja

	// Modificación de un elemento
	numeros[2] = 99
	fmt.Printf("Array 'numeros' modificado: %v\n", numeros)
	
	// ---
	fmt.Println("\n=== 4. Iteración (Uso de 'range') ===")

	// Iteración usando el bucle for...range
	fmt.Println("Precios del inventario:")
	for indice, valor := range precios {
		fmt.Printf("Índice %d: $%.2f\n", indice, valor)
	}
	
	// Ignorar el índice (solo necesitamos el valor)
	sumaPrecios := 0.0
	for _, valor := range precios {
		sumaPrecios += valor
	}
	fmt.Printf("\nSuma total de precios: $%.2f\n", sumaPrecios)
	
	// =========================================================
	// 5. Diferencia Clave: Asignación por Valor
	// =========================================================
	fmt.Println("\n=== 5. Asignación por Valor (Copia) ===")
	
	// Cuando asignas un array a otro, se crea una COPIA COMPLETA.
	a1 := [3]int{1, 2, 3}
	a2 := a1 // Se copia todo el array a2
	
	a2[0] = 99 // Modificamos a2

	fmt.Printf("a1 después de modificar a2: %v\n", a1) // a1 sigue siendo {1 2 3}
	fmt.Printf("a2 modificado: %v\n", a2)           // a2 es {99 2 3}

	// Un array solo se puede asignar a un array del MISMO TIPO
	// [3]int NO es el mismo tipo que [4]int o [3]string
	// var a3 [4]int = a1 // Esto causaría un error de compilación
}