package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"struct-dz/bins"
	"struct-dz/file"
)

func SaveBin(bins bins.BinList) error {
	content, err := json.Marshal(bins)
	if err != nil {
		return fmt.Errorf("Не удалось преобразовать данные в JSON. %w", err)
	}

	file, err := os.Create("bins.json")
	if err != nil {
		return fmt.Errorf("Не удалось создать файл. %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(string(content))
	if err != nil {
		return fmt.Errorf("Не удалось записать данные в файл. %w", err)
	}
	fmt.Println("Данные bin сохранены в файл bins.json")
	return nil
}

func ReadBin(name string) (bins.BinList, error) {
	var bin bins.BinList
	isJson := file.CheckFile(name)
	if !isJson {
		return nil, errors.New("Файл не является JSON")
	}
	data, err := file.ReadFile(name)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &bin)
	if err != nil {
		return nil, errors.New("Не удалось преобразовать JSON данные в слайс структур.")
	}
	return bin, nil
}
