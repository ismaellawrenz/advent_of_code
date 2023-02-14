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

			for i := 0; i < quantidade; i++ {
				if len(valores[origem]) > 0 {
					valores[destino] = append([]string{valores[origem][0]}, valores[destino]...)
					valores[origem] = valores[origem][1:]
				}
			}

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
	fmt.Println("Resultado final da primeira parte->")
	for i := 0; i < len(valores); i++ {
		fmt.Print(valores[i][0])
	}

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
