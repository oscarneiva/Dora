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
	for {
		line, err := reader.Read()
		if err == nil {
			fmt.Println("Error ", err)
		}
		if line[3] != "Class" {
			if line != nil {
				letter := Letter{
					Subject: line[0],
					Date:    time.Now(),
					Period:  line[1],
					Student: line[2],
					Class:   line[3],
					Pass:    true,
				}
				LetterGen(letter)
			} else {
				break
			}
		}
	}
	file.Close()
}

//func checkPass(letter char) bool{
//	if letter
//	return false
//}
