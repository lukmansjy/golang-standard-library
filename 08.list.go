package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()

	data.PushBack("Lukman")
	data.PushBack("Sanjaya")
	data.PushBack("Wonogiri")

	var head *list.Element = data.Front()
	fmt.Println(head.Value) // Lukman

	next := head.Next()
	fmt.Println(next.Value) // Sanjaya

	next = next.Next()
	fmt.Println(next.Value) // Wonogiri

	// dengan iterasi
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
