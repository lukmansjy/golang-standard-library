package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	csvString := "Lukman, Wonogiri, L\n" +
		"Budi, Solo, L\n" +
		"Ani, Jogja, P\n"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}

	writer := csv.NewWriter(os.Stdout)
	_ = writer.Write([]string{"Lukman", "Wonogiri", "L"})
	_ = writer.Write([]string{"Budi", "Solo", "L"})
	_ = writer.Write([]string{"Ani", "Jogja", "P"})

	writer.Flush()
}
