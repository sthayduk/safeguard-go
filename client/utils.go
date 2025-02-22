package client

import (
	"os/exec"
	"runtime"
)

// openBrowser opens the specified URL in the default web browser.
func openBrowser(url string) {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	case "linux":
		cmd = "xdg-open"
	default:
		log.Printf("Platform %s is not supported", runtime.GOOS)
		return
	}

	args = append(args, url)
	if err := exec.Command(cmd, args...).Start(); err != nil {
		log.Printf("Error opening browser: %v", err)
	}
}
