package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}
