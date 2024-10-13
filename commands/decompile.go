package commands

import (
	"bebra/helpers"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var decompileCmd = &cobra.Command{
    Use:   "decompile [apk file]",
    Short: "Decompile the apk file",
	Args: cobra.ExactArgs(1),
    Run: decompileHanlder,
}

func decompileHanlder(cmd *cobra.Command, args []string) {
	if !helpers.FileExists(args[0]) {
		fmt.Printf("The given apk file(%s) not found!\n", args[0])
		os.Exit(1)
	}

	fmt.Println("Decompilation started...")

	outputPath, _ := cmd.Flags().GetString("output")

    if !cmd.Flags().Changed("output") {
		outputPath = BebraConfig.DecompiledOutDir
    }

    osCmd := exec.Command(BebraConfig.Apktool, "d", args[0], "-o", outputPath)
    _, err := osCmd.CombinedOutput()

	if err != nil {
        log.Fatalf("Decompilation failed: %s\n", err)
    }

	fmt.Printf("The apk file was decompiled! The output is saved in %s\n", outputPath)
}

func init() {
    decompileCmd.Flags().StringP("output", "o", "", "The output of decompiled apk")
}