package main

import (
	"bufio"
	"fmt"
	"os"
)

func File2lines(filePath string) []string {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	return lines
}

func main() {
	lines := File2lines("9dt/台語世界 九重天 鄭邦鎮 1997.dml3")
	for _, line := range lines {
		fmt.Println(line)
	}
}
