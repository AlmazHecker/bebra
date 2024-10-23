package commands

import (
	"bebra/helpers"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var installAPKCmd = &cobra.Command{
    Use:   "install [*.apk OR directory with apks]",
    Args: cobra.ExactArgs(1),
    Short: "Install single or splitted apks",
    Run:   installHandler, 
}

func installHandler (cmd *cobra.Command, args []string) {
	apks := collectAPKFiles(args[0])
	if len(apks) == 0 {
		fmt.Println("No APK files found.")
		return
	}

	err := installMultipleAPKs(apks)
	if err != nil {
		fmt.Printf("Error installing APKs: %v\n", err)
	}

	fmt.Println("APKs installed successfully.")

}

func collectAPKFiles(inputPath string) []string {
    var apks []string

    info, err := os.Stat(inputPath)
    if err != nil {
        helpers.ErrorLog(fmt.Sprintf("Cannot read given input %s", err))
    }

    if info.IsDir() {
		err := helpers.TraverseDir(inputPath, collectAPKFilesCallback(&apks))
		if err != nil {
			helpers.ErrorLog(fmt.Sprintf("Error traversing directory: %v\n", err))
		}
	} else {
		if filepath.Ext(info.Name()) == ".apk" {
			apks = append(apks, inputPath)
		} else {
		    helpers.ErrorLog(fmt.Sprintf("The given input is not valid!"))
            os.Exit(1)
		}
	}

    return apks
}

func collectAPKFilesCallback(apks *[]string) func(string, os.FileInfo) error {
	return func(path string, info os.FileInfo) error {
		if !info.IsDir() && filepath.Ext(info.Name()) == ".apk" {
			*apks = append(*apks, path)
		}
		return nil
	}
}

func installMultipleAPKs(apks []string) error {
	args := append([]string{"install-multiple"}, apks...)
	cmd := exec.Command("adb", args...)

	fmt.Printf("Running command: %s\n", cmd.String())

	_, err := cmd.CombinedOutput()
	if err != nil {
        helpers.ErrorLog(fmt.Sprintf("ADB install failed: %v\n", err))
		return err
	}

	return nil
}


func init() {
}

