package main

import "fmt"

func main() {
	// --- 1. Operadores Aritméticos ---
	fmt.Println("=== 1. Operadores Aritméticos (Números) ===")
	
	a := 10
	b := 3
	c := 3.0
	
	fmt.Printf("a = %d, b = %d, c = %.1f\n", a, b, c)

	fmt.Printf("Suma (a + b): %d\n", a + b)           // 13
	fmt.Printf("Resta (a - b): %d\n", a - b)          // 7
	fmt.Printf("Multiplicación (a * b): %d\n", a * b) // 30
	
	// División Entera (ignora el residuo)
	fmt.Printf("División (a / b): %d (División Entera)\n", a / b) // 3
	
	// Módulo (Residuo de la división)
	fmt.Printf("Módulo (a %% b): %d\n", a % b)           // 1
	
	// División de Flotantes (requiere que al menos un operando sea flotante)
	fmt.Printf("División Flotante (float64(a) / c): %.2f\n", float64(a) / c) // 3.33

	// Operadores unarios
	a++ // Incremento: a = a + 1
	fmt.Printf("Incremento (a++): %d\n", a) // 11
	b-- // Decremento: b = b - 1
	fmt.Printf("Decremento (b--): %d\n", b) // 2

	// --- 2. Operadores Relacionales (Comparación) ---
	fmt.Println("\n=== 2. Operadores Relacionales (Booleanos) ===")
	
	x := 5
	y := 8
	
	fmt.Printf("x = %d, y = %d\n", x, y)
	
	fmt.Printf("Igualdad (x == y): %t\n", x == y)   // false
	fmt.Printf("No Igualdad (x != y): %t\n", x != y) // true
	fmt.Printf("Mayor Que (x > y): %t\n", x > y)     // false
	fmt.Printf("Menor Que (x < y): %t\n", x < y)     // true
	fmt.Printf("Mayor o Igual (x >= 5): %t\n", x >= 5) // true
	fmt.Printf("Menor o Igual (y <= 8): %t\n", y <= 8) // true

	// --- 3. Operadores Lógicos ---
	fmt.Println("\n=== 3. Operadores Lógicos (Booleanos) ===")
	
	p := true
	q := false
	
	fmt.Printf("p = %t, q = %t\n", p, q)
	
	// AND lógico (ambos deben ser true)
	fmt.Printf("AND lógico (p && q): %t\n", p && q) // false
	
	// OR lógico (al menos uno debe ser true)
	fmt.Printf("OR lógico (p || q): %t\n", p || q)   // true
	
	// NOT lógico (niega el valor)
	fmt.Printf("NOT lógico (!p): %t\n", !p)         // false
	
	// --- 4. Operadores de Asignación ---
	fmt.Println("\n=== 4. Operadores de Asignación Compuesta ===")
	
	i := 20
	fmt.Printf("Inicial: i = %d\n", i) // 20

	i += 5 // i = i + 5
	fmt.Printf("i += 5: %d\n", i)      // 25

	i -= 10 // i = i - 10
	fmt.Printf("i -= 10: %d\n", i)     // 15

	i *= 2 // i = i * 2
	fmt.Printf("i *= 2: %d\n", i)      // 30

	i /= 4 // i = i / 4
	fmt.Printf("i /= 4: %d (División Entera)\n", i) // 7

	i %= 3 // i = i % 3
	fmt.Printf("i %%= 3: %d\n", i)      // 1

	// --- 5. Operadores a Nivel de Bits (Bitwise) ---
	// Más especializados, se usan para manipulación de bits
	fmt.Println("\n=== 5. Operadores a Nivel de Bits (Bitwise) ===")
	
	m := 6 // Binario: 0110
	n := 3 // Binario: 0011
	
	fmt.Printf("m = %d (0110), n = %d (0011)\n", m, n)

	// AND Bitwise (&): 0110 & 0011 = 0010 (2)
	fmt.Printf("AND Bitwise (m & n): %d\n", m & n) 
	
	// OR Bitwise (|): 0110 | 0011 = 0111 (7)
	fmt.Printf("OR Bitwise (m | n): %d\n", m | n)
	
	// XOR Bitwise (^): 0110 ^ 0011 = 0101 (5)
	fmt.Printf("XOR Bitwise (m ^ n): %d\n", m ^ n)
	
	// Desplazamiento a la izquierda (<<): 0110 << 1 = 1100 (12)
	fmt.Printf("Desplazamiento a la Izquierda (m << 1): %d\n", m << 1)
}