package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	for scanner.Scan() {
		runesScanned := scanner.Text()

		fmt.Printf("\n%s", runesScanned)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
