package main

import (
	"fmt"
	"log"
	rn "reverse-number/reverse" // "reverse-number/reverse" and can also be line 17
)

func main() {
	fmt.Print("Enter a number: ")

	var num int64

	if _, err := fmt.Scan(&num); err != nil {
		log.Fatalln("Scan for number failed:", err)
	}
	reversed := rn.ReverseNumber(num) // reverse.ReverseNumber(num) {rn is just alias for reverse package}
	fmt.Printf("Reversed number: %d\n", reversed)
}
