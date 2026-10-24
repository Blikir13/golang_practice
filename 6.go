package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Print("Введите количество чисел: ")

	if _, err := fmt.Scan(&n); err != nil || n < 0 {
		fail("Количество должно быть неотрицательным целым числом")
	}

	fmt.Println("Введите целые числа:")
	positive, negative, zeros := 0, 0, 0

	for i := 0; i < n; i++ {
		var x int
		if _, err := fmt.Scan(&x); err != nil {
			fail("Некорректное число")
		}
		if x > 0 {
			positive++
		} else if x < 0 {
			negative++
		} else {
			zeros++
		}
	}
	
	fmt.Println("Положительных:", positive)
	fmt.Println("Отрицательных:", negative)
	fmt.Println("Нулей:", zeros)
}

func fail(message string) {
	fmt.Fprintln(os.Stderr, "Ошибка:", message)
	os.Exit(1)
}
