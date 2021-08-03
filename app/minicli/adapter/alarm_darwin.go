//go:build darwin
// +build darwin

package adapter

import "os/exec"

func (d dataManager) Alarm() {
	for i := 0; i < 3; i++ {
		_ = exec.Command("afplay", "/System/Library/Sounds/Tink.aiff", "-v", "15", "-q", "1").Run()
	}
}
