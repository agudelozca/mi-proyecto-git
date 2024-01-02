package main

import "fmt"

func main() {
	promedio := calcularPromedio(5, 5, 5, 5, 5, 5, 4, 4, 4, 4)
	if promedio == 0 {
		fmt.Println("No se realizo el calculo del promedio debido a error en el input")
	} else {
		fmt.Println("El promedio de notas es: ", promedio)
	}
}

func calcularPromedio(notas ...float64) float64 {
	var sumaTotal float64
	var prom float64
	cantidad := len(notas)
	for _, i := range notas {
		if i < 0 {
			fmt.Println("Se ingresaron calificaciones negativas, ingresar solo mayores o igual 0")
			return 0
		}
		sumaTotal += i
	}
	prom = sumaTotal / float64(cantidad)

	return prom
}
