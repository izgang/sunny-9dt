package main

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"strings"
)

type Sentence struct {
	Wm  string
	Hvg string
	Hv  string
	Dv  string
	Di  string
}

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

func ParseDml3(filepath string) []Sentence {
	lines := File2lines(filepath)

	var sentences []Sentence
	for _, line := range lines {
		if strings.HasPrefix(line, ".wm.") {
			s := Sentence{}
			s.Wm = line[4:]
			sentences = append(sentences, s)
			continue
		}
		if strings.HasPrefix(line, ".hvg.") {
			sentences[len(sentences)-1].Hvg = line[5:]
			continue
		}
		if strings.HasPrefix(line, ".hv.") {
			sentences[len(sentences)-1].Hv = line[4:]
			continue
		}
		if strings.HasPrefix(line, ".dv.") {
			sentences[len(sentences)-1].Dv = line[4:]
			continue
		}
		if strings.HasPrefix(line, ".di.") {
			sentences[len(sentences)-1].Di = line[4:]
			continue
		}
	}

	return sentences
}

func main() {
	sentences := ParseDml3("9dt/台語世界 九重天 鄭邦鎮 1997.dml3")

	t := template.Must(template.ParseFiles("index.tpml"))
	fo, _ := os.Create("output/index.html")
	defer fo.Close()
	t.Execute(fo, sentences)
}
