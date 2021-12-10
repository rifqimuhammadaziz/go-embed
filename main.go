package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed version.txt
var version string

//go:embed signal.png
var image []byte

//go:embed files/*.txt
var path embed.FS

func main() {
	fmt.Println(version)

	err := ioutil.WriteFile("image_new.png", image, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dirEntryies, _ := path.ReadDir("files")
	for _, entry := range dirEntryies { // loop directory
		if !entry.IsDir() {
			fmt.Println(entry.Name()) // print all files name
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file)) // print content of file
		}
	}
}
