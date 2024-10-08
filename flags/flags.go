package flags

import (
	"flag"
	"fmt"
)

var (
	TestFlag  = flag.Bool("test", false, "Run in test mode")
	StartFlag = flag.Bool("start", false, "Start the app")
	BuildFlag = flag.Bool("build", false, "Build the app")
)

// ParseFlags parses the flags
func ParseFlags() {
	flag.Parse()
}

// HandleFlags processes the flags after parsing
func HandleFlags() {
	if *TestFlag {
		fmt.Println("Test mode activated!")
	}

	if *StartFlag {
		fmt.Println("Starting the app!")
	}

	if *BuildFlag {
		fmt.Println("Building the app!")
	}
}
