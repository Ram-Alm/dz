package main

import (
	"fmt"
	"struct-dz/bins"
	"struct-dz/file"
	"struct-dz/storage"
)

func main() {
	st1 := storage.NewStorage(file.NewJsonDb("bins.json"))

	list := bins.NewBinList()
	list = append(list, *bins.NewBin("1", false, "Ram-bin"))

	err := st1.SaveBin(list)
	if err != nil {
		fmt.Println(err)
		return
	}
	loadedBin, err := st1.ReadBin("bins.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(loadedBin)
}
