// NOT TESTED YET

package commands

import (
	"bebra/helpers"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var SignerCmd = &cobra.Command{
    Use:   "signer [apk]",
    Short: "A tool to sign APK files",
    Args: cobra.ExactArgs(1),
    Run: signerHandler,
}

func signerHandler(cmd *cobra.Command, args []string) {
	if !helpers.FileExists(args[0]) {
		fmt.Printf("The given file(%s) not found!\n", args[0])
        os.Exit(1)
	}

    input, _ := cmd.Flags().GetString("input")
    output, _ := cmd.Flags().GetString("output")
    keystore, _ := cmd.Flags().GetString("keystore")

    if err := signAPK(input, output, keystore); err != nil {
        log.Fatalf("Error signing APK: %v", err)
    }
    fmt.Println("APK signed successfully:", output)

}


func signAPK(inputAPK, outputAPK, keystore string) error {
    cmd := exec.Command("apksigner", "sign", "--ks", keystore, "--out", outputAPK, inputAPK)
    return cmd.Run()
}

func init() {

    rootCmd.Flags().StringP("kestore", "k", "", "Path to the keystore (JKS file)")
    rootCmd.Flags().StringP("input", "i", "", "Path to the input APK file")
    rootCmd.Flags().StringP("output", "o", "signed_app.apk", "Path for the output signed APK file")

    rootCmd.MarkFlagRequired("ks")
    rootCmd.MarkFlagRequired("input")

}