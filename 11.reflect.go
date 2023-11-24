package main

import (
	"fmt"
	"reflect"
)

// angap aja kita tidak tahu isi struct nya apa aja, dan kita bisa melihatnya dengan package reflect
type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name    string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Email   string `required:"true" max:"10"`
}

func readField(value any) {
	valueType := reflect.TypeOf(value)
	fmt.Println("Type Name", valueType.Name())

	for i := 0; i < valueType.NumField(); i++ {
		var structField = valueType.Field(i)
		fmt.Println(structField.Name, "with type", structField.Type) // membaca nama field

		fmt.Println(structField.Tag.Get("required")) // get tag
		fmt.Println(structField.Tag.Get("max"))
	}
}

func IsValid(value any) (result bool) {
	result = true
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if result == false {
				return result
			}
		}
	}

	return result
}
func main() {
	sample := Sample{"Lukman"}
	sampleType := reflect.TypeOf(sample)
	structField := sampleType.Field(0) // membaca nama field

	fmt.Println(structField.Name)

	readField(sample)
	readField(Person{"Sanjaya", "", ""})

	person := Person{
		Name:    "Lukman",
		Address: "Wonogiri",
		Email:   "lukman@mail.com",
	}

	fmt.Println(IsValid(person))
}
