package commands

import (
	"bebra/config"
	"bebra/helpers"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var compileCmd = &cobra.Command{
    Use:   "compile",
    Short: "Compile the application",
    Run:   compileHandler, 
}

func compileHandler(cmd *cobra.Command, args []string) {


    configPath, _ := cmd.Flags().GetString("config")
	if !helpers.FileExists(configPath) {
        println("Config is not defined! Define bebra.config.json file")
		os.Exit(1)
    } 

    conf := config.GetConfig(configPath)
	fmt.Println(conf)


    fmt.Println("Compiling the application...")
    fmt.Println("Using config path:", configPath)
}

func init() {
    compileCmd.Flags().StringP("config", "c", "bebra.config.json", "Specify config path")
}
