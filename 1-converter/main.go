package main

import (
	"fmt"
	"strings"
)

var exchangeRates = map[string]float64{
	"USD": 1.0,
	"EUR": 0.86,
	"RUB": 86.3,
}

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

// Функция для проверки наличия валюты в мапе
func checkCurrency(currency string) bool {
	_, ok := exchangeRates[currency]
	return ok
}

// Функция ввода и проверки валюты
func getCurrency(prompt string) string {
	var currency string
	for {
		fmt.Printf("%s (доступны: USD, EUR, RUB): ", prompt)
		_, err := fmt.Scanln(&currency)

		// Переводим в верхний регистр, чтобы ввод "usd" или "Usd" тоже работал
		currency = strings.ToUpper(strings.TrimSpace(currency))

		if err != nil || !checkCurrency(currency) {
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
	amountInUSD := amount / exchangeRates[from]

	// Затем переводим из USD в целевую валюту
	result := amountInUSD * exchangeRates[to]

	fmt.Printf("\nРезультат: %.2f %s = %.2f %s\n", amount, from, result, to)
}
