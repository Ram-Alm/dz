package file

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type JsonDb struct {
	fileName string
}

func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		fileName: name,
	}
}

func (db *JsonDb) Read() ([]byte, error) {
	data, err := os.ReadFile(db.fileName)
	if err != nil {
		return nil, errors.New("Не удалось прочитать файл.")
	}
	return data, nil
}

func (db *JsonDb) Write(b []byte) error {
	file, err := os.Create(db.fileName)
	if err != nil {
		return err
	}
	_, err = file.WriteString(string(b))
	return err
}

func CheckFile(name string) bool {
	ext := filepath.Ext(name)
	if strings.EqualFold(ext, ".json") {
		return true
	}
	return false
}
