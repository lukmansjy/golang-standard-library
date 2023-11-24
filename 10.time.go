package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now) // menggunakan time zone yg komputer gunakan
	fmt.Println(now.Local())

	utc := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	parse, _ := time.Parse(time.RFC3339, "2006-02-02T15:04:05Z")
	fmt.Println(parse)

	formatter := "2006-01-02 15:04:05"
	value := "2020-10-10 10:10:10"
	valueTime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(valueTime)
	}

	// Duration
	duration1 := time.Second * 100     // membuat durasi 100 detik
	duration2 := time.Millisecond * 10 // membuat durasi 10 ms
	duration3 := duration1 - duration2

	fmt.Println(duration3)
	fmt.Printf("duration: %d\n", duration3)

}
