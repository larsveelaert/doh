package main

import (
	"fmt"
	"os"
	"runtime"
	"os/exec"
)

/* A Simple function to verify error */
func CheckError(err error) {
	if err  != nil {
		fmt.Println("Error: " , err)
		//GUI
		if runtime.GOOS == "darwin" { // macos
			exec.Command("sh", "-c", "osascript -e 'tell app \"System Events\" to display dialog \"DoH failed to start/exited. Unencrypted DNS requests could leak to network\"'").Run()
		} else if runtime.GOOS == "windows" {
			//TODO add windows gui error
		}
		os.Exit(0)
	}
}
