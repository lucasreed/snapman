package cmd

import (
	"github.com/spf13/cobra"
)

// DeleteCmd will delete snapshots based on tags
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete EBS snapshots based on tags",
	Long:  `This command will delete EBS snapshots filtered by tags.`,
	Run: func(cmd *cobra.Command, args []string) {
		NotImplemented()
	},
}
