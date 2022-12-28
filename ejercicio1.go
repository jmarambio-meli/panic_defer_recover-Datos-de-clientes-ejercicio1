package main

import (
	"fmt"
	"os"
)

func main() {

	defer func() {
		fmt.Println("ejecución finalizada")
	}()

	_, err := os.Open("customer.txt")
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
	fmt.Println("fin")
}
