package main

import (
	"fmt"
	"os"
)

// jalankan dengan go run 03.os.go Lukman Sanjaya atau go run 03.os.go "Lukman Sanjaya"
func main() {
	args := os.Args
	for _, arg := range args {
		fmt.Println(arg)
	}

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error", err.Error())
	}
}
