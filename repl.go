package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func startREPL(conf *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("wp defuxer > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		command := words[0]
		if cmd, ok := getCommands()[command]; ok {
			if err := cmd.callback(conf, words[1:]...); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Printf("%s is not a command", command)
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "exit the defuxer",
			callback:    exitDefuxer,
		},
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type config struct {
	baseURL  string
	client   *http.Client
	postList []string
	//apiKEY  string
}