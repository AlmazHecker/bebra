package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "bebra",
    Short: "A command-line application",
}

func Init() {
    rootCmd.AddCommand(compileCmd)
    rootCmd.AddCommand(decompileCmd)
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        return
    }
}
