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

	var matris [200][200]string

	for i := 0; i < len(matris); i++ {
		for e := 0; e < len(matris[i]); e++ {
			matris[i][e] = "."
		}
	}
	matris[40][40] = "H"
	ultimaLinha := 40
	ultimaColuna := 40

	for scanner.Scan() {
		d := strings.Split(scanner.Text(), " ")

		quantidade, _ := strconv.Atoi(d[1])

		if d[0] == "L" {
			novoValor := ultimaColuna - quantidade
			if novoValor < 0 {
				novoValor = 0
			}
			matris[ultimaLinha][ultimaColuna] = "#"
			matris[ultimaLinha][novoValor] = "H"
			ultimaColuna = novoValor
		}
		if d[0] == "R" {
			novoValor := ultimaColuna + quantidade
			if novoValor < 0 {
				novoValor = 0
			}
			matris[ultimaLinha][ultimaColuna] = "#"
			matris[ultimaLinha][novoValor] = "H"
			ultimaColuna = novoValor
		}
		if d[0] == "U" {
			novoValor := ultimaLinha - quantidade
			if novoValor < 0 {
				novoValor = 0
			}
			matris[ultimaLinha][ultimaColuna] = "#"
			matris[novoValor][ultimaColuna] = "H"
			ultimaLinha = novoValor

		}
		if d[0] == "D" {
			novoValor := ultimaLinha + quantidade
			if novoValor < 0 {
				novoValor = 0
			}
			matris[ultimaLinha][ultimaColuna] = "#"
			matris[novoValor][ultimaColuna] = "H"
			ultimaLinha = novoValor
		}

	}

	fmt.Print(matris)

}
