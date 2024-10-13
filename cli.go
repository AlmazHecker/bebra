package main

import (
	"bebra/commands"
)

func main() {
	// by default
	// config.InitConfig("bebra.config.json")

	commands.Init()
	commands.Execute()
}

