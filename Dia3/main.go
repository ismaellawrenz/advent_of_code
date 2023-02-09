package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

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

		tamanho := len(scanner.Text())
		primeiraParte := scanner.Text()[0 : tamanho/2]
		segundaParte := scanner.Text()[tamanho/2 : tamanho]

		for i := 0; i < len(primeiraParte); i++ {
			if strings.Count(segundaParte, string(primeiraParte[i])) > 0 {

				acumulador += buscarPrioridade(string(primeiraParte[i]))
				break
			}

		}

	}
	fmt.Println("valor total é ", acumulador)

}

func buscarPrioridade(letra string) int {
	alfabeto := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	valor := strings.Index(alfabeto, strings.ToUpper(letra)) + 1

	if strings.ToUpper(letra) == letra {
		valor = valor + 26
	}

	return valor
}
