package screenlock

import (
	"os/exec"
)

func Lock() error {
	_, err := exec.Command("rundll32.exe", "user32.dll", "LockWorkStation").Output()
	if err != nil {
		return err
	}
	return nil
}

