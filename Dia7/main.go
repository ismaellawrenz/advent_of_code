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

func main() {

	primeiraParte()
	//segundaParte()
}

type Pasta struct {
	Nome     string
	Arquivos []Arquivo
}

type Arquivo struct {
	Nome    string
	Tamanho int
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

	//primeiraLinha := true
	var pastas []Pasta

	pasta := Pasta{Nome: "/"}
	//pastas = append(pastas, pasta)

	var arquivos []Arquivo
	//pattern := regexp.MustCompile(`(\d+)`)
	nivel := ""
	for scanner.Scan() {

		if strings.Contains(scanner.Text(), "$ cd ") && scanner.Text() != "$ cd .." {
			dados := strings.Split(scanner.Text(), " ")
			nivel += "-"
			fmt.Println(nivel, dados[2], " (dir)")
		} else if scanner.Text() == "$ cd .." {
			nivel = nivel[0 : len(nivel)-1]
		} else if scanner.Text() != "$ ls" {
			fmt.Println(nivel + scanner.Text())
		}

		/*
			if !primeiraLinha {

				if strings.Contains(scanner.Text(), "cd ") && scanner.Text() != "$ cd .." {
					dados := strings.Split(scanner.Text(), " ")
					pasta.Arquivos = append(pasta.Arquivos, arquivos...)
					arquivos = nil
					pastas = append(pastas, pasta)
					pasta = Pasta{Nome: dados[2]}
				}

				match := pattern.FindAllString(scanner.Text(), -1)
				if len(match) > 0 {
					dados := strings.Split(scanner.Text(), " ")
					tamanho, _ := strconv.Atoi(dados[0])
					arquivo := Arquivo{Nome: dados[1], Tamanho: tamanho}
					arquivos = append(arquivos, arquivo)
				}

			} else {
				primeiraLinha = false
			}
		*/
	}

	pasta.Arquivos = append(pasta.Arquivos, arquivos...)
	pastas = append(pastas, pasta)
	somaTotal := 0
	somaDiretorio := 0
	for i := 0; i < len(pastas); i++ {
		for j := 0; j < len(pastas[i].Arquivos); j++ {
			somaDiretorio += pastas[i].Arquivos[j].Tamanho
		}
		if somaDiretorio <= 100000 {
			somaTotal = somaDiretorio + somaTotal
		}
		somaDiretorio = 0
	}

	fmt.Println("Resultado final da primeira parte->", somaTotal)

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
	pattern := regexp.MustCompile(`(\d+)`)
	var valores = make([][]string, 9)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), ("move")) {
			match := pattern.FindAllString(scanner.Text(), -1)
			quantidade, _ := strconv.Atoi(match[0])
			origem, _ := strconv.Atoi(match[1])
			destino, _ := strconv.Atoi(match[2])

			origem -= 1
			destino -= 1

			contador := quantidade - 1
			for i := 0; i < quantidade; i++ {

				valores[destino] = append([]string{valores[origem][contador-i]}, valores[destino]...)

			}
			valores[origem] = valores[origem][quantidade:]

			fmt.Print("")

		} else {
			if scanner.Text() != "" && scanner.Text()[1:2] != "1" {
				contador := 1
				for i := 0; i < len(valores); i++ {
					if scanner.Text()[contador:contador+1] != " " {
						valores[i] = append(valores[i], scanner.Text()[contador:contador+1])
					}
					contador += 4
				}
			}
		}
	}
	fmt.Println("Resultado final da segunda parte->")
	for i := 0; i < len(valores); i++ {
		fmt.Print(valores[i][0])
	}

}
