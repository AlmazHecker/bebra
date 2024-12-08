package commands

import (
	"bebra/helpers"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var signerCmd = &cobra.Command{
	Use:   "signer [dir with apk]",
	Short: "A tool to sign APK files",
	Args:  cobra.ExactArgs(1),
	Run:   signerHandler,
}

func signerHandler(cmd *cobra.Command, args []string) {
	if !helpers.DirExists(args[0]) {
		helpers.ErrorLog(fmt.Sprintf("The given file(%s) not found!\n", args[0]))
		os.Exit(1)
	}

	output, _ := cmd.Flags().GetString("output")
	keystore, _ := cmd.Flags().GetString("keystore")
	pass, _ := cmd.Flags().GetString("pass")

	if !helpers.FileExists(keystore) {
		helpers.ErrorLog("The keystore file not found!")
		os.Exit(1)
	}

	if err := helpers.Signer(args[0], output, keystore, BebraConfig.Signer, pass); err != nil {
		helpers.ErrorLog(fmt.Sprintf("Error signing APK: %v", err))
		os.Exit(1)
	}

	fmt.Println("APK signed successfully:", output)

}

func init() {
	signerCmd.Flags().StringP("keystore", "k", "", "Path to the keystore (JKS file)")
	signerCmd.Flags().StringP("pass", "p", "", "keystore password")
	signerCmd.Flags().StringP("output", "o", "./signed_apks", "Path for the output signed APK files")

	signerCmd.MarkFlagRequired("k")
	signerCmd.MarkFlagRequired("p")

}
