package main

import "fmt"

func main() {
	const usdToEur = 0.86
	const usdToRub = 86.3

	eurToRub := (1 / usdToEur) * usdToRub

	fmt.Printf("Рассчитанный курс EUR к RUB: %.2f\n", eurToRub)
}
