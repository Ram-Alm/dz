package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"struct-dz/bins"
	"struct-dz/file"
)

type Db interface {
	Read() ([]byte, error)
	Write(b []byte) error
}

type Storage struct {
	db Db
}

func NewStorage(db Db) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) SaveBin(bins bins.BinList) error {
	content, err := json.Marshal(bins)
	if err != nil {
		return fmt.Errorf("Не удалось преобразовать данные в JSON. %w", err)
	}

	err = s.db.Write(content)
	if err != nil {
		return fmt.Errorf("Не удалось создать файл. %w", err)
	}
	return nil
}

func (s *Storage) ReadBin(name string) (bins.BinList, error) {
	var bin bins.BinList
	isJson := file.CheckFile(name)
	if !isJson {
		return nil, errors.New("Файл не является JSON")
	}
	data, err := s.db.Read()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &bin)
	if err != nil {
		return nil, errors.New("Не удалось преобразовать JSON данные в слайс структур.")
	}
	return bin, nil
}
