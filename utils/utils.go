package utils

import (
	"fmt"
	"os/user"
)

// GetHomeDir gives the home directory of the current user
func GetHomeDir() string {
	user, err := user.Current()

	if err != nil {
		panic(fmt.Errorf("An error occurred while trying to obtain current user's home directory. Error : Couldn't get details of the current user "))
	}

	return user.HomeDir
}
