package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	//primeiraParte()
	segundaParte()
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
	linha := 0
	primeiraLinha := ""
	segundaLinha := ""
	terceiraLinha := ""
	for scanner.Scan() {
		linha++
		if linha == 4 {
			for i := 0; i < len(primeiraLinha); i++ {
				if strings.Count(segundaLinha, string(primeiraLinha[i])) > 0 {
					if strings.Count(terceiraLinha, string(primeiraLinha[i])) > 0 {
						acumulador += buscarPrioridade(string(primeiraLinha[i]))
						break
					}
				}
			}
			linha = 1
		}

		if linha == 1 {
			primeiraLinha = scanner.Text()
		} else if linha == 2 {
			segundaLinha = scanner.Text()
		} else if linha == 3 {
			terceiraLinha = scanner.Text()
		}
	}

	if linha == 3 {
		//calcula após passar por todas as linhas do arquivo
		for i := 0; i < len(primeiraLinha); i++ {
			if strings.Count(segundaLinha, string(primeiraLinha[i])) > 0 {
				if strings.Count(terceiraLinha, string(primeiraLinha[i])) > 0 {
					acumulador += buscarPrioridade(string(primeiraLinha[i]))
					break
				}
			}
		}
	}

	fmt.Println("valor total da segunda parte é ", acumulador)
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
	fmt.Println("valor total da primeira parte é ", acumulador)

}

func buscarPrioridade(letra string) int {
	alfabeto := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	valor := strings.Index(alfabeto, strings.ToUpper(letra)) + 1

	if strings.ToUpper(letra) == letra {
		valor = valor + 26
	}

	return valor
}
