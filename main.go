package main

import (
	"bufio"
	"fmt"
	"os"
)
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			res := scanner.Text()
			words := cleanInput(res)
			fmt.Printf("Your command was: %s\n", words[0])
		}
	}
}
