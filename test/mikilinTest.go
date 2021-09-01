package main

import (
	mikilin "github.com/SimonAlong/Mikilin-go"
	log "github.com/sirupsen/logrus"
)

type ValidateEntity struct {
	Name string `validate:"max=3"`
	Age  int    `validate:"max=3"`
}

type MyEntity struct {
	Name string
	Age  int `match:"value={asdf, 322}"`
}

func main() {
	myTag()
}

func myTag() {

	entity := MyEntity{Age: 32}

	result, errMsg := mikilin.Check(entity)
	if !result {
		log.Error(errMsg)
	}
}
