package commands

import (
	"fmt"
)

// HandleCommand processes the main command (like bebra)
func HandleCommand(command string) {
	switch command {
	case "bebra":
		fmt.Println("Bebra command executed!")
	default:
		fmt.Printf("Unknown command: %s\n", command)
	}
}
