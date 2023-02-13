package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	primeiraParte()
	segundaParte()
}

func primeiraParte() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	acumulador := 0
	for scanner.Scan() {
		valores := strings.Split(scanner.Text(), ",")
		primeiraSecao := strings.Split(valores[0], "-")
		segundaSecao := strings.Split(valores[1], "-")

		primeiroNumero, _ := strconv.Atoi(primeiraSecao[0])
		segundoNumero, _ := strconv.Atoi(primeiraSecao[1])

		terceiroNumero, _ := strconv.Atoi(segundaSecao[0])
		quartoNumero, _ := strconv.Atoi(segundaSecao[1])

		if primeiroNumero >= terceiroNumero && segundoNumero <= quartoNumero {
			acumulador += 1
		} else if terceiroNumero >= primeiroNumero && quartoNumero <= segundoNumero {
			acumulador += 1
		}
	}
	fmt.Println("valor total da primeira parte é ", acumulador)

}

func segundaParte() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	acumulador := 0
	for scanner.Scan() {
		valores := strings.Split(scanner.Text(), ",")
		primeiraSecao := strings.Split(valores[0], "-")
		segundaParte := strings.Split(valores[1], "-")

		primeiroNumero, _ := strconv.Atoi(primeiraSecao[0])
		segundoNumero, _ := strconv.Atoi(primeiraSecao[1])

		terceiroNumero, _ := strconv.Atoi(segundaParte[0])
		quartoNumero, _ := strconv.Atoi(segundaParte[1])

		if terceiroNumero >= primeiroNumero && terceiroNumero <= segundoNumero {
			acumulador += 1
		} else if quartoNumero >= primeiroNumero && quartoNumero <= segundoNumero {
			acumulador += 1
		} else if primeiroNumero >= terceiroNumero && primeiroNumero <= quartoNumero {
			acumulador += 1
		} else if segundoNumero >= terceiroNumero && segundoNumero <= quartoNumero {
			acumulador += 1
		}

	}
	fmt.Println("valor total da segunda parte é ", acumulador)

}
