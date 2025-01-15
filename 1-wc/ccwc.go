package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getFileData(fileName string) []byte {
	file, err := os.ReadFile(fileName)
	check(err)
	return file
}

func getScanner(fileData []byte) *bufio.Scanner {
	return bufio.NewScanner(strings.NewReader(string(fileData)))
}

func getCount(fileData []byte, splitFunction bufio.SplitFunc) int {
	scanner := getScanner(fileData)
	scanner.Split(splitFunction)
	count := 0
	for scanner.Scan() {
		count++
	}

	check(scanner.Err())
	return count
}

func flagsProvided(flags [4]*bool) bool {
	for _, f := range flags {
		if *f {
			return true
		}
	}
	return false
}

func main() {
	printByteCountPointer := flag.Bool("c", false, "prints the byte count")
	printLineCountPointer := flag.Bool("l", false, "prints the newline count")
	printWordCountPointer := flag.Bool("w", false, "prints the word count")
	printCharacterCountPointer := flag.Bool("m", false, "prints the character count")

	flag.Parse()

	if len(flag.Args()) == 1 {
		fileName := flag.Arg(0)
		fileData := getFileData(fileName)

		newLineCount := getCount(fileData, bufio.ScanLines)
		wordCount := getCount(fileData, bufio.ScanWords)
		byteCount := getCount(fileData, bufio.ScanBytes)

		if !flagsProvided([4]*bool{printByteCountPointer, printLineCountPointer, printWordCountPointer, printCharacterCountPointer}) {
			fmt.Printf("%v %v %v %v\n", newLineCount, wordCount, byteCount, fileName)
		}

		if *printByteCountPointer {
			fmt.Printf("%v %v\n", byteCount, fileName)
		}

		if *printLineCountPointer {
			fmt.Printf("%v %v\n", newLineCount, fileName)
		}

		if *printWordCountPointer {
			fmt.Printf("%v %v\n", wordCount, fileName)
		}

		if *printCharacterCountPointer {
			characterCount := getCount(fileData, bufio.ScanRunes)
			fmt.Printf("%v %v\n", characterCount, fileName)
		}

	} else if len(flag.Args()) == 0 {
		panic("no file provided")
	} else if len(flag.Args()) > 1 {
		panic("more than 1 file provided")
	}
}
