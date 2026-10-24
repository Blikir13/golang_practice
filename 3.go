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
		sum += float64(i + 1)
	}
	
	fmt.Println("Сумма:", sum)
	fmt.Println("Среднее:", sum/float64(n))
}

func fail(message string) {
	fmt.Fprintln(os.Stderr, "Ошибка:", message)
	os.Exit(1)
}
