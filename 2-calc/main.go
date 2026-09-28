package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

func main() {
	operation := getOperation()
	numbers := getNumbers()
	result := calculate(operation, numbers)
	fmt.Println(result)
}

func calculate(opearatin string, numbers []int) int {
	var result int
	switch opearatin {
	case "SUM":
		for _, value := range numbers {
			result += value
		}
		return result
	case "AVG":
		for _, value := range numbers {
			result += value
		}
		return result / len(numbers)
	case "MED":
		slices.Sort(numbers)
		if len(numbers)%2 != 0 {
			return numbers[len(numbers)/2]
		} else if len(numbers)%2 == 0 {
			return (numbers[len(numbers)/2-1] + numbers[len(numbers)/2]) / 2
		}
	}
	return 0
}

func getOperation() string {
	var operation string

	fmt.Print("Введите тип операции (AVG, SUM, MED): ")
	for {
		fmt.Scan(&operation)
		if operation == "AVG" || operation == "SUM" || operation == "MED" {
			return operation
		}
	}
}

func getNumbers() []int {
	var inputNumbers string
	nums := make([]int, 0)
	fmt.Print("Введите числа через запятую (пример 1, 2, 3): ")
	fmt.Scan(&inputNumbers)

	numbers := strings.Split(strings.TrimSpace(inputNumbers), ",")

	for _, value := range numbers {
		num, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}
	return nums
}
