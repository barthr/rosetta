package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/barthr/rosetta/fetcher"
	"github.com/barthr/rosetta/settings"
)

const (
	searchCmd   = "search"
	languageCmd = "language"
	envVariable = "rosetta_lang"
)

var (
	s                  = new(settings.User)
	repo               = make(chan []string)
	languagePreference = "Go"
)

func setLanguagePreference(input string) {
	os.Setenv(envVariable, input)

}

func main() {
	executeCommand(os.Args[1:])
}

func executeCommand(arg []string) {
	switch arg[0] {
	case searchCmd:
		items := <-repo // Wait for program to complete fetching the tasks

		if len(arg) == 1 {
			printOptions(items)
			return
		}

		searchTerm := strings.ToLower(arg[1])

		matches := matcher(items, searchTerm)

		printOptions(matches)

		fmt.Print("Enter selection number: ") // Ask for input from the search result
		var input int
		fmt.Scanln(&input)

		pref := s.ReadSettings()

		fetcher.OpenWebsite(matches[input], pref.Language)

	case languageCmd:
		inputLang := arg[1]
		s.Language = inputLang

		s.WriteSettings()
	}
}

// Pretty Print the options provided by the matcher function
func printOptions(matches []string) {
	if len(matches) <= 0 {
		fmt.Println("Try again!")
		return
	}

	for i := 0; i < len(matches); i++ {
		fmt.Printf("%d) %s \n", i, matches[i])
	}
}

// getLanguagePreference returns the Prefered language
func getLanguagePreference() string {
	return os.Getenv(envVariable)
}

// matcher match the tasks against the search term
// if 1 of the tasks contains the search term add them to the result
// @Returns slice of Programming tasks which contains the search term
func matcher(items []string, term string) (result []string) {
	for _, task := range items {
		if strings.Contains(strings.ToLower(task), term) {
			result = append(result, task)
		}
	}
	return
}

// Asynchronous fetch the tasks from the website
func init() {
	os.Setenv(envVariable, languagePreference)

	go func() {
		categoryOutput, err := fetcher.GetProgrammingTasks()
		if err != nil {
			fmt.Println(err)
			return
		}
		repo <- categoryOutput
	}()
}
