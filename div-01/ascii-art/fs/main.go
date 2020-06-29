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
	if len(args) == 0 {
		log.Fatal("\tError! Empty input.")
	} else if len(args) > 0 && len(args) <= 2 {
		arg := os.Args[1]
		style := "standard.txt"
		if len(args) == 2 {
			style = args[1]
			if style != "standard" && style != "shadow" && style != "thinkertoy" {
				log.Fatal("\tError! Invalid style.")
			}
			style += ".txt"
		}
		for _, val := range arg {
			if isValid(val) == false {
				log.Fatal("\tError! Not ASCII.")
			}
		}
		printLines(arg, style)
	} else {
		log.Fatal("\tError! More than 2 arguments.")
	}
}

func isValid(val rune) bool {
	if val >= ' ' && val <= '~' {
		return true
	}
	return false
}

func rangeLines(arg string, style string) []string {
	content, err := ioutil.ReadFile(style)
	if err != nil {
		content, err = ioutil.ReadFile("../" + style)
		if err != nil {
			log.Fatal("\tError! No such file in directory.")
		}
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

func printLines(arg string, style string) {
	outputSlice := rangeLines(arg, style)
	var sliceByte []byte
	for _, slice := range outputSlice {
		for _, bytes := range slice {
			sliceByte = append(sliceByte, byte(bytes))
		}
	}
	fmt.Print(string(sliceByte))
}
