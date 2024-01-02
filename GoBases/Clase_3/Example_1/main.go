package main

import "fmt"

/* Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar
 el sueldo, para cumplir el objetivo es necesario crear una función que devuelva el impuesto de
un salario.
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 % del sueldo
 y si gana más de $150.000 se le descontará además un 10 % (27% en total).*/

func main() {
	impuesto := SalaryImpuesto(75000)
	fmt.Printf("Para un salario de 75000, el impuesto calculado es: $%.2f\n", impuesto)

}

func SalaryImpuesto(salary float64) float64 {
	var impuesto float64
	if salary > 50000 {
		impuesto = salary * 0.17
	} else if salary > 150000 {
		impuesto = salary * 0.27
	}
	return impuesto
}
