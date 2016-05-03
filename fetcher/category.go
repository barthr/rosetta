package fetcher

import (
	"fmt"

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
	open.Run(WebsiteURL(url, defaultLang))
}

// WebsiteURL returns the formatted URL
func WebsiteURL(url, defaultLang string) string {
	return fmt.Sprintf(selectionURL, url, defaultLang)
}
