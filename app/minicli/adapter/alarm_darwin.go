//go:build darwin
// +build darwin

package adapter

import (
	"os/exec"
)

func (d dataManager) Alarm() {
	go func() {
		_ = exec.Command("osascript", "-e", "tell application (path to frontmost application as text) to display dialog \"The tomato clock is completed.\" buttons {\"OK\"} with icon caution").Run()
	}()
	go func() {
		for i := 0; i < 3; i++ {
			_ = exec.Command("afplay", "/System/Library/Sounds/Tink.aiff", "-v", "15", "-q", "1").Run()
		}
	}()
}
