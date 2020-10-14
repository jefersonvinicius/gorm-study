package helpers

import (
	"os"
	"os/exec"
	"runtime"
)

// ClearScreen : clear screen
func ClearScreen() {
	clear := make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear[runtime.GOOS]()
}

// GetMapKeys : get map keys
func GetMapKeys(data map[string]func()) []string {
	keys := make([]string, 0, len(data))
	for key := range data {
		keys = append(keys, key)
	}
	return keys
}
