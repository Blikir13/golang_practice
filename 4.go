package main

import (
	"fmt"
	"math/big"
	"os"
)

func main() {
	var n int
	fmt.Print("Введите целое число: ")

	if _, err := fmt.Scan(&n); err != nil {
		fail("Некорректное число")
	}
	
	printTable(n)
}

func printTable(n int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d × %d = %s\n", n, i, new(big.Int).Mul(big.NewInt(int64(n)), big.NewInt(int64(i))))
	}
}

func fail(message string) {
	fmt.Fprintln(os.Stderr, "Ошибка:", message)
	os.Exit(1)
}
