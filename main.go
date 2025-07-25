package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("normaltext.txt")

	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	// for scanner.Scan() {
	// 	processedString := scanner.Text()
	// 	fmt.Println(processedString)
	// }

	scanner.Split(bufio.ScanRunes) //split by character

	var collected strings.Builder

	previousChar := ""

	for scanner.Scan() {
		currentRune := scanner.Text()

		if currentRune == "{" {
			currentRune += "(==========)"
		}

		if previousChar == "{" && currentRune == "\n" {
			currentRune += "(==========)"
		}

		collected.WriteString(currentRune) // Collects the runes to a single "collected" variable

		fmt.Printf("\n%s", currentRune)

		previousChar = currentRune
	}

	fmt.Println(collected.String())

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
