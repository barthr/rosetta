package settings

import (
	"encoding/gob"
	"fmt"
	"os"
	"unicode"
)

// User holds the settings for the user of Rosetta
type User struct {
	Language string
}

// WriteSettings write's the user object to Disk
// Writesettings create's a settings.gob when it is not there
// The method uses the encoding/gob from the std lib encode the User struct
// The user struct is then written and persisted to disk
// @ERROR error: exits the program when writing of the settings is not possible
func (u *User) WriteSettings() {
	// create a file
	dataFile, err := os.Create("settings.gob")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	u.toUpperCase()
	fmt.Println(u.Language)
	// serialize the data
	dataEncoder := gob.NewEncoder(dataFile)
	dataEncoder.Encode(u)

	dataFile.Close()
}

func (u *User) toUpperCase() {
	chars := []rune(u.Language)
	chars[0] = unicode.ToUpper(chars[0])
	u.Language = string(chars)
}

// ReadSettings reads the 'settings.gob' file from disk
// Decoding with encoding/gob from the std lib
// Decodes the bytes back to an User struct
// @RETURNS user: The User struct with user settings
func (u *User) ReadSettings() User {
	var data User

	// open data file
	dataFile, err := os.Open("settings.gob")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dataDecoder := gob.NewDecoder(dataFile)
	err = dataDecoder.Decode(&data)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dataFile.Close()

	return data
}

// DeleteSettings removes the settings.gob file
// @ERROR error: prints the error and returns
func (u *User) DeleteSettings() {
	file := "settings.gob"
	err := os.Remove(file)

	if err != nil {
		fmt.Println(err)
		return
	}
}
