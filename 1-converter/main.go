package main

import "fmt"

func calculate(num int, currency1, currency2 string) {

}

func userInput() {
	var num int
	fmt.Scan(&num)
}

func main() {
	const usdToEur = 0.86
	const usdToRub = 86.3

	eurToRub := (1 / usdToEur) * usdToRub

	fmt.Printf("Рассчитанный курс EUR к RUB: %.2f\n", eurToRub)
}
