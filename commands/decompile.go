package commands

import (
	"bebra/helpers"
	"fmt"

	"github.com/spf13/cobra"
)



var decompileCmd = &cobra.Command{
    Use:   "decompile",
    Short: "Decompile the application",
    Run: decompileHanlder,
}

func decompileHanlder(cmd *cobra.Command, args []string) {
	configPath, _ := cmd.Flags().GetString("config")
    if !helpers.FileExists(configPath) {
        println("Config is not defined! Define bebra.config.json file")
    } 
	fmt.Println("Decompiling the application...")
	fmt.Println("Using config path:", configPath)
}

func init() {
    decompileCmd.Flags().StringP("config", "c", "bebra.config.json", "Specify config path")
}