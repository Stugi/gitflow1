package main

import (
	"fmt"
	"log"
)

func main() {
	input := ""
	fmt.Print("Введите данные: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Вы ввели: %s\n", input)
}
