package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var balance float64
	fmt.Print("Введите баланс: ")

	if _, err := fmt.Scan(&balance); err != nil || balance < 0 || math.IsNaN(balance) || math.IsInf(balance, 0) {
		fail("Баланс должен быть конечным неотрицательным числом")
	}

	fmt.Println("1 — Посмотреть баланс\n2 — Пополнить баланс\n3 — Снять средства")
	var action int

	if _, err := fmt.Scan(&action); err != nil {
		fail("Некорректное действие")
	}

	switch action {
	case 1:
	case 2, 3:
		var amount float64
		fmt.Print("Введите сумму: ")
		if _, err := fmt.Scan(&amount); err != nil || amount <= 0 || math.IsNaN(amount) || math.IsInf(amount, 0) {
			fail("Сумма должна быть конечным положительным числом")
		}
		if action == 2 {
			if math.IsInf(balance+amount, 0) {
				fail("Баланс выходит за допустимый диапазон")
			}
			balance += amount
		} else {
			if amount > balance {
				fail("Недостаточно средств")
			}
			balance -= amount
		}
	default:
		fail("Неизвестное действие")
	}
	
	fmt.Printf("Баланс: %.2f\n", balance)
}

func fail(message string) {
	fmt.Fprintln(os.Stderr, "Ошибка:", message)
	os.Exit(1)
}
