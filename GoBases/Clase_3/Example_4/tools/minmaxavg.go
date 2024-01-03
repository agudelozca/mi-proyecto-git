package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

// Función para realizar cálculos (mínimo, promedio o máximo)
func operation(opType string) (func(...int) float64, error) {
	switch opType {
	case minimum:
		return func(nums ...int) float64 {
			if len(nums) == 0 {
				return 0
			}
			minValue := nums[0]
			for _, num := range nums[1:] {
				if num < minValue {
					minValue = num
				}
			}
			return float64(minValue)
		}, nil
	case average:
		return func(nums ...int) float64 {
			if len(nums) == 0 {
				return 0
			}
			sum := 0
			for _, num := range nums {
				sum += num
			}
			return float64(sum) / float64(len(nums))
		}, nil
	case maximum:
		return func(nums ...int) float64 {
			if len(nums) == 0 {
				return 0
			}
			maxValue := nums[0]
			for _, num := range nums[1:] {
				if num > maxValue {
					maxValue = num
				}
			}
			return float64(maxValue)
		}, nil
	default:
		return nil, errors.New("Operación no definida")
	}
}

func main() {
	// Ejemplos de uso
	minFunc, err := operation(minimum)
	averageFunc, err := operation(average)
	maxFunc, err := operation(maximum)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Calcular valores
	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	// Imprimir resultados
	fmt.Printf("Valor mínimo: %.2f\n", minValue)
	fmt.Printf("Valor promedio: %.2f\n", averageValue)
	fmt.Printf("Valor máximo: %.2f\n", maxValue)
}
