package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	CreateFile()
	Readfile()
}

func CreateFile() {
	fmt.Println("Creating a file in golang")
	file, err := os.Create("file.txt")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer file.Close()

	len, err := file.WriteString("Understanding file system in golang")
	if err != nil {
		log.Fatalf("failed writing to a file: %s", err)
	}

	fmt.Printf("\nFile Name: %s", file.Name())
	fmt.Printf("\nLength: %d bytes", len)
}

func Readfile() {
	fmt.Println("\nREading a file in golang")
	fileName := "file.txt"
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panicf("failed reading from file %s", err)
	}
	fmt.Printf("\nFile Name: %s", fileName)
	fmt.Printf("\nFile Size: %s", len(data))
	fmt.Printf("\nFile Data: %s", data)
}
