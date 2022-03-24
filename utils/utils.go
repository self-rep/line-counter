package utils

import (
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func Sleep(tm int) {
	time.Sleep(time.Duration(tm) * time.Millisecond)
}

func Clear() {
	// detect OS
	operating_system := runtime.GOOS
	if operating_system == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			log.Fatal(err.Error())
		}
	} else if operating_system == "linux" {
		// Not sure if working, i work in windows environment
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			log.Fatal(err.Error())
		}
	}
}
