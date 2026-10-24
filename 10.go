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
	var sum float64
	var min, max, even, odd int

	for i := 0; i < n; i++ {
		var x int
		if _, err := fmt.Scan(&x); err != nil {
			fail("Некорректное число")
		}
		sum += float64(x)
		if i == 0 || x < min {
			min = x
		}
		if i == 0 || x > max {
			max = x
		}
		if x%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	
	fmt.Println("Сумма:", sum)
	fmt.Println("Среднее:", sum/float64(n))
	fmt.Println("Максимум:", max)
	fmt.Println("Минимум:", min)
	fmt.Println("Чётных:", even)
	fmt.Println("Нечётных:", odd)
}

func fail(message string) {
	fmt.Fprintln(os.Stderr, "Ошибка:", message)
	os.Exit(1)
}
