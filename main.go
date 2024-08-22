package main

import (
	"fmt"
	"log"
)

func main() {
	input := ""
	fmt.Print("Введите целое число: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Вы ввели число: %s\n", input)
}
