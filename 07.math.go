package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Ceil(1.40))  // Bulatkan ke atas
	fmt.Println(math.Floor(1.60)) // Bulatkan ke bawah
	fmt.Println(math.Round(1.60)) // Bulatkan ke terdekat
	fmt.Println(math.Max(10, 11))
	fmt.Println(math.Min(10, 11))
}
