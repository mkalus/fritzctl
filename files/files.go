package files

import (
	"fmt"
	"os/user"
	"path"
)

type currentUser func() (*user.User, error)

var obtainCurrentUser = user.Current

// InHomeDir returns the path to a file inside the home directory.
func InHomeDir(filename string) (string, error) {
	usr, err := obtainCurrentUser()
	if err != nil {
		return "", fmt.Errorf("cannot determine current user: %s", err)
	}
	return path.Join(usr.HomeDir, filename), nil
}
