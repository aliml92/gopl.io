package main

import (
	"bufio"
	"fmt"
	"unicode"

	// "strconv"
	"strings"
)

func main() {
	// An artificial input source.
	const input = "1234 5678 7854\n2565 5499\n5458 6313 1234567901234567890"
	wordCount := 0
	lineCount := 0
	inWord := true
	spaceStart := true
	scanner := bufio.NewScanner(strings.NewReader(input))
	// Create a custom split function by wrapping the existing ScanWords function.
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanLines(data, atEOF)
		if err == nil && token != nil {
			lineCount++
			for _, b := range token {
				if unicode.IsSpace(rune(b)) && spaceStart  {
					wordCount++
					inWord = false
				}
				if !unicode.IsSpace(rune(b)) {
					
				}
				if inWord {
					continue
				}  
			}
		}
		return
	}
	// Set the split function for the scanning operation.
	scanner.Split(split)
	// Validate the input
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}
}