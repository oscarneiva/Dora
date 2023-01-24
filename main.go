package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func main() {
	readFile()
}

func readFile() {
	file, err := os.Open("./data/Recuperacao.csv")
	if err != nil {
		fmt.Println("Error ", err)
	}
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err == nil {
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
			LetterGen(letter)
		} else {
			break
		}
	}
	file.Close()
}

func checkPass(letter string) bool {
	if letter == "y" || letter == "Y" {
		return true
	}
	return false
}
