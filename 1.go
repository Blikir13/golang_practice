package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var a, b float64
	var op string

	fmt.Print("Введите два числа и операцию (+, -, *, /): ")
	if _, err := fmt.Scan(&a, &b, &op); err != nil {
		fail("Некорректный ввод")
	}

	if math.IsNaN(a) || math.IsInf(a, 0) || math.IsNaN(b) || math.IsInf(b, 0) {
		fail("Введите конечные числа")
	}

	var result float64

	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			fail("Деление на ноль")
		}
		result = a / b
	default:
		fail("Неизвестная операция")
	}
	
	if math.IsNaN(result) || math.IsInf(result, 0) {
		fail("Результат выходит за допустимый диапазон")
	}

	fmt.Println("Результат:", result)
}

func fail(message string) {
	fmt.Fprintln(os.Stderr, "Ошибка:", message)
	os.Exit(1)
}
