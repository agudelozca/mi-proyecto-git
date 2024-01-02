package main

import "fmt"

func main() {
	salarioEmpleado := calcuSalary(1000, "C")
	fmt.Println("Mi salario es: ", salarioEmpleado)
}

func calcuSalary(minutos int, categoria string) float64 {
	salarioC := 1000
	salarioB := 1500
	salarioA := 3000
	switch {
	case categoria == "C":
		salario := (salarioC * minutos) / 60
		return float64(salario)
	case categoria == "B":
		salario := (salarioB * minutos) / 60
		salario2 := float64(salario) * 1.20
		return salario2
	case categoria == "A":
		salario := (salarioA * minutos) / 60
		salario2 := float64(salario) * 1.50
		return salario2
	default:
		//
	}
	return 0
}
