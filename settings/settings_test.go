package settings

import (
	"os"
	"testing"
)

func TestSettings(t *testing.T) {
	t.Log("Testing creating and reading user Settings")
	user := &User{
		Language: "Go",
	}

	user.WriteSettings()
	userExpected := user.ReadSettings()

	if *user != userExpected {
		t.FailNow()
	}
}

func TestDeleteSettings(t *testing.T) {
	t.Log("Testing deleting a user")
	user := &User{
		Language: "Go",
	}

	if _, err := os.Stat("settings.gob"); err != nil {
		t.Log(err)
		t.FailNow()
	}
	user.DeleteSettings()

}
