package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Print("Введите количество чисел: ")

	if _, err := fmt.Scan(&n); err != nil || n <= 0 {
		fail("Количество должно быть положительным целым числом")
	}

	fmt.Println("Введите целые числа:")
	var min, max int

	for i := 0; i < n; i++ {
		var x int
		if _, err := fmt.Scan(&x); err != nil {
			fail("Некорректное число")
		}
		if i == 0 || x < min {
			min = x
		}
		if i == 0 || x > max {
			max = x
		}
	}
	
	fmt.Println("Максимум:", max)
	fmt.Println("Минимум:", min)
}

func fail(message string) {
	fmt.Fprintln(os.Stderr, "Ошибка:", message)
	os.Exit(1)
}
