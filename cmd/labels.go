package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(labelsCmd)
}

var labelsCmd = &cobra.Command{
	Use:   "labels",
	Short: "Distribution of the labels across issues and pull requests",
	Long:  "Return all the labels on a project, and how they are distributed across issues and pull requests",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello labels!")
	},
}
