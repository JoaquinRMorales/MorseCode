package Morse

import (
	"bufio"
	"os"
	"strings"
)

//User put his code
func Scan(str string) string {
	print(str, "  :  ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}

//Change code for letters
func Coder(str string) string {

	var acumulator string

	astr := strings.Split(str, "/")

	letter := make([]string, len(astr))

	for i, v := range astr {
		if v == ".-" {
			letter[i] = "a"
		}
		if v == "-..." {
			letter[i] = "b"
		}
		if v == "-.-." {
			letter[i] = "c"
		}
		if v == "-.." {
			letter[i] = "d"
		}
		if v == "." {
			letter[i] = "e"
		}
		if v == "..-." {
			letter[i] = "f"
		}
		if v == "--." {
			letter[i] = "g"
		}
		if v == "...." {
			letter[i] = "h"
		}
		if v == ".." {
			letter[i] = "i"
		}
		if v == ".---" {
			letter[i] = "j"
		}
		if v == "-.-" {
			letter[i] = "k"
		}
		if v == ".-.." {
			letter[i] = "l"
		}
		if v == "--" {
			letter[i] = "m"
		}
		if v == "-." {
			letter[i] = "n"
		}
		if v == "---" {
			letter[i] = "o"
		}
		if v == ".--." {
			letter[i] = "p"
		}
		if v == "--.-" {
			letter[i] = "q"
		}
		if v == ".-." {
			letter[i] = "r"
		}
		if v == "..." {
			letter[i] = "s"
		}
		if v == "-" {
			letter[i] = "t"
		}
		if v == "..-" {
			letter[i] = "u"
		}
		if v == "...-" {
			letter[i] = "v"
		}
		if v == ".--" {
			letter[i] = "w"
		}
		if v == "-..-" {
			letter[i] = "x"
		}
		if v == "-.--" {
			letter[i] = "y"
		}
		if v == "--.." {
			letter[i] = "z"
		}
		if v == "-----" {
			letter[i] = "0"
		}
		if v == ".----" {
			letter[i] = "1"
		}
		if v == "..---" {
			letter[i] = "2"
		}
		if v == "...--" {
			letter[i] = "3"
		}
		if v == "....-" {
			letter[i] = "4"
		}
		if v == "....." {
			letter[i] = "5"
		}
		if v == "-...." {
			letter[i] = "6"
		}
		if v == "--..." {
			letter[i] = "7"
		}
		if v == "---.." {
			letter[i] = "8"
		}
		if v == "----." {
			letter[i] = "9"
		}
		acumulator = acumulator + letter[i]
	}

	return acumulator + " "
}

//Is a word Separator
func Identificator(str string) string {

	sentences := strings.Split(str, "//")
	var acumulator string

	for _, v := range sentences {
		acumulator = acumulator + Coder(v)
	}

	return acumulator
}

func Morse() {
	fmt.Println("A .-   J .---   S ...   1 .----")
	fmt.Println("B -... K -.-    T -     2 ..---")
	fmt.Println("C -.-. L .-..   U ..-   3 ...--")
	fmt.Println("D -..  M --     V ...-  4 ....-")
	fmt.Println("E .    N -.     W .--   5 .....")
	fmt.Println("F ..-. O ---    X -..-  6 -....")
	fmt.Println("G --.  P .--.   Y -.--  7 --...")
	fmt.Println("H .... Q --.-   Z --..  8 ---..")
	fmt.Println("I ..   R .-.    0 ----- 9 ----.")
}
