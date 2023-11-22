package main

import (
	"flag"
	"fmt"
)

// jalankan dengan go run 04.flag.go -host 127.0.0.1 -username lukman -password rahasia -port 123
func main() {
	host := flag.String("host", "localhost", "Put your DB host")
	username := flag.String("username", "root", "Put your DB username")
	password := flag.String("password", "root", "Put your DB password")
	port := flag.Int("port", 0, "Put your DB port")

	flag.Parse()

	fmt.Println("host", *host)
	fmt.Println("username", *username)
	fmt.Println("password", *password)
	fmt.Println("port", *port)
}
