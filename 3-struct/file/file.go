package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("Не удалось прочитать файл. Ошибка:", err)
	}

	isJson := checkFile(name)
	if !isJson {
		return nil, errors.New("Файл не json!")
	}
	return data, nil
}

func checkFile(name string) bool {
	ext := filepath.Ext(name)
	if strings.EqualFold(ext, ".json") {
		return true
	}
	return false
}
