package main

import (
	"fmt"
	goutils "github.com/armando-couto/goutils"
)

func main() {
	files := goutils.ListFolderFiles("./arquivos")
	for _, f := range files {
		fmt.Println(f.Name())
		if f.Name() != ".DS_Store" {
			texto := goutils.ReadingFiles("./arquivos/", f.Name())
			linhas := goutils.ReturnsTheRows(texto)

			for _, l := range linhas {
				goutils.CreateFileDayInfoNotDate(l)
			}
		}
	}
}
