package cmd

import (
	"fmt"

	"github.com/lucasreed/snapman/internal/sess"
	"github.com/lucasreed/snapman/internal/snapshot"
	"github.com/spf13/cobra"
)

// ListCmd will list out snapshots based on tags. Will show 20 max unless passed -a or --num NUMBER
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List EBS snapshots based on tags",
	Long: `This command will list out your snapshots. By default it will
list all snapshots that match your tags. The tag flag must be used. --tags or -t
with a comma separated list of tags to match (ex. --tags KEY=VALUE,KEY=VALUE).`,
	Run: func(cmd *cobra.Command, args []string) {
		listSnaps(tags)
	},
}

func init() {
	TagsFlag(ListCmd)
}

func listSnaps(tags []string) {
	client := sess.GetClient(&awsRegion, &awsProfile)
	snaps := snapshot.Get(client, tags)
	for _, v := range snaps {
		fmt.Println(*v.SnapshotId, *v.Description)
	}
	fmt.Printf("=======\nTotal Number of snapshots: %v\n", len(snaps))
}
