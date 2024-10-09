package commands

import (
	"fmt"
	"os/exec"
)
func Decompile() {
	cmd := exec.Command("echo", "Hello, World!")

    // Run the command and capture the output
    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println("Error executing command:", err)
        return
    }

    // Print the output
    fmt.Println(string(output))
}