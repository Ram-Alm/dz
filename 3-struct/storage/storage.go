package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"struct-dz/bins"
)

func SaveBin(bins bins.BinList) {
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

func ReadBin(bin *bins.BinList) (bins.BinList, error) {
	data, err := os.ReadFile("bins.json")
	if err != nil {
		return nil, errors.New("Не удалось прочитать файл.")
	}
	err = json.Unmarshal(data, &bin)
	if err != nil {
		return nil, errors.New("Не удалось преобразовать JSON данные в слайс структур.")
	}
	return *bin, nil
}
