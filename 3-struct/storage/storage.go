package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"struct-dz/bins"
)

func saveBin(bins bins.BinList) {
	content, err := json.Marshal(bins)
	if err != nil {
		fmt.Println("Не удалось преобразовать данные в JSON.", err)
		return
	}

	file, err := os.Create("bins.json")
	if err != nil {
		fmt.Println("Не удалось создать файл.", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(string(content))
	if err != nil {
		fmt.Println("Не удалось записать данные в файл.", err)
		return
	}
	fmt.Println("Данные bin сохранены в файл bins.json")
}

func readBin(bins bins.BinList) {
	data, err := os.ReadFile("bins.json")
	if err != nil {
		fmt.Println("Не удалось прочитать файл.", err)
	}
	err = json.Unmarshal(data, bins)
	if err != nil {
		fmt.Println("Не удалось преобразовать JSON данные в слайс структур.", err)
	}
	fmt.Println(bins)
}
