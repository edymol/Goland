package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	// Execute shell command to get list of processes
	cmd := exec.Command("ps", "-A", "-o", "pid,comm")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Parse output to extract PID and command name
	lines := strings.Split(string(output), "\n")
	for _, line := range lines[1:] { // Skip the header line
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			pid := fields[0]
			name := fields[1]
			fmt.Printf("PID: %s, Name: %s\n", pid, name)
		}
	}
}
