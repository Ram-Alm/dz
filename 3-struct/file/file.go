package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func readFile(name string) {
	ext := filepath.Ext(name)
	if strings.EqualFold(ext, ".json") {
		data, err := os.ReadFile(name)
		if err != nil {
			fmt.Println("Не удалось прочитать файл. Ошибка:", err)
		}
		fmt.Println(string(data))
	} else {
		fmt.Println("Не поддерживаемый тип файла.")
		return
	}
}
