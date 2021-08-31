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
	Name string `matcher:"size=2"`
	Age  int    `matcher:"value={12, 32};range=(12,30]"`
}

func main() {
	myTag()
}

func myTag() {

	myentity := MyEntity{}

	result, err := mikilin.Check(myentity)
	if !result {
		log.Print(err.Error())
	}
}
