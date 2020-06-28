package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		arg := os.Args[1]
		for _, val := range arg {
			if isValid(val) != true {
				log.Fatal("\tError! Not ASCII.")
			}
		}
		outputSlice := rangeLines(arg)
		var sliceByte []byte
		for _, slice := range outputSlice {
			for _, bytes := range slice {
				sliceByte = append(sliceByte, byte(bytes))
			}
		}
		fmt.Print(string(sliceByte))
	} else if len(args) == 0 {
		fmt.Println("Empty input!")
	} else if len(args) > 1 {
		fmt.Println("More than 1 argument.")
	}
}

func isValid(val rune) bool {
	if val >= ' ' && val <= '~' {
		return true
	}
	return false
}

func rangeLines(arg string) []string {
	content, err := ioutil.ReadFile("standard.txt")
	if err != nil {
		log.Fatal("\t", err)
	}

	splited := strings.Split(arg, "\\n")
	fileLine := 0
	var slice []string
	for _, word := range splited {
		for line := 1; line <= 8; line++ {
			for _, val := range word {
				fileLine = (int(val)-32)*9 + 1 + line
				start, end := getBorder(content, fileLine)
				slice = append(slice, string(content[start:end]))
			}
			slice = append(slice, string("\n"))
		}
	}
	return slice
}

func getBorder(content []byte, fileLine int) (int, int) {
	borderFileLine := 0
	borderEnd := 0
	b := -1
	for idx := range content {
		if content[idx] == '\n' {
			borderFileLine++
		}
		if borderFileLine == fileLine-1 {
			borderEnd = idx + 1
			b++
		}
	}
	return (borderEnd - b), borderEnd
}
