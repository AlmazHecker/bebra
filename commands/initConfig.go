package commands

import (
	"github.com/spf13/cobra"
)

var initializeCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a configuration",
	Run:   initializeHandler,
}

func initializeHandler(cmd *cobra.Command, args []string) {

	// here will be code that creates new bebra.config.json with comments

	// configPath, _ := cmd.Flags().GetString("config")
	// if !helpers.FileExists(configPath) {
	// 	println("Config is not defined! Define bebra.config.json file")
	// }

	// conf := config.GetConfig(configPath)
	// fmt.Println(conf)

	// fmt.Println("Decompiling the application...")
	// fmt.Println("Using config path:", configPath)
}

func init() {
	initializeCmd.Flags().StringP("config", "c", "bebra.config.json", "Specify config path")
}