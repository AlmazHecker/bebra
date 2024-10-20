package commands

import (
	"bebra/config"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)


var configPath string
var BebraConfig config.Config

var rootCmd = &cobra.Command{
    Use:   "bebra",
    Short: "A command-line application",
    PersistentPreRun: func(cmd *cobra.Command, args []string) {
		BebraConfig = config.InitConfig(configPath)
	},

}

func Init() {
    rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "./bebra.config.json", "Bebra configuration file")
    
    rootCmd.AddCommand(compileCmd)
    rootCmd.AddCommand(decompileCmd)
    rootCmd.AddCommand(initializeCmd)
    rootCmd.AddCommand(keystoreCmd)
    rootCmd.AddCommand(signerCmd)
    rootCmd.AddCommand(zipCommand, unzipCommand)


}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
