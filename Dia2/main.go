package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// A- pedra, B-papel, C-tesoura
// X-pedra, Y-papel, Z-tesoura
const (
	vitoria = 6
	empate  = 3
	derrota = 0
	pedra   = 1
	papel   = 2
	tesoura = 3
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
	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var pontuacao int = 0
	for scanner.Scan() {

		valores := strings.Split(scanner.Text(), " ")

		if valores[0] == "C" {
			if valores[1] == "Y" {
				pontuacao += derrota + papel
			} else if valores[1] == "X" {
				pontuacao += vitoria + pedra
			} else if valores[1] == "Z" {
				pontuacao += empate + tesoura
			}
		} else if valores[0] == "B" {
			if valores[1] == "Y" {
				pontuacao += empate + papel
			} else if valores[1] == "X" {
				pontuacao += derrota + pedra
			} else if valores[1] == "Z" {
				pontuacao += vitoria + tesoura
			}
		} else if valores[0] == "A" {
			if valores[1] == "Y" {
				pontuacao += vitoria + papel
			} else if valores[1] == "X" {
				pontuacao += empate + pedra
			} else if valores[1] == "Z" {
				pontuacao += derrota + tesoura
			}
		}

	}
	file.Close()
	fmt.Println("Pontuação final primeira parte =", pontuacao)
}

// A- pedra, B-papel, C-tesoura
// x-perder, y-empatar, z-vencer
func segundaParte() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	var pontuacao int = 0
	for scanner.Scan() {

		valores := strings.Split(scanner.Text(), " ")

		if valores[0] == "C" {
			if valores[1] == "Y" {
				pontuacao += empate + tesoura
			} else if valores[1] == "X" {
				pontuacao += derrota + papel
			} else if valores[1] == "Z" {
				pontuacao += vitoria + pedra
			}
		} else if valores[0] == "B" {
			if valores[1] == "Y" {
				pontuacao += empate + papel
			} else if valores[1] == "X" {
				pontuacao += derrota + pedra
			} else if valores[1] == "Z" {
				pontuacao += vitoria + tesoura
			}
		} else if valores[0] == "A" {
			if valores[1] == "Y" {
				pontuacao += empate + pedra
			} else if valores[1] == "X" {
				pontuacao += derrota + tesoura
			} else if valores[1] == "Z" {
				pontuacao += vitoria + papel
			}
		}

	}
	file.Close()
	fmt.Println("Pontuação final segunda parte =", pontuacao)
}
