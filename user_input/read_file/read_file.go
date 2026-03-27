package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("rawfile.txt")
	if err!=nil{
		fmt.Println("Error opening file", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineNum := 0

	for scanner.Scan(){
		lineNum++
		line := scanner.Text()
		fmt.Println("line number ",lineNum," :",line)
	}
	if err := scanner.Err(); err!=nil{
		fmt.Println("Error for scanning: ",err)
	}
}