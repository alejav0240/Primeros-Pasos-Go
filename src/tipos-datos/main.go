package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// --- Tipos Numéricos Enteros (Con Signo) ---
	fmt.Println("--- Tipos Numéricos Enteros (Con Signo) ---")

	var myInt8Var int8 = 12
	fmt.Printf("int8: Valor: %d, Tipo: %T, Bytes: %d, Bits: %d\n", myInt8Var, myInt8Var, unsafe.Sizeof(myInt8Var), unsafe.Sizeof(myInt8Var)*8)

	var myInt16Var int16 = 12345
	fmt.Printf("int16: Valor: %d, Tipo: %T, Bytes: %d, Bits: %d\n", myInt16Var, myInt16Var, unsafe.Sizeof(myInt16Var), unsafe.Sizeof(myInt16Var)*8)

	var myInt32Var int32 = 1234567890
	fmt.Printf("int32: Valor: %d, Tipo: %T, Bytes: %d, Bits: %d\n", myInt32Var, myInt32Var, unsafe.Sizeof(myInt32Var), unsafe.Sizeof(myInt32Var)*8)

	var myInt64Var int64 = 987654321098765432
	fmt.Printf("int64: Valor: %d, Tipo: %T, Bytes: %d, Bits: %d\n", myInt64Var, myInt64Var, unsafe.Sizeof(myInt64Var), unsafe.Sizeof(myInt64Var)*8)

	// int es dependiente de la arquitectura (32 o 64 bits)
	var myIntVar int = 42
	fmt.Printf("int: Valor: %d, Tipo: %T, Bytes: %d, Bits: %d\n", myIntVar, myIntVar, unsafe.Sizeof(myIntVar), unsafe.Sizeof(myIntVar)*8)

	// --- Tipos Numéricos Enteros (Sin Signo) ---
	fmt.Println("\n--- Tipos Numéricos Enteros (Sin Signo) ---")

	var myUint8Var uint8 = 255 // alias: byte
	fmt.Printf("uint8 (byte): Valor: %d, Tipo: %T, Bytes: %d, Bits: %d\n", myUint8Var, myUint8Var, unsafe.Sizeof(myUint8Var), unsafe.Sizeof(myUint8Var)*8)

	var myUint16Var uint16 = 65535
	fmt.Printf("uint16: Valor: %d, Tipo: %T, Bytes: %d, Bits: %d\n", myUint16Var, myUint16Var, unsafe.Sizeof(myUint16Var), unsafe.Sizeof(myUint16Var)*8)

	var myUint32Var uint32 = 4294967295
	fmt.Printf("uint32: Valor: %d, Tipo: %T, Bytes: %d, Bits: %d\n", myUint32Var, myUint32Var, unsafe.Sizeof(myUint32Var), unsafe.Sizeof(myUint32Var)*8)

	var myUint64Var uint64 = 18446744073709551615
	fmt.Printf("uint64: Valor: %d, Tipo: %T, Bytes: %d, Bits: %d\n", myUint64Var, myUint64Var, unsafe.Sizeof(myUint64Var), unsafe.Sizeof(myUint64Var)*8)

	// uint es dependiente de la arquitectura (32 o 64 bits)
	var myUintVar uint = 88
	fmt.Printf("uint: Valor: %d, Tipo: %T, Bytes: %d, Bits: %d\n", myUintVar, myUintVar, unsafe.Sizeof(myUintVar), unsafe.Sizeof(myUintVar)*8)

	// uintptr almacena punteros, también dependiente de la arquitectura
	var myUintptrVar uintptr = 0xdeadbeef
	fmt.Printf("uintptr: Valor: %x, Tipo: %T, Bytes: %d, Bits: %d\n", myUintptrVar, myUintptrVar, unsafe.Sizeof(myUintptrVar), unsafe.Sizeof(myUintptrVar)*8)

	// rune es un alias para int32 y representa un punto de código Unicode
	var myRuneVar rune = 'G'
	fmt.Printf("rune (int32): Valor: %d (%c), Tipo: %T, Bytes: %d, Bits: %d\n", myRuneVar, myRuneVar, myRuneVar, unsafe.Sizeof(myRuneVar), unsafe.Sizeof(myRuneVar)*8)

	// --- Tipos Numéricos de Coma Flotante ---
	fmt.Println("\n--- Tipos Numéricos de Coma Flotante ---")

	var myFloat32Var float32 = 3.14159
	fmt.Printf("float32: Valor: %f, Tipo: %T, Bytes: %d, Bits: %d\n", myFloat32Var, myFloat32Var, unsafe.Sizeof(myFloat32Var), unsafe.Sizeof(myFloat32Var)*8)

	var myFloat64Var float64 = 3.1415926535
	fmt.Printf("float64: Valor: %f, Tipo: %T, Bytes: %d, Bits: %d\n", myFloat64Var, myFloat64Var, unsafe.Sizeof(myFloat64Var), unsafe.Sizeof(myFloat64Var)*8)

	// --- Tipos Numéricos Complejos ---
	fmt.Println("\n--- Tipos Numéricos Complejos ---")

	var myComplex64Var complex64 = 1 + 2i
	fmt.Printf("complex64: Valor: %v, Tipo: %T, Bytes: %d, Bits: %d\n", myComplex64Var, myComplex64Var, unsafe.Sizeof(myComplex64Var), unsafe.Sizeof(myComplex64Var)*8)

	var myComplex128Var complex128 = 3 + 4i
	fmt.Printf("complex128: Valor: %v, Tipo: %T, Bytes: %d, Bits: %d\n", myComplex128Var, myComplex128Var, unsafe.Sizeof(myComplex128Var), unsafe.Sizeof(myComplex128Var)*8)

	// --- Tipo Booleano ---
	fmt.Println("\n--- Tipo Booleano ---")

	var myBoolVar bool = true
	// El tamaño de bool es 1 byte, pero puede variar por alineación.
	fmt.Printf("bool: Valor: %t, Tipo: %T, Bytes: %d, Bits: %d\n", myBoolVar, myBoolVar, unsafe.Sizeof(myBoolVar), unsafe.Sizeof(myBoolVar)*8)

	// --- Tipo String (Cadena de Caracteres) ---
	fmt.Println("\n--- Tipo String ---")

	var myStringVar string = "Hola, Go!"
	// Una string es una estructura que contiene un puntero a los datos y una longitud.
	fmt.Printf("string: Valor: %s, Tipo: %T, Bytes: %d, Bits: %d\n", myStringVar, myStringVar, unsafe.Sizeof(myStringVar), unsafe.Sizeof(myStringVar)*8)

    var myStringVar2 string = `texto con
    interlineado `
    fmt.Println(myStringVar2)
	fmt.Printf("string: Valor: %s, Tipo: %T, Bytes: %d, Bits: %d\n", myStringVar2, myStringVar2, unsafe.Sizeof(myStringVar2), unsafe.Sizeof(myStringVar2)*8)
	// Nota importante: El tamaño de una string (8 o 16 bytes) es el tamaño de la ESTRUCTURA del header de la string,
	// no el tamaño de los datos que contiene la string.
}