package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

func random() string {
	file, err := os.Open("words.txt")
	if err != nil {
		fmt.Println("Error:", err)
		//	return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	words := []string{}
	for scanner.Scan() {
		words = append(words, strings.Split(scanner.Text(), " ")...)
	}
	randomNumber := rand.Intn(len(words))
	mot := words[randomNumber]
	fmt.Println(mot)
	return mot
}
func motcache(mot string) string {
	cache := []rune{}
	for range mot {
		cache = append(cache, '_')
	}
	println(string(cache))

	// Transforme le mot en liste de lettres
	lettres := []rune{}
	for _, lettre := range mot {
		lettres = append(lettres, rune(lettre))
	}

	long := len(mot)
	if long < 5 {
		for i := 0; i < 2; i++ {
			rev := rand.Intn(long - 1)
			cache[rev] = lettres[rev]
		}
	} else if long >= 5 {
		for i := 0; i < 4; i++ {
			rev := rand.Intn(long - 1)
			cache[rev] = lettres[rev]
		}
	}
	return string(cache)
}

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

func affichage(cache string, mot string) {
	// permet d'afficher le bon pendu selon le nombre de faute, le fichier texte fonctionnant 8 par 8
	faute := 0
	if faute == 0 {
		display(0, 8)
		faute = 1
	} else {
		display(faute*8, (faute*8)+8)
		faute += 1
	}
	if faute < 10 {
		fmt.Println("Mot caché:", cache)
		input(cache, mot)
	}
}

func input(cache string, mot string) {
	var i rune
	mot1 := []rune{}
	cache1 := []rune{}
	fmt.Print("Devinez une lettre \n")
	fmt.Scan(&i)
	for _, lettreeeee := range mot {
		mot1 = append(mot1, rune(lettreeeee))
	}
	for _, caca := range cache {
		cache1 = append(cache1, rune(caca))
	}
	for o := 0; o <= len(mot1)-1; o++ {
		if i == mot1[o] {
			cache1[o] = i
			fmt.Println(string(cache1))
		}
	}
	cache = string(cache1)
	fmt.Print(cache)
	affichage(cache, mot)
}

func main() {
	mot := random()
	cache := motcache(mot)
	fmt.Println("Mot caché:", cache)
	input(cache, mot)
}
