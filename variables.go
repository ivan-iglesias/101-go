package main

import (
	"fmt"
	"strings"
)

func numbers() {
	// Infiriendo el tipo de valor
	entero := 10                  // int
	decimal := 3.14               // float64
	suma := entero + int(decimal) // int

	fmt.Println(suma)
}

func text() {
	mensaje := "Hola, "
	concatenado := mensaje + "Iván"
	enMayuscula := strings.ToUpper(concatenado)

	fmt.Println(enMayuscula)
}

func arrays() {
	// Tamaño explícito y fijo
	arrayFijo1 := [3]int{1, 2, 3}

	// Inferir automáticamente el tamaño del array a partir del número de elementos proporcionados durante la declaración.
	arrayFijo2 := [...]int{1, 2, 3, 4, 5}

	fmt.Println(arrayFijo1, arrayFijo2)
}

func slices() {
	// Tamaño es dinámico y no se especifica ni infiere. Se inicializa con N elementos, pero puede crecer o reducirse.
	sliceVariable := []int{4, 5, 6}
	sliceVariable = append(sliceVariable, 7)
	sliceVariable = append(sliceVariable, 3, 4)

	fmt.Println(sliceVariable)
	fmt.Println(len(sliceVariable)) // Longitud, cuanto elementos hay en el array
	fmt.Println(cap(sliceVariable)) // Capacidad, cual es la capcacidad del array

	// slices.Equal(sliceVariable1, sliceVariable2) -> Comprar dos slices
}
