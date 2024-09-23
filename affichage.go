package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func display(a int, b int) {
	// lis le fichier texte et imprime les lignes en particulier, a étant l'index de la ligne de départ et b l'index de la dernière ligne à imprimer
	readFile, err := os.Open("hangman.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	readFile.Close()
	for i := a; i <= b; i++ {
		fmt.Println(lines[i])
	}
}

func affichage() {
	// permet d'afficher le bon pendu selon le nombre de faute, le fichier texte fonctionnant 8 par 8
	faute := 0
	for i := 0; i <= 8; i++ {
		if faute == 0 {
			display(0, 8)
			faute = 1
		} else {
			display(faute*8, (faute*8)+8)
			faute += 1
		}
	}
}

func main() {
	display(64, 72)
}
