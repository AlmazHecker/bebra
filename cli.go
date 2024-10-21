package main

import (
	"bebra/assets"
	"bebra/commands"
)


func main() {
	println(assets.ASCII_LOGO)

	commands.Init()
	commands.Execute()
}

