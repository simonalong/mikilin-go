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
	Name string `match:"value={songjiang, chenzhen}"`
	Age  int
}

func main() {
	myTag()
}

func myTag() {

	entity := MyEntity{Name: "songjiang"}

	result, errMsg := mikilin.Check(entity)
	if !result {
		log.Print(errMsg)
	}
}
