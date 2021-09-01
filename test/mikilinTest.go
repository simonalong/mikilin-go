package main

import (
	mikilin "github.com/SimonAlong/Mikilin-go"
	"log"
)

type ValidateEntity struct {
	Name string `validate:"max=3"`
	Age  int    `validate:"max=3"`
}

type MyEntity struct {
	Name string
	Age  int `match:"value={12, 32}"`
}

func main() {
	myTag()
}

func myTag() {

	myentity := MyEntity{Age: 21}

	result, errMsg := mikilin.Check(myentity)
	if !result {
		log.Print(errMsg)
	}
}
