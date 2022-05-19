package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	inputFile, err := os.Open("vcard.gob")
	if err != nil {
		fmt.Println(err)
		return
	}

	dec := gob.NewDecoder(inputFile)

	var card VCard
	dec.Decode(&card)
	fmt.Println(card)
	var add = card.Addresses
	fmt.Println(*add[0])
	fmt.Println(*card.Addresses[0])
}
