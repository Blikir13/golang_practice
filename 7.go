package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Print("Введите положительное целое N: ")

	if _, err := fmt.Scan(&n); err != nil || n <= 0 {
		fail("N должно быть положительным целым числом")
	}

	var sum float64
	
	for i := 0; i < n; i++ {
		x := i + 1
		if x%2 == 0 && x%3 == 0 {
			sum += float64(x)
		}
	}

	fmt.Println("Сумма:", sum)
}

func fail(message string) {
	fmt.Fprintln(os.Stderr, "Ошибка:", message)
	os.Exit(1)
}
