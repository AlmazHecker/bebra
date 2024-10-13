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

// var outputPath string = ""


var decompileCmd = &cobra.Command{
    Use:   "decompile [apk file]",
    Short: "Decompile the apk file",
	Args: cobra.ExactArgs(1),
    Run: decompileHanlder,
}

func decompileHanlder(cmd *cobra.Command, args []string) {
	configPath, _ := cmd.Flags().GetString("config")
    conf := config.GetConfig(configPath)

	if !helpers.FileExists(args[0]) {
		fmt.Printf("The given apk file(%s) not found!\n", args[0])
		os.Exit(1)
	}

	fmt.Println("Decompilation started...")

	outputPath, _ := cmd.Flags().GetString("output")

    if outputPath == "" {
		outputPath = conf.DecompileOutDir
    }

    osCmd := exec.Command(conf.Apktool, "d", args[0], "-o", outputPath)
    _, err := osCmd.CombinedOutput()

	if err != nil {
        log.Fatalf("Decompilation failed: %s\n", err)
    }

	fmt.Println("Decompilation done!")
}

func init() {
    decompileCmd.Flags().StringP("config", "c", "bebra.config.json", "Specify config path")
    decompileCmd.Flags().StringP("output", "o", "", "The output of decompiled apk")
}