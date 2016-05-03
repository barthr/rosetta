package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/barthr/rosetta/fetcher"
	"github.com/barthr/rosetta/settings"
	"github.com/codegangsta/cli"
)

var (
	s          = new(settings.User)
	repo       = make(chan []string)
	searchLang string
)

func main() {
	app := cli.NewApp()

	app.Name = "Rosetta snippets"
	app.Usage = "Quickly find code snippets for your language"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:   "language",
			Usage:  "Set the language for Rosetta",
			Action: languageCommand,
		},
		{
			Name:   "reset",
			Usage:  "Removes and resets all your settings",
			Action: removeCommand,
		},
		{
			Name:   "settings",
			Usage:  "Show all your settings",
			Action: showCommand,
		},
		{
			Name:   "search",
			Usage:  "Search the rosettacode.org snippets repository",
			Action: searchCommand,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "l",
					Usage:       "Search for a specific language",
					Destination: &searchLang,
				},
				cli.BoolFlag{
					Name:  "r",
					Usage: "Providing this flag only returns the url, not opening it",
				},
			},
		},
	}

	app.Run(os.Args)
}

func showCommand(c *cli.Context) {
	settings := s.ReadSettings()
	fmt.Printf("UR SETTINGS:\n    Search language: %s", settings.Language)
}

func removeCommand(c *cli.Context) {
	s.DeleteSettings()
	fmt.Println("Deleted Settings!")
	os.Exit(0)
}

func languageCommand(c *cli.Context) {
	if c.NArg() > 0 {
		s.Language = c.Args()[0]
		s.WriteSettings()
		return
	}
	fmt.Println("Please provide a language!")
	os.Exit(13)
}

func searchCommand(c *cli.Context) {
	items := <-repo // Wait for program to complete fetching the tasks

	if c.NArg() > 0 { // Check if args are provided

		args := c.Args()[0] // Search term

		var searchInput string // Placeholder for the language to be returned with

		matches := match(items, args) // Match the search term against the items from Rosetta

		printOptions(matches) // Print the correspondending options

		if len(c.String("l")) > 0 { // If the there is a -l provided, than use that search term instead of the one from the settings
			searchInput = settings.ToUpperCaseFirst(searchLang)
		} else {
			pref := s.ReadSettings()
			searchInput = pref.Language
		}

		input := getSelectionFromUser() // Get the input selection from the user, which of the tasks he selects

		for !validIndex(input, matches) { // Check if user selects something valid
			fmt.Println("U cannot choose a number which is not in the list")
			input = getSelectionFromUser() // Retry
		}

		if c.Bool("r") { // if -r flag is provided than do not open the webpage but only return the url
			fmt.Println(fetcher.WebsiteURL(matches[input], searchInput))
			return
		}

		fetcher.OpenWebsite(matches[input], searchInput)
	} else {
		printOptions(items)
	}
	os.Exit(0)
}

func getSelectionFromUser() int {
	fmt.Print("Enter selection number: ") // Ask for input from the search result
	var input int
	fmt.Scanln(&input)
	return input
}

// validIndex takes the absolute value of the index
// The absolute value gets checked against the length of the options
// By using math.Abs from the std lib, no negatives allowed
func validIndex(index int, options []string) bool {
	if int(math.Abs(float64(index))) > len(options) {
		return false
	}
	return true
}

// Pretty Print the options provided by the match function
func printOptions(matches []string) {
	if len(matches) <= 0 {
		fmt.Println("Try again!")
		os.Exit(1)
	}

	for i := 0; i < len(matches); i++ {
		fmt.Printf("%d) %s \n", i, matches[i])
	}
}

// match match the tasks against the search term
// if 1 of the tasks contains the search term add them to the result
// @Returns slice of Programming tasks which contains the search term
func match(items []string, term string) (result []string) {
	for _, task := range items {
		if strings.Contains(strings.ToLower(task), strings.ToLower(term)) {
			result = append(result, task)
		}
	}
	return
}

// Asynchronous fetch the tasks from the website
func init() {
	go func() {
		categoryOutput, err := fetcher.GetProgrammingTasks()
		if err != nil {
			fmt.Println(err)
			return
		}
		repo <- categoryOutput
	}()
}
