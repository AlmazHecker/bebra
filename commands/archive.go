package commands

import (
	"bebra/helpers"
	"fmt"

	"github.com/spf13/cobra"
)


var unzipCommand = &cobra.Command{
    Use:   "unzip [xapk]",
    Args: cobra.ExactArgs(1),
    Short: "unzip the xapk file",
    Run:   unzipHandler, 
}

var zipCommand = &cobra.Command{
    Use:   "zip [directory]",
    Args: cobra.ExactArgs(1),
    Short: "zip the files of xapk file",
    Run:   zipHandler, 
}


func unzipHandler(cmd *cobra.Command, args []string) {
	outputPath, _ := cmd.Flags().GetString("output")

	if !cmd.Flags().Changed("output") {
		fmt.Printf("The output is not provided. Output will be saved in %s dir.\n", outputPath)
	}

	helpers.Unzip(args[0], outputPath)

	fmt.Printf("Unzipped. The output is in %s dir", outputPath)
}

func zipHandler(cmd *cobra.Command, args []string) {
	outputPath, _ := cmd.Flags().GetString("output")

	if !cmd.Flags().Changed("output") {
		fmt.Printf("The output is not provided. Output will be saved as %s.xapk file.\n", outputPath)
	}

	if err := helpers.Zip(args[0], outputPath); err != nil {
		fmt.Printf("Error zipping: %v\n", err)
		return
	}
	fmt.Printf("Zipped. The output is in %s dir", outputPath)

}


func init() {
    unzipCommand.Flags().StringP("output", "o", "./unzipped", "Specify output directory")
    zipCommand.Flags().StringP("output", "o", "./android.xapk", "Specify output directory")
}
