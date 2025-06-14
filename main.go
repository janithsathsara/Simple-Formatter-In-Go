package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("normaltext.txt")
	if err != nil {
		log.Fatal(err)
	}
	content := bufio.NewScanner(file)

	for content.Scan() {
		processedString := content.Text()
		fmt.Println(processedString)
	}

	if err := content.Err(); err != nil {
		log.Fatal(err)
	}
}
