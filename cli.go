package main

import (
	"fmt"

	"bebra/commands"
	"bebra/config"
)

func main() {
	conf := config.GetConfig()
	fmt.Println(conf)

	commands.Init()
	commands.Execute()
}

