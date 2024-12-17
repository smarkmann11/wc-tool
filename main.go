package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

type Counts struct {
	Lines      int
	Words      int
	Bytes      int
	Characters int
}

func countFunction(file *os.File) Counts {
	counts := Counts{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		counts.Lines++
		counts.Words += len(strings.Fields(line))
		counts.Bytes += len(line) + 1
		counts.Characters += utf8.RuneCountInString(line) + 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return counts
}

func main() {

	linesFlag := flag.Bool("l", false, "Show only lines of input")
	wordsFlag := flag.Bool("w", false, "Show only words of input")
	bytesFlag := flag.Bool("c", false, "Show only bytes of input")
	charactersFlag := flag.Bool("m", false, "Show only characters of input")

	var filepath string
	var args []string

	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-") {
			args = append(args, arg)
		} else {
			filepath = arg
		}
	}

	os.Args = append([]string{os.Args[0]}, args...)
	flag.Parse()

	if filepath == "" {
		fmt.Printf("Option 1: go run main.go [flags] <files>\nOption 2: go run main.go <file> [flags]\nDefault: go run main.go <file>")
		os.Exit(1)
	}

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	counts := countFunction(file)

	if *linesFlag {
		fmt.Printf("Lines: %d\n", counts.Lines)
	}
	if *wordsFlag {
		fmt.Printf("Words: %d\n", counts.Words)
	}
	if *bytesFlag {
		fmt.Printf("Bytes: %d\n", counts.Bytes)
	}
	if *charactersFlag {
		fmt.Printf("Characters: %d\n", counts.Characters)
	}
	if !*linesFlag && !*wordsFlag && !*bytesFlag && !*charactersFlag {
		fmt.Printf("Lines: %d\n", counts.Lines)
		fmt.Printf("Words: %d\n", counts.Words)
		fmt.Printf("Bytes: %d\n", counts.Bytes)
		fmt.Printf("Characters: %d\n", counts.Characters)
	}
}
