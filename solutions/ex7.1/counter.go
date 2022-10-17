package main

import (
	// "bufio"
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"text/scanner"
)


type LineCounter struct {
	lines int
}

func (c *LineCounter) Write(p []byte) (n int, err error) {
	for _, b := range p {
		if b == '\n' {
			c.lines++
		}
	}
	return len(p), nil
}





func main(){
	f, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	c := &LineCounter{}
	c.Write(b)
	fmt.Printf("Line count is: %d\n", c.lines)
	
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	// count := 0
	// for scanner.Scan() {
	// 	count++
	// }
	// if err := scanner.Err(); err != nil {
	// 	fmt.Fprintln(os.Stderr, "reading standard input:", err)
	// }
	// fmt.Printf("%d\n", count)
	// scanner.Split(bufio.ScanLines)
	// // Count the words.
	// count := 0
	// for scanner.Scan() {
	// 	count++
	// 	// fmt.Println(scanner.Text())
	// }
	// if err := scanner.Err(); err != nil {
	// 	fmt.Fprintln(os.Stderr, "reading input:", err)
	// }
	// fmt.Printf("%d\n", count)
}