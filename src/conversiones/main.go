package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	// =========================================================
	// I. Exploración de Tipos de Datos Primitivos (unsafe.Sizeof)
	// =========================================================
	fmt.Println("=== I. Exploración de Tipos de Datos y Tamaño ===")

	// 1. Enteros con Signo
	var myInt8Var int8 = 12
	var myIntVar int = 42
	fmt.Printf("int8:  Valor: %d, Tipo: %T, Bytes: %d, Bits: %d\n", myInt8Var, myInt8Var, unsafe.Sizeof(myInt8Var), unsafe.Sizeof(myInt8Var)*8)
	fmt.Printf("int:   Valor: %d, Tipo: %T, Bytes: %d, Bits: %d (Arquitectura)\n", myIntVar, myIntVar, unsafe.Sizeof(myIntVar), unsafe.Sizeof(myIntVar)*8)

	// 2. Enteros sin Signo
	var myUint8Var uint8 = 255 // alias: byte
	var myRuneVar rune = 'G' // alias: int32
	fmt.Printf("byte:  Valor: %d, Tipo: %T, Bytes: %d, Bits: %d\n", myUint8Var, myUint8Var, unsafe.Sizeof(myUint8Var), unsafe.Sizeof(myUint8Var)*8)
	fmt.Printf("rune:  Valor: %d, Tipo: %T, Bytes: %d, Bits: %d\n", myRuneVar, myRuneVar, unsafe.Sizeof(myRuneVar), unsafe.Sizeof(myRuneVar)*8)

	// 3. Flotantes y Booleano
	var myFloat64Var float64 = 3.14
	var myBoolVar bool = true
	fmt.Printf("float64: Valor: %f, Tipo: %T, Bytes: %d, Bits: %d\n", myFloat64Var, myFloat64Var, unsafe.Sizeof(myFloat64Var), unsafe.Sizeof(myFloat64Var)*8)
	fmt.Printf("bool:    Valor: %t, Tipo: %T, Bytes: %d, Bits: %d\n", myBoolVar, myBoolVar, unsafe.Sizeof(myBoolVar), unsafe.Sizeof(myBoolVar)*8)

	// 4. String (Tamaño del descriptor, no del contenido)
	var myStringVar string = "GoLang"
	fmt.Printf("string:  Valor: %s, Tipo: %T, Bytes: %d, (Descriptor)\n", myStringVar, myStringVar, unsafe.Sizeof(myStringVar))


	// =========================================================
	// II. Conversión Explícita (Casting)
	// =========================================================
	fmt.Println("\n=== II. Conversión Explícita (Casting) ===")
	
	var entero int = 100
	var flotante float64 = 3.14159

	// int a float64
	f_from_i := float64(entero)
	fmt.Printf("int (%d) a float64: %f (Tipo: %T)\n", entero, f_from_i, f_from_i)

	// float64 a int (Se trunca el decimal)
	i_from_f := int(flotante)
	fmt.Printf("float64 (%f) a int: %d (Tipo: %T) <-- Truncado\n", flotante, i_from_f, i_from_f)

	// int32 (rune) a string
	s_from_r := string(myRuneVar)
	fmt.Printf("rune ('G') a string: %s (Tipo: %T)\n", s_from_r, s_from_r)


	// =========================================================
	// III. Formateo a String (fmt.Sprintf) - Como en la imagen
	// =========================================================
	fmt.Println("\n=== III. Formateo a String (fmt.Sprintf) ===")

	// 1. Float a String con precisión
	floatVar := 33.11
	fmt.Printf("float original: type: %T, value: %f\n", floatVar, floatVar)

	floatStrVar := fmt.Sprintf("%.2f", floatVar)
	fmt.Printf("float a string: type: %T, value: %s\n", floatStrVar, floatStrVar)
	
	// 2. Int a String
	intVar := 22
	fmt.Printf("int original: type: %T, value: %d\n", intVar, intVar)
	
	intStrVar := fmt.Sprintf("%d", intVar)
	fmt.Printf("int a string: type: %T, value: %s\n", intStrVar, intStrVar)


	// =========================================================
	// IV. Parseo/Conversión String ↔ Otros Tipos (strconv)
	// =========================================================
	fmt.Println("\n=== IV. Parseo/Conversión (strconv) ===")

	// --- A. String a Tipo Numérico/Booleano (Parseo) ---
	strNum := "567"
	strFloat := "1.414"
	strBool := "false"

	// Parseo de String a int (Atoi es un atajo a ParseInt)
	i_parsed, err := strconv.Atoi(strNum)
	if err == nil {
		fmt.Printf("string (\"%s\") a int: %d (Tipo: %T)\n", strNum, i_parsed, i_parsed)
	}

	// Parseo de String a float (float64)
	f_parsed, err := strconv.ParseFloat(strFloat, 64)
	if err == nil {
		fmt.Printf("string (\"%s\") a float64: %f (Tipo: %T)\n", strFloat, f_parsed, f_parsed)
	}

	// Parseo de String a bool
	b_parsed, err := strconv.ParseBool(strBool)
	if err == nil {
		fmt.Printf("string (\"%s\") a bool: %t (Tipo: %T)\n", strBool, b_parsed, b_parsed)
	}


	// --- B. Tipo Numérico/Booleano a String (Formateo) ---
	num := 789
	pi := 3.141592
	logico := true

	// Int a String (Itoa es un atajo a FormatInt)
	s_from_num := strconv.Itoa(num)
	fmt.Printf("int (%d) a string (Itoa): %s (Tipo: %T)\n", num, s_from_num, s_from_num)

	// Float a String (FormatFloat)
	// 'g' = formato general, -1 = precisión automática, 64 = float64
	s_from_pi := strconv.FormatFloat(pi, 'g', -1, 64)
	fmt.Printf("float (%f) a string (FormatFloat): %s (Tipo: %T)\n", pi, s_from_pi, s_from_pi)

	// Bool a String
	s_from_logico := strconv.FormatBool(logico)
	fmt.Printf("bool (%t) a string (FormatBool): %s (Tipo: %T)\n", logico, s_from_logico, s_from_logico)
}