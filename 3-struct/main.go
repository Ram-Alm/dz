package main

import (
	"struct-dz/file"
	"struct-dz/storage"
)

func main() {
	storage.NewStorage(file.NewJsonDb("bins.json"))
}
