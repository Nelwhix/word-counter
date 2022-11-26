package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"flag"
)

func main() {
	lines := flag.Bool("lines", false, "Count lines")
	bytes := flag.Bool("bytes", false, "Count Bytes")
	flag.Parse()
	fmt.Println(count(os.Stdin, *lines, *bytes))
}

func count(r io.Reader, countLines bool, countBytes bool) int {
	scanner := bufio.NewScanner(r)

	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	if countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	count := 0

	for scanner.Scan() {
		count++
	}

	return count
}