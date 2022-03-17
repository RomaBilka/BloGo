package main

import (
	"fmt"
	"reflect"
)

type person struct {
	id    int    `comment:"Identification number"`
	name  string `comment:"Full name person"`
	phone string `comment:"Phone number of person"`
	ip    string `comment:"IP address when person registration"`
}

func main() {
	p := person{
		id:    1,
		name:  "John Doe",
		phone: "+123456789",
		ip:    "85.12.35.158",
	}

	pValue := reflect.ValueOf(p)
	pType := reflect.TypeOf(p)
	for i := 0; i < pValue.NumField(); i++ {
		field := pType.Field(i)
		if comment, ok := field.Tag.Lookup("comment"); ok {
			fmt.Printf("%s: %v\n", comment, pValue.Field(i))
		}
	}
}
