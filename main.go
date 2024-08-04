package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/mattn/go-tty"
)

func main() {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()
	dic, _ := os.ReadFile("dictionary.txt")
	wordmap = strings.Split(string(dic), "\n")

	for {
		r, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		if r == 0 {
			continue
		}
		ph := Serch2(string(r))
		if ph != "0" {
			fmt.Print("\033[2K")
			fmt.Print("\033[G")
			fmt.Printf("\r" + color.HiBlackString(ph))
		} else {
			fmt.Print("\033[2K")
			fmt.Print("\033[G")
		}
	}
}

var wordmap []string = make([]string, 0)

func Serch(text string) (int, string) {
	for _, word := range wordmap {
		index := strings.Index(word, text)
		if index != -1 {
			return index, word
		}
	}
	return -1, "0"
}

func Serch2(text string) string {
	for _, word := range wordmap {
		index := strings.Index(word, text)
		if index == 0 {
			return word
		}
	}
	return "0"
}

func PlaceHolder() {

}
