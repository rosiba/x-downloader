package main

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

func getURL(xURL string) (string, error) {
	cmd := exec.Command("yt-dlp", "-g", "-f", "b", xURL)

	var out bytes.Buffer
	cmd.Stdout = &out

	log.Println("yt-dlp output:", out.String())

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(out.String()), nil
}
