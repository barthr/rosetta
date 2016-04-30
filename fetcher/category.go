package fetcher

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/skratchdot/open-golang/open"
)

const (
	taskURL      = "http://rosettacode.org/wiki/Category:Programming_Tasks"
	selectionURL = "http://rosettacode.org/wiki/%s#%s"
)

var tasks = make([]string, 0)

// GetCategories prints the programming Tasks supported by Rosetta
// @Returns error when function cannot access the Rosetta website
func getTasks() error {
	page, err := goquery.NewDocument(taskURL)

	if err != nil {
		return err
	}

	// Regex to check if the page is valid
	// Todo Check for statuscode (way faster & easier)
	r := regexp.MustCompile("currently no text in this page")

	// Match the block where the page returns not found
	if r.MatchString(page.Find(".noarticletext").Text()) {
		// If there is a match the URL is not valid anymore
		return errors.New("The used Url for getting categories doesn't exist anymore")
	}

	elements := page.Find(".mw-category-group").Find("ul").Find("li")

	elements.Each(func(_ int, s *goquery.Selection) {
		task := s.Text()
		tasks = append(tasks, task)
	})
	return nil
}

// GetProgrammingTasks returns a []string of the available programming tasks from Rosetta
// If the array is already filled it will return the array
func GetProgrammingTasks() ([]string, error) {
	var err error
	if len(tasks) <= 0 {
		if err = getTasks(); err == nil {
			return tasks, nil
		}
		return nil, err
	}
	return tasks, nil
}

// OpenWebsite Formats the url to open with the browser
// Open's the browser with the url
func OpenWebsite(url, defaultLang string) {
	runnable := fmt.Sprintf(selectionURL, url, defaultLang)
	open.Run(runnable)
}
