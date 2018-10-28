package cmd

import (
	"github.com/spf13/cobra"
)

// ReportCmd will spit out a report of the snapshots based on tags
var ReportCmd = &cobra.Command{
	Use:   "report",
	Short: "Report will generate statistics about EBS snapshots and output them to stdout",
	Long: `This command will generate a report about the snapshots in your account + region.
As always, the snapshots can be filtered by tags.`,
	Run: func(cmd *cobra.Command, args []string) {
		NotImplemented()
	},
}
