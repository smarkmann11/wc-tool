package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func args() (string, string) {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Provide a file_path")
		os.Exit(1)
	}

	if len(args) == 1 {
		return "w", args[0]
	}

	return args[0], args[1]
}

func readFile(file *os.File, mode string) int {
	valueCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		switch mode {
		case "-l":
			valueCount++
		case "-w":
			words := strings.Fields(line)
			valueCount += len(words)
		case "-c":
			valueCount += len(line) + 1
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading File: %v\n", err)
		os.Exit(1)
	}

	return valueCount
}

func openFile(filepath string, mode string) int {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	return readFile(file, mode)
}

func main() {

	mode, filepath := args()
	count := openFile(filepath, mode)

	switch mode {
	case "-l":
		fmt.Printf("%d lines in file\n", count)
	case "-w":
		fmt.Printf("%d words in file\n", count)
	case "-c":
		fmt.Printf("%d bytes in file\n", count)
	default:
		fmt.Println("Unknown mode. Use -l, -w, or -c.")
	}
}
