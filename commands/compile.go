package commands

import (
	"bebra/helpers"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var compileCmd = &cobra.Command{
	Use:   "compile [decompiled-apk-dir]",
	Short: "Compile the application",
	Args:  cobra.ExactArgs(1),
	Run:   compileHandler,
}

func compileHandler(cmd *cobra.Command, args []string) {
	if !helpers.DirExists(args[0]) {
		helpers.ErrorLog(fmt.Sprintf("The given dir(%s) not found!\n", args[0]))
		os.Exit(1)
	}

	fmt.Println("Compilation started...")

	outputPath, _ := cmd.Flags().GetString("output")

	if !cmd.Flags().Changed("output") {
		outputPath = BebraConfig.CompiledOutDir
	}

	osCmd := exec.Command(BebraConfig.Apktool, "b", args[0], "-o", outputPath)
	output, err := osCmd.CombinedOutput()

	if len(output) > 0 {
		fmt.Printf("Command output:\n%s\n", string(output))
	}

	if err != nil {
		helpers.ErrorLog(fmt.Sprintf("Compilation failed: %s\n", err))
		os.Exit(1)
	}

	fmt.Printf("The APK file is compiled! The output is saved in %s\n", outputPath)
}

func init() {
	compileCmd.Flags().StringP("output", "o", "", "The output of decompiled apk")
}
