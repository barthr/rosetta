package fetcher

import (
	"strings"
	"testing"
)

func TestTaskFetcher(t *testing.T) {
	err := getTasks()
	t.Log("Testing the tasks from the rosettacode.org website")
	if err != nil {
		t.Fail()
		t.Error(err)
	}
	if len(tasks) <= 0 {
		t.Fail()
		t.Error("tasks array is empty, the site dom has changed")
	}
}

func TestProgrammingTask(t *testing.T) {
	t.Log("Testing if the array is filled")
	_, err := GetProgrammingTasks()
	if err != nil {
		t.Fail()
	}
}

func TestWebsiteURL(t *testing.T) {
	t.Log("Testing if the formatting if the URL is correct")
	result := WebsiteURL("Simple_windowed_application", "Java")
	expectedResult := "http://rosettacode.org/wiki/Simple_windowed_application#Java"
	if strings.Compare(result, expectedResult) != 0 {
		t.FailNow()
	}
}
