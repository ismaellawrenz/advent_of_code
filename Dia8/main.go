package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Pasta struct {
	Nome     string
	Arquivos []Arquivo
}

type Arquivo struct {
	Nome    string
	Tamanho int
}
type TamanhoPasta struct {
	Nome    string
	Tamanho int
}

func main() {

	primeiraParte()

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
	var matriz [][]string
	for scanner.Scan() {
		d := strings.Split(scanner.Text(), "")
		matriz = append(matriz, d)
	}

	visivel := len(matriz) + len(matriz[0])

	for linha := 1; linha < len(matriz)-1; linha++ {

		for coluna := 1; coluna < len(matriz[linha])-1; coluna++ {
			maior := false
			//verifica a esquerda
			for d := 0; d < coluna; d++ {
				if matriz[linha][d] > matriz[linha][coluna] {
					maior = true
					break
				}
			}
			if !maior {
				visivel += 1
			}

			//se a verificacao anterior achou uma arvore maior procura em direção a direita
			if maior {
				//verifica a direita
				for d := coluna + 1; d < len(matriz[linha]); d++ {
					if matriz[linha][d] > matriz[linha][coluna] {
						maior = true
						break
					}
				}
				if !maior {
					visivel += 1

				}
			}

			if maior {
				//verifica a coluna superior
				for d := 0; d < linha; d++ {
					if matriz[d][coluna] > matriz[linha][coluna] {
						maior = true
						break
					}
				}
				if !maior {
					visivel += 1

				}
			}

			if maior {
				//verifica a coluna inferior
				for d := linha + 1; d < len(matriz); d++ {
					if matriz[d][coluna] > matriz[linha][coluna] {
						maior = true
						break
					}
				}
				if !maior {
					visivel += 1

				}
			}
		}
	}

	fmt.Println("Resultado final da primeira parte->", visivel)

}
