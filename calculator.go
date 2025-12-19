package main

import (
	"fmt"
)

func Add(a, b float64) float64 {
	return a + b
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) float64 {
	return a / b
}

func main() {
	for {
		var num1, num2 float64
		var operator string

		fmt.Print("Введите первое число: ")
		fmt.Scanln(&num1)

		fmt.Print("Введите оператор (+, -, *, /): ")
		fmt.Scanln(&operator)

		fmt.Print("Введите второе число: ")
		fmt.Scanln(&num2)

		var result float64

		switch operator {
		case "+":
			result = Add(num1, num2)

		case "-":
			result = Subtract(num1, num2)

		case "*":
			result = Multiply(num1, num2)

		case "/":
			if num2 == 0 {
				fmt.Println("Ошибка: деление на ноль!")
				continue
			}
			result = Divide(num1, num2)

		default:
			fmt.Println("Ошибка: неизвестный оператор!")
			continue
		}

		fmt.Println("Результат:", result)
		fmt.Println()
	}
}
