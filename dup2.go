package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)          // key: word,  value: number of how many times the word appears
	belongs := make(map[string][]string)    // key: word,  value: list of files in which the word appears
	test := make(map[string]map[string]int) // key: word, value: map  { key: file in which the word appears, value: number of how many times the word appears in this file}
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, belongs, test)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			}
			countLines(f, counts, belongs, test)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, belongs[line])
			for key, value := range test[line] {
				fmt.Printf("%s\t%d\n", key, value)
			}
			fmt.Println("  ")
		}
	}
}

func countLines(f *os.File, counts map[string]int, belongs map[string][]string, dict map[string]map[string]int) {
	fileName := f.Name()
	input := bufio.NewScanner(f)
	for input.Scan() {
		if !contains(belongs[input.Text()], fileName) {
			belongs[input.Text()] = append(belongs[input.Text()], fileName)
		}
		if _, containsKey := dict[input.Text()]; !containsKey {
			dict[input.Text()] = make(map[string]int)
		}
		dict[input.Text()][fileName]++
		counts[input.Text()]++
	}
}

func contains(s []string, target string) bool {
	for _, str := range s {
		if str == target {
			return true
		}
	}
	return false
}
