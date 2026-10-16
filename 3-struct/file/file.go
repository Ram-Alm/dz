package file

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, errors.New("Не удалось прочитать файл.")
	}
	return data, nil
}

func CheckFile(name string) bool {
	ext := filepath.Ext(name)
	if strings.EqualFold(ext, ".json") {
		return true
	}
	return false
}
