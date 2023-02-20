package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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

	somaTotal := 0
	pattern := regexp.MustCompile(`(\d+)`)

	var arquivos []Arquivo
	var caminho []string
	var pastas []Arquivo
	for scanner.Scan() {

		if strings.Contains(scanner.Text(), "$ cd") {
			if strings.Contains(scanner.Text(), "$ cd ..") {
				caminho = caminho[0 : len(caminho)-1]
			} else {
				dados := strings.Split(scanner.Text(), " ")
				caminho = append(caminho, dados[2])
			}
		}

		match := pattern.FindAllString(scanner.Text(), -1)
		if len(match) > 0 {
			dados := strings.Split(scanner.Text(), " ")
			tamanho, _ := strconv.Atoi(dados[0])
			arquivo := Arquivo{Nome: strings.Join(caminho, `|`) + `|` + dados[1], Tamanho: tamanho}
			arquivos = append(arquivos, arquivo)
		}

	}

	for i := 0; i < len(arquivos); i++ {
		fmt.Println(arquivos[i].Nome, arquivos[i].Tamanho)
		d := strings.Split(arquivos[i].Nome, "|")
		var caminhoInteiro = ""
		for j := 0; j < len(d)-1; j++ {
			caminhoInteiro += d[j] + "|"
			achou := false
			for q := 0; q < len(pastas); q++ {
				if pastas[q].Nome == caminhoInteiro {
					pastas[q].Tamanho += arquivos[i].Tamanho
					achou = true
				}
			}
			if !achou {
				pasta := Arquivo{Nome: caminhoInteiro, Tamanho: arquivos[i].Tamanho}
				pastas = append(pastas, pasta)
			}
		}

	}

	for i := 0; i < len(pastas); i++ {
		if pastas[i].Tamanho < 100000 {
			somaTotal += pastas[i].Tamanho
			fmt.Println("----------------", pastas[i].Nome, " tem ", pastas[i].Tamanho)
		}

	}

	fmt.Println("Resultado final da primeira parte->", somaTotal)

}
