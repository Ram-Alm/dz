package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var operations = map[string]func([]float64) float64{
	"SUM": sum,
	"AVG": avg,
	"MED": med,
}

func main() {
	operation := getOperation()
	numbers := getNumbers()
	result := operations[operation](numbers)
	fmt.Println(result)
}

func sum(numbers []float64) float64 {
	var result float64
	for _, value := range numbers {
		result += value
	}
	return result
}

func avg(numbers []float64) float64 {
	var result float64
	for _, value := range numbers {
		result += value
	}
	return result / float64(len(numbers))
}

func med(numbers []float64) float64 {
	slices.Sort(numbers)
	if len(numbers)%2 != 0 {
		return numbers[len(numbers)/2]
	} else if len(numbers)%2 == 0 {
		return (numbers[len(numbers)/2-1] + numbers[len(numbers)/2]) / 2
	}
	return 0
}

func getOperation() string {
	var operation string

	for {
		fmt.Print("Введите тип операции (AVG, SUM, MED): ")
		fmt.Scan(&operation)
		if operation == "AVG" || operation == "SUM" || operation == "MED" {
			return operation
		}
	}
}

func getNumbers() []float64 {
	var inputNumbers string
	nums := make([]float64, 0)
	fmt.Print("Введите числа через запятую (пример 1, 2, 3): ")
	inputNumbers, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	numbers := strings.Split(inputNumbers, ",")

	for _, value := range numbers {
		num, err := strconv.ParseFloat(strings.TrimSpace(value), 64)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}
	return nums
}
