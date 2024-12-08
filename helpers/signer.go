package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func Signer(inputFolder string, outputFolder string, keystore string, signer string, pass string) error {
	files, err := filepath.Glob(filepath.Join(inputFolder, "*.apk"))
	if err != nil {
		ErrorLog(fmt.Sprintf("Error reading APK files: %s\n", err))
		os.Exit(1)
	}

	for _, file := range files {
		fmt.Printf("Signing %s...\n", file)

		outputFile := filepath.Join(outputFolder, filepath.Base(file))

		fmt.Println(file, outputFile, keystore)
		execCmd := signerCmd(file, outputFile, keystore, signer, pass)

		output, err := execCmd.CombinedOutput()

		if len(output) > 0 {
			fmt.Printf("Command output:\n%s\n", string(output))
		}

		if err != nil {
			ErrorLog(fmt.Sprintf("Failed to sign %s: %s\n", file, err))
			os.Exit(1)
		}

	}

	return err
}

func signerCmd(file string, outputFile string, keystore string, signer string, pass string) *exec.Cmd {
	outputDir := filepath.Dir(outputFile)
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		ErrorLog(fmt.Sprintf("Failed to create output directory: %s\n", err))
		return nil
	}
	return exec.Command(signer, "sign", "--ks", keystore, "--ks-pass", "pass:"+pass, "--out", outputFile, file)
}
