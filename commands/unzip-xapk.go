package commands

import (
	"bebra/helpers"
	"fmt"

	"github.com/spf13/cobra"
)


var unzipXAPKCmd = &cobra.Command{
    Use:   "unzip [xapk]",
    Args: cobra.ExactArgs(1),
    Short: "unzip the xapk file",
    Run:   unzipXAPKHandler, 
}

func unzipXAPKHandler(cmd *cobra.Command, args []string) {
	outputPath, _ := cmd.Flags().GetString("output")

	if !cmd.Flags().Changed("output") {
		fmt.Printf("The output is not provided. Output will be saved in %s dir.\n", outputPath)
	}

	helpers.Unzip(args[0], outputPath)

	fmt.Printf("Unzipped. The output is in %s dir", outputPath)
}


func init() {
    unzipXAPKCmd.Flags().StringP("output", "o", "./unzipped", "Specify output directory")
}
