package commands

import (
	"bebra/config"
	"bebra/helpers"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)


var compileCmd = &cobra.Command{
    Use:   "compile [decompiled-apk-dir]",
    Short: "Compile the application",
    Args: cobra.ExactArgs(1),
    Run:   compileHandler, 
}

func compileHandler(cmd *cobra.Command, args []string) {
	configPath, _ := cmd.Flags().GetString("config")
    conf := config.GetConfig(configPath)

	if !helpers.DirExists(args[0]) {
		fmt.Printf("The given dir(%s) not found!\n", args[0])
        os.Exit(1)
	}

	fmt.Println("Compilation started...")

	outputPath, _ := cmd.Flags().GetString("output")
	
    if outputPath == "" {
		outputPath = conf.CompiledOutDir
    }

    osCmd := exec.Command(conf.Apktool, "b", args[0], "-o", outputPath)
    _, err := osCmd.CombinedOutput()

	if err != nil {
        log.Fatalf("Compilation failed: %s\n", err)
    }

	fmt.Println("Compilation done!")
}

func init() {
    compileCmd.Flags().StringP("config", "c", "bebra.config.json", "Specify config path")
    compileCmd.Flags().StringP("output", "o", "./compiled.apk", "The output of decompiled apk")
}
