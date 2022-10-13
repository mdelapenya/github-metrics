package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ghm",
	Short: "Github Metrics is a very fast Github metrics calculator",
	Long:  `A Fast and Flexible Github metrics calculator with 🧡 by mdelapenya and friends in Go.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello ghm!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
