package main

import (
	"time"
)

func main() {
	letter := Letter{
		Subject: "subject",
		Date: time.Now(),
		Period: 0,
		Student: "John Doe",
		Class: "5A",
		Pass: true,
	}

	LetterGen(letter)
}