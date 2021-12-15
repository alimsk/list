package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

// source: https://gist.github.com/hyg/9c4afcd91fe24316cbf0
func openBrowser(url string) (err error) {
	switch runtime.GOOS {
	case "linux", "android":
		// use termux in android
		err = exec.Command("xdg-open", url).Run()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Run()
	case "darwin":
		err = exec.Command("open", url).Run()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	return
}
