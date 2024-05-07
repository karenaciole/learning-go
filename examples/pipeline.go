//Faça um pipeline em que uma thread gera strings aleatórias, 
//enquanto uma segunda filtra as strings que contém somente valores alpha, 
//e uma terceira escreva os valores filtrados na saída padrão

package main

import (
	"fmt"
	"math/rand"
	"time"
	"unicode"
)

//below random string functions are based on Jon Calhoun code
const charset = "abcdefghijklmnopqrstuvwxyz1234567890"

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length) 
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func RandString(length int) string {
	return StringWithCharset(length, charset)
}

func isLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func generateContent(out chan string) {
	for {
		out <- RandString(5)
	}
}

func filterContent(in chan string, out chan string) {
	for {
		word := <-in
		if isLetter(word) {
			out <- word
		}
	}
}

func main() {
	rawContent := make(chan string)
	filteredContent := make(chan string)

	go generateContent(rawContent)
	go filterContent(rawContent, filteredContent)

	for {
		fmt.Printf("alpha: <%s>\n", <-filteredContent)
	}
}
