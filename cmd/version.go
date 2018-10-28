package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// VersionCmd will output snapman's version
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of snapman",
	Long:  `All software has versions. This is snapman's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("snapman version v0.1.0")
	},
}
