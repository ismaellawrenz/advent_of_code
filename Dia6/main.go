package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	primeiraParte()
	segundaParte()
}

func primeiraParte() {
	arquivo, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	dados := string(arquivo)

	fluxo := ""
	for i := 0; i < len(dados); i++ {
		if len(fluxo) > 3 {
			fluxo = fluxo[1:4]
		}
		fluxo += string(dados[i])
		if len(fluxo) == 4 {
			repete := false
			for j := 0; j < len(fluxo); j++ {
				if strings.Count(fluxo, string(fluxo[j])) > 1 {
					repete = true
				}
			}
			if !repete {
				fmt.Println("Resultado final da primeira parte->", i+1)
				break
			}

		}
	}

}

func segundaParte() {
	arquivo, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	dados := string(arquivo)

	fluxo := ""
	for i := 0; i < len(dados); i++ {
		if len(fluxo) > 13 {
			fluxo = fluxo[1:14]
		}
		fluxo += string(dados[i])
		if len(fluxo) == 14 {
			repete := false
			for j := 0; j < len(fluxo); j++ {
				if strings.Count(fluxo, string(fluxo[j])) > 1 {
					repete = true
				}
			}
			if !repete {
				fmt.Println("Resultado final da segunda parte->", i+1)
				break
			}

		}
	}

}
