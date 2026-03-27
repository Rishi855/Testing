package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var allWords[] string
	// newScanner := bufio.NewScanner(os.Stdin)
	lineCount := 0

	fmt.Println("Enter line one by one:")
	for scanner.Scan(){
		line := strings.TrimSpace(scanner.Text())
		if len(line)==0{
			break
		}
		lineCount++
		words:= strings.Fields(line)
		allWords = append(allWords, words...)
	}
	letters := 0
	digit := 0
	for _, word := range allWords{
		for _, r := range word{
			if unicode.IsLetter(r){
				letters++
			}else if unicode.IsDigit(r){
				letters++
			}
		}
	}
	fmt.Printf("\nLines: %d",lineCount)
	fmt.Printf("\nWords: %d",len(allWords))
	fmt.Printf("\nLetters: %d",letters)
	fmt.Printf("\nDigits: %d",digit)
}
