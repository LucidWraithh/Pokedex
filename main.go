package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	call        func() error
}

var commands map[string]cliCommand
var shouldExit bool

func main() {

	commands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays help message and all current commands available",
			call:        commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Ceases the Pokedex functioning",
			call:        commandExit,
		},
	}

	shouldExit = false

	scanner := bufio.NewScanner(os.Stdin)

	for !shouldExit {
		fmt.Printf("Pokedex > ")
		scanner.Scan()

		inputText := scanner.Text()

		com := commands[inputText]

		com.call()
	}

}

func commandHelp() error {

	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("Usage:\n\n\n\n")

	for command := range commands {
		name := commands[command].name
		desc := commands[command].description

		fmt.Printf("%s: %s\n\n", name, desc)
	}

	return nil
}

func commandExit() error {
	fmt.Println("Exit command used")

	os.Exit(0)

	return nil
}

func commandMap() error {
	return nil
}

func commandMapb() error {
	return nil
}
