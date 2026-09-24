package main

import (
	"fmt"
	"strings"
)

// Константы для курсов валют (базовая валюта - USD)
const (
	usdToEur = 0.86
	usdToRub = 86.3
)

func main() {
	fmt.Println("=== Добро пожаловать в Калькулятор Валют ===")

	// 1. Шаг: Ввод исходной валюты
	currency1 := getCurrency("Введите исходную валюту")

	// 2. Шаг: Ввод суммы
	amount := getAmount()

	// 3. Шаг: Ввод целевой валюты
	currency2 := getCurrency("Введите целевую валюту")

	// 4. Шаг: Расчет и вывод результата
	calculate(amount, currency1, currency2)
}

// Функция ввода и проверки валюты
func getCurrency(prompt string) string {
	var currency string
	for {
		fmt.Printf("%s (доступны: USD, EUR, RUB): ", prompt)
		_, err := fmt.Scanln(&currency)

		// Переводим в верхний регистр, чтобы ввод "usd" или "Usd" тоже работал
		currency = strings.ToUpper(strings.TrimSpace(currency))

		if err != nil || (currency != "USD" && currency != "EUR" && currency != "RUB") {
			fmt.Println("Ошибка: некорректная валюта. Пожалуйста, выберите из списка.")
			continue
		}
		return currency
	}
}

// Функция ввода и проверки числа (суммы)
func getAmount() float64 {
	var amount float64
	for {
		fmt.Print("Введите сумму для конвертации: ")
		_, err := fmt.Scanln(&amount)

		if err != nil || amount < 0 {
			fmt.Println("Ошибка: введите корректное положительное число.")
			// Очищаем буфер ввода в случае ошибки чтения строки
			var discard string
			fmt.Scanln(&discard)
			continue
		}
		return amount
	}
}

// Функция расчета и вывода итога
func calculate(amount float64, from, to string) {
	// Сначала переводим любую исходную валюту в промежуточный USD
	var amountInUSD float64

	switch from {
	case "USD":
		amountInUSD = amount
	case "EUR":
		amountInUSD = amount / usdToEur
	case "RUB":
		amountInUSD = amount / usdToRub
	}

	// Затем переводим из USD в целевую валюту
	var result float64

	switch to {
	case "USD":
		result = amountInUSD
	case "EUR":
		result = amountInUSD * usdToEur
	case "RUB":
		result = amountInUSD * usdToRub
	}

	fmt.Printf("\nРезультат: %.2f %s = %.2f %s\n", amount, from, result, to)
}

