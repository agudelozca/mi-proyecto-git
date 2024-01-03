package main

import (
	"testing"
)

// Función de prueba para calcular el mínimo de calificaciones
func TestCalcularMinimo(t *testing.T) {
	calificaciones := []int{80, 90, 85, 95, 88}
	resultado := calcularMinimo(calificaciones...)
	expected := 80
	if resultado != expected {
		t.Errorf("El mínimo de calificaciones no es correcto. Se esperaba %v, pero se obtuvo %v", expected, resultado)
	}
}

// Función de prueba para calcular el máximo de calificaciones
func TestCalcularMaximo(t *testing.T) {
	calificaciones := []int{80, 90, 85, 95, 88}
	resultado := calcularMaximo(calificaciones...)
	expected := 95
	if resultado != expected {
		t.Errorf("El máximo de calificaciones no es correcto. Se esperaba %v, pero se obtuvo %v", expected, resultado)
	}
}

// Función de prueba para calcular el promedio de calificaciones
func TestCalcularPromedio(t *testing.T) {
	calificaciones := []int{80, 90, 85, 95, 88}
	resultado := calcularPromedio(calificaciones...)
	expected := 87.6
	if resultado != expected {
		t.Errorf("El promedio de calificaciones no es correcto. Se esperaba %v, pero se obtuvo %v", expected, resultado)
	}
}
