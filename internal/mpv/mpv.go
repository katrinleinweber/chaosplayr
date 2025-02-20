package mpv

import (
	"log"
	"os/exec"
)

// launch MPV video player with url as parameter
func LaunchMPV(videoURL string) {
	mpvPath := "mpv"

	parameters := []string{
		videoURL,
	}

	cmd := exec.Command(mpvPath, parameters...)
	err := cmd.Start()
	if err != nil {
		log.Fatalf("Failed to start MPV: %v", err)
	}
}
