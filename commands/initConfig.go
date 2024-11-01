package commands

import (
	"bebra/config"
	"bebra/helpers"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var initializeCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a configuration",
	Run:   initializeHandler,
}

func initializeHandler(cmd *cobra.Command, args []string) {
	newConfig := config.Config{
		Apktool: "apktool-path",
		Adb: "adb-path",
		Signer: "apk-signer",
		DecompiledOutDir: "decompiled code out dir(not required)",
	}
	if helpers.FileExists("bebra.config.json") {
		fmt.Println("Bebra config already exists!")
		os.Exit(1)
	}

	file := helpers.CreateFile("bebra.config.json")
	defer file.Close()

	helpers.JSONEncoder(file, newConfig)

	println("Config has been created!")
}


  