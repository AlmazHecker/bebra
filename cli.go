package main

import (
	"fmt"
	"os"

	"bebra/commands"
	"bebra/flags"
)

func printUsage() {
	fmt.Println("Usage: bebra <command> [OPTIONS]")
	fmt.Println("Command:")
	fmt.Println("  bebra      The main command without dashes")
	fmt.Println("Options:")
	fmt.Println("  -test      Run in test mode")
	fmt.Println("  -start     Start the bebra")
	fmt.Println("  -build     Build the bebra")
}

func main() {
	// Check if we have at least one argument (for the command)
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	// First argument is the command
	command := os.Args[1]

	// Parse the flags
	flags.ParseFlags()

	// Handle the command
	commands.HandleCommand(command)

	// Handle the flags
	flags.HandleFlags()
}
