package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var acumulador float64
	var maiorValor float64 = 0
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {

		if scanner.Text() != "" {
			valor, err := strconv.ParseFloat(scanner.Text(), 64)
			if err != nil {
				fmt.Println(err)
			} else {
				acumulador += valor
			}
		}
		if scanner.Text() == "" {
			if maiorValor < acumulador {
				maiorValor = acumulador
			}
			acumulador = 0
		}
	}

	fmt.Println("maior valor ", maiorValor)

}
