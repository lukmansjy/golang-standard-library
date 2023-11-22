package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("Error", err.Error())
	}

	resultInt, err := strconv.Atoi("1000")
	if err == nil {
		fmt.Println(resultInt)
	} else {
		fmt.Println("Error", err.Error())
	}

	resultBinary := strconv.FormatInt(999, 2)
	fmt.Println(resultBinary)

	resultString := strconv.Itoa(999)
	fmt.Println(resultString)
}
