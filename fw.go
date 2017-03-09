package main

import (
	"strings"
	"fmt"
	"os"
	"io"
	"bufio"
)

func toFw(s string) string {
	retval := ""
	for _, runeValue := range s {
		if runeValue >= 0x21 && runeValue <= 0x7E {
			retval += string(rune(runeValue + 65248))
		} else if runeValue == rune(' ') {
			retval += string(0x3000)
		} else {
			retval += string(runeValue)
		}
	}
	return retval
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	if len(os.Args) > 1 {
		fmt.Println(toFw(strings.Join(os.Args[1:], string(0x3000))))
	} else {
		for {
			text, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}
			fmt.Print(toFw(text))
		}
	}
}
