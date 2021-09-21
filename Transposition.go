package main

import (
	"fmt"
	"os"
)

var key = ""
var text string = ""
var order []int
var nbOfLetters = 0
var nbOfColumn = 0
var nbOfLine = 0
var moduloResult = 0
var results = make(map[int]string)

func main() {
	get_args()
	find_order()
	init_vars()
	fillMapWithColumnString()
	display()
}

func fillMapWithColumnString() {
	for i := 1; i <= nbOfColumn; i++ {
		if inArray(order[:moduloResult], i) {
			results[i] = string(text[0 : nbOfLine+1])
			text = text[nbOfLine+1 : len(text)]

		} else {
			results[i] = string(text[0:nbOfLine])
			text = text[nbOfLine:len(text)]
		}
	}
}

func display() {
	fmt.Println("Result :")
	for i := 0; i < nbOfLine+1; i++ {
		if i == nbOfLine {
			for j := 0; j < moduloResult; j++ {
				fmt.Print(string(results[order[j]][i]))
			}
		} else {
			for j := 0; j < nbOfColumn; j++ {

				fmt.Print(string(results[order[j]][i]))
			}
		}

	}
}

func inArray(array []int, in int) bool {
	for i := range array {
		if array[i] == in {
			return true
		}
	}
	return false
}

func find_order() {
	order = make([]int, len(key))
	for i := range key {
		order[i] = 1
		char := key[i]
		for j := range key {
			if char > key[j] || (char == key[j] && i > j) {
				order[i]++
			}
		}
	}
}

func get_args() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 2 {
		fmt.Println("Pls give the text & the key in arg to the run command")
		os.Exit(3)
	}
	text = argsWithoutProg[0]
	key = argsWithoutProg[1]
}

func init_vars() {
	nbOfLetters = len(text)
	nbOfColumn = len(order)
	nbOfLine = nbOfLetters / nbOfColumn
	moduloResult = nbOfLetters % nbOfColumn
}
