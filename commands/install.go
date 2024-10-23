package commands

import (
	"bebra/helpers"

	"github.com/spf13/cobra"
)

var installAPK = &cobra.Command{
    Use:   "install [*.apk OR directory with apks]",
    Args: cobra.ExactArgs(1),
    Short: "Install single or splitted apks",
    Run:   installHandler, 
}

func installHandler (cmd *cobra.Command, args []string) {
    if helpers.FileExists(args[0]) {
        // install single apk
    }

    if helpers.DirExists(args[0]) {
        //install splitted apks
    }


    helpers.ErrorLog("File or directory not found!")
}

func init() {
}
