package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "snapman",
	Short:   "Snapman helps you manage your EBS snapshots in AWS",
	Long:    `A much easier way to manage snapshots in AWS with reporting built in.`,
	Version: "v0.1.0",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var (
	awsProfile string
	awsRegion  string
	tags       []string
)

// Execute is the main entry point for cobra
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&awsProfile, "profile", "default", "which profile to use from $HOME/.aws/config")
	rootCmd.PersistentFlags().StringVar(&awsRegion, "region", "us-east-1", "which aws region to use")
	addCommands()
}

func addCommands() {
	rootCmd.AddCommand(DeleteCmd)
	rootCmd.AddCommand(ListCmd)
	rootCmd.AddCommand(ReportCmd)
	rootCmd.AddCommand(VersionCmd)
}

// NotImplemented should be applied to commands that are still a WIP
func NotImplemented() {
	fmt.Println("Not yet implemented")
}

// TagsFlag will apply the tag flag to a given command, this is not a global flag
func TagsFlag(c *cobra.Command) {
	c.Flags().StringSliceVarP(&tags, "tags", "t", []string{}, "comma separated tags to match, ex. -t key=value,key=value")
}
