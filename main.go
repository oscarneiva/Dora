package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	readFile()
}

func readFile() {
	dirReader, err := os.Open("./data/")
	if err != nil {
		fmt.Println("Error ", err)
	}
	files, err := dirReader.ReadDir(0)
	if err != nil {
		fmt.Println("Error ", err)
	}
	for i := range files {
		file, err := os.Open("./data/" + files[i].Name())
		if err != nil {
			fmt.Println("Error ", err)
		}
		reader := csv.NewReader(file)
		data, err := reader.ReadAll()
		if err != nil {
			fmt.Println("Error ", err)
		}
		for _, line := range data[1:] {
			if line != nil {
				letter := Letter{
					Subject: line[0],
					Date:    time.Now(),
					Period:  line[1],
					Student: line[2],
					Class:   line[3],
					Pass:    checkPass(line[4]),
				}
				recLetterGen(letter)
			} else {
				break
			}
		}
		file.Close()
	}
}

func checkPass(pass string) bool {
	if strings.Contains(pass, "y") || strings.Contains(pass, "Y") {
		return true
	}
	return false
}
