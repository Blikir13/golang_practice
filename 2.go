package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Print("Введите целое число: ")

	if _, err := fmt.Scan(&n); err != nil {
		fail("Некорректное число")
	}

	if n > 0 {
		fmt.Println("Положительное")
	} else if n < 0 {
		fmt.Println("Отрицательное")
	} else {
		fmt.Println("Ноль")
	}
	if n % 2 == 0 {
		fmt.Println("Чётное")
	} else {
		fmt.Println("Нечётное")
	}

	fmt.Println("Делится на 3:", n % 3 == 0)
	fmt.Println("Делится на 5:", n % 5 == 0)
	fmt.Println("Делится на 3 и на 5:", n % 3 == 0 && n % 5 == 0)
}

func fail(message string) {
	fmt.Fprintln(os.Stderr, "Ошибка:", message)
	os.Exit(1)
}
