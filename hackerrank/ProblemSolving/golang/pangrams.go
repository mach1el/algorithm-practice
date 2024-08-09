package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(strings.ToLower(s))

	cleaned := ""
	for _, ch := range s {
		if unicode.IsLetter(ch) {
			cleaned += string(ch)
		}
	}

	letterSet := make(map[rune]bool)
	for _, ch := range cleaned {
		letterSet[ch] = true
	}

	if len(letterSet) == 26 {
		fmt.Println("pangram")
	} else {
		fmt.Println("not pangram")
	}
}
