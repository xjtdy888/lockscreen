//+build !windows,!darwin

package screenlock

import (
	"errors"
)

func Lock() error {
	return errors.New("Unsupport")
}

