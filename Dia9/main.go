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

	var matris [20][20]string

	for i := 0; i < len(matris); i++ {
		for e := 0; e < len(matris[i]); e++ {
			matris[i][e] = "."
		}
	}
	matris[10][10] = "H"
	ultimaLinha := 10
	ultimaColuna := 10

	for i := 0; i < len(matris); i++ {
		for e := 0; e < len(matris[i]); e++ {
			fmt.Print(matris[i][e])
		}
		fmt.Println()
	}
	fmt.Println()
	for scanner.Scan() {
		d := strings.Split(scanner.Text(), " ")

		quantidade, _ := strconv.Atoi(d[1])

		if d[0] == "L" {
			novoValor := ultimaColuna - quantidade
			if novoValor < 0 {
				novoValor = 0
			}
			matris[ultimaLinha][novoValor] = "H"

			matris[ultimaLinha][ultimaColuna+1] = "#" //ultimo lugar do T
			for i := novoValor + 1; i < ultimaColuna+quantidade; i++ {
				matris[ultimaLinha][i] = "#" //movimenta o T
			}
			matris[ultimaLinha][novoValor+1] = "T"
			ultimaColuna = novoValor
		}
		if d[0] == "R" {
			novoValor := ultimaColuna + quantidade
			if novoValor < 0 {
				novoValor = 0
			}

			matris[ultimaLinha][novoValor] = "H"
			matris[ultimaLinha][ultimaColuna-1] = "#" //ultimo lugar do T
			for i := ultimaColuna; i < novoValor; i++ {
				matris[ultimaLinha][i] = "#"
			}
			matris[ultimaLinha][novoValor-1] = "T"

			ultimaColuna = novoValor
		}
		if d[0] == "U" {
			novoValor := ultimaLinha - quantidade
			if novoValor < 0 {
				novoValor = 0
			}

			matris[novoValor][ultimaColuna] = "H"

			matris[ultimaLinha+1][ultimaColuna] = "#"
			for i := novoValor; i < ultimaLinha; i++ {
				matris[i][ultimaColuna] = "#" //movimenta o T
			}
			matris[novoValor+1][ultimaColuna] = "T"

			ultimaLinha = novoValor

		}
		if d[0] == "D" {
			novoValor := ultimaLinha + quantidade
			if novoValor < 0 {
				novoValor = 0
			}

			matris[novoValor][ultimaColuna] = "H"
			matris[ultimaLinha-1][ultimaColuna] = "#"
			for i := ultimaLinha; i < novoValor; i++ {
				matris[i][ultimaColuna] = "#" //movimenta o T
			}
			matris[novoValor-1][ultimaColuna] = "T"

			ultimaLinha = novoValor
		}
		for i := 0; i < len(matris); i++ {
			for e := 0; e < len(matris[i]); e++ {
				fmt.Print(matris[i][e])
			}
			fmt.Println()
		}
		fmt.Println()
	}

}
