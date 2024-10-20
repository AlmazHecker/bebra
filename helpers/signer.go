package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func Signer(inputFolder string, outputFolder string, keystore string, buildToolsDir string) error {
    files, err := filepath.Glob(filepath.Join(inputFolder, "*.apk"))
    if err != nil {
        fmt.Printf("Error reading APK files: %s\n", err)
        os.Exit(1)
    }

    for _, file := range files {
        fmt.Printf("Signing %s...\n", file)

        outputFile := filepath.Join(outputFolder, filepath.Base(file))

        fmt.Println(file, outputFile, keystore)
		execCmd := signerCmd(file, outputFile, keystore, buildToolsDir)

        execCmd.Stdin = os.Stdin
        execCmd.Stdout = os.Stdout
        execCmd.Stderr = os.Stderr

		if err := execCmd.Run(); err != nil {
            fmt.Printf("Failed to sign %s: %s\n", file, err)
            os.Exit(1)
        }
    }

    return err;
}

func signerCmd(file string, outputFile string, keystore string, buildToolsDir string) *exec.Cmd {
    var apkSignerFile string

    if runtime.GOOS == "windows" {
        apkSignerFile = filepath.Join(buildToolsDir, "apksigner.bat")
    } else {
        apkSignerFile = filepath.Join(buildToolsDir, "apksigner.sh")
    }
    if !FileExists(apkSignerFile) {
        fmt.Printf("File doesn't exist", buildToolsDir)
        os.Exit(1)
    }

    outputDir := filepath.Dir(outputFile)
    if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
        fmt.Printf("Failed to create output directory: %s\n", err)
       return nil
    }

	return exec.Command(apkSignerFile, "sign", "--ks", keystore, "--ks-pass", "pass:random", "--out", outputFile, file)
}