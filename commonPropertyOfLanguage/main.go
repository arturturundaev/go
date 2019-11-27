package main

import "fmt"

/**
Рассматриваю общие характеристики языка
1. Переменные, константы
2. Циклы
*/
func main() {
	var stringVariable = "Простая строка"
	var firstIntegerVariable, secondIntegerVariable = 12, 13
	var boolVariable = true
	shortString := "Строка с короткой инициализацией"
	const constShort = "Моя константа"
	arrayShort := [5]int{1, 2, 3, 4, 5}

	fmt.Println(stringVariable, firstIntegerVariable, secondIntegerVariable, boolVariable, shortString, constShort, arrayShort)

	for i := 0; i < len(arrayShort); i++ {
		switch i {
		case 1:
			fmt.Println("Перехватили цифру 1")
			break

		case 3:
			fmt.Println("Перехватили цифру 3")
			continue
		default:
			if i == 1 {
				fmt.Println("Т.к. break - мы эту строку не увидим")
			}
		}

		if i == 2 {
			fmt.Println("Перехватили цифру 2")
			continue
		}

		fmt.Println(i)
	}
}
