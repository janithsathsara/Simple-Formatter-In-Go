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

	for scanner.Scan() {
		runesScanned := scanner.Text()

		if runesScanned == ")" {
			runesScanned += "(==========)"
		}

		collected.WriteString(runesScanned) // Collects the runes to a single "collected" variable

		fmt.Printf("\n%s", runesScanned)
	}

	fmt.Println(collected.String())

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
