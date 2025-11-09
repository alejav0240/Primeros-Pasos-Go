package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Inicializamos el generador de números aleatorios
	// Esto es para que rand.Intn genere resultados diferentes cada vez.
	rand.Seed(time.Now().UnixNano())

	// =========================================================
	// 1. Sentencia IF-ELSE IF-ELSE
	// =========================================================
	fmt.Println("=== 1. Sentencia IF-ELSE IF-ELSE ===")
	
	puntuacion := 75

	if puntuacion >= 90 {
		fmt.Printf("Puntuación %d: ¡Sobresaliente!\n", puntuacion)
	} else if puntuacion >= 70 {
		fmt.Printf("Puntuación %d: Notable. Buen trabajo.\n", puntuacion)
	} else if puntuacion >= 50 {
		fmt.Printf("Puntuación %d: Aprobado. Justo.\n", puntuacion)
	} else {
		fmt.Printf("Puntuación %d: Suspenso. Necesita mejorar.\n", puntuacion)
	}

	// =========================================================
	// 2. Sentencia IF con Inicializador (Short Statement)
	// =========================================================
	// La variable 'num' se declara y evalúa SÓLO dentro del scope del if.
	// Esto ayuda a mantener el código limpio y la variable 'num' acotada.
	fmt.Println("\n=== 2. IF con Inicializador (num) ===")

	if num := rand.Intn(10); num > 7 {
		fmt.Printf("¡Alto! El número aleatorio (%d) es grande.\n", num)
	} else if num > 3 {
		fmt.Printf("El número aleatorio (%d) es medio.\n", num)
	} else {
		fmt.Printf("¡Bajo! El número aleatorio (%d) es pequeño.\n", num)
	}
	
	// ERROR: Intentar acceder a 'num' aquí fallará, porque solo existe dentro del 'if'.
	// fmt.Println(num) 

	// =========================================================
	// 3. Sentencia SWITCH (Estructura de ramificación múltiple)
	// =========================================================
	fmt.Println("\n=== 3. Sentencia SWITCH (Valor específico) ===")

	dia := "miércoles"

	switch dia {
	case "lunes", "martes", "miércoles", "jueves", "viernes":
		fmt.Printf("%s: Es un día laboral.\n", dia)
	case "sábado", "domingo":
		fmt.Printf("%s: ¡Es fin de semana!\n", dia)
	default:
		fmt.Printf("Error: \"%s\" no es un día válido.\n", dia)
	}

	// 💡 Nota clave de Go: A diferencia de C o Java, Go no necesita
	// la palabra clave 'break' después de cada 'case'. La ejecución se detiene automáticamente.


	// =========================================================
	// 4. Sentencia SWITCH sin Expresión (Actúa como IF-ELSE IF)
	// =========================================================
	// El switch compara booleanos (true) de forma implícita,
	// permitiendo usar expresiones en los 'case'.
	fmt.Println("\n=== 4. SWITCH sin Expresión (Rango) ===")

	edad := 19

	switch {
	case edad < 18:
		fmt.Printf("Edad %d: Es menor de edad.\n", edad)
	case edad >= 18 && edad < 65:
		fmt.Printf("Edad %d: Es un adulto en edad laboral.\n", edad)
	default:
		fmt.Printf("Edad %d: Es una persona mayor.\n", edad)
	}
}