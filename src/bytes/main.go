package main

import "fmt"

func main() {
	// === I. Uso del Tipo 'byte' (uint8) para Caracteres Individuales ===
	fmt.Println("=== I. Uso de 'byte' Individual ===")

	var A byte = 'A' // ASCII 65
	var r byte = 82  // Valor numérico 82 (ASCII de 'R')

	// Impresión del valor numérico (ASCII)
	fmt.Printf("Byte A: Valor numérico: %d, Carácter: %c\n", A, A)
	fmt.Printf("Byte r: Valor numérico: %d, Carácter: %c\n", r, r)

	// Conversión a string (carácter único)
	fmt.Printf("string(A): %s\n", string(A))
	
	// --- Separador ---
	fmt.Println("\n" + "--------------------------------------")

	// === II. Array (Slice) de Bytes para Formar una String ===
	fmt.Println("=== II. Array (Slice) de Bytes ===")
	
	// 1. Declaración de un slice de bytes (código ASCII/UTF-8 de "HOLA")
	// Los valores son: H=72, O=79, L=76, A=65
	sliceBytes := []byte{72, 79, 76, 65} 
	
	fmt.Printf("Slice de Bytes: %v (Tipo: %T)\n", sliceBytes, sliceBytes)

	// 2. Creación de una string a partir del slice de bytes
	// Esta es la forma canónica de convertir un slice de bytes a string en Go.
	palabra := string(sliceBytes) 
	
	fmt.Printf("Slice convertido a String: %s (Tipo: %T)\n", palabra, palabra)

	// 3. Conversión de una string a un slice de bytes
	// Go hace la conversión inversa.
	otraPalabra := "GoLang"
	otraSliceBytes := []byte(otraPalabra)
	
	fmt.Printf("String \"%s\" a Slice de Bytes: %v (Tipo: %T)\n", otraPalabra, otraSliceBytes, otraSliceBytes)
}