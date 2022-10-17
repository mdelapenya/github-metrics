package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/mdelapenya/github-metrics/formatters"
	"github.com/spf13/cobra"
)

var Format string
var OutputFile string
var Owner string
var Repository string

func init() {
	rootCmd.PersistentFlags().StringVarP(&Format, "format", "F", "console", "Response format")
	rootCmd.PersistentFlags().StringVarP(&OutputFile, "output", "o", "output.txt", "Output file where to write the results")
	rootCmd.PersistentFlags().StringVarP(&Owner, "owner", "O", "", "Github owner (organization or user)")
	rootCmd.PersistentFlags().StringVarP(&Repository, "repository", "R", "", "Github repository")
}

var rootCmd = &cobra.Command{
	Use:   "ghm",
	Short: "Github Metrics is a very fast Github metrics calculator",
	Long:  `A Fast and Flexible Github metrics calculator with 🧡 by mdelapenya and friends in Go.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if !formatters.IsValid(Format) {
			return fmt.Errorf("%s is not a valid format. Accepted: console", Format)
		}

		if !strings.EqualFold("console", Format) && OutputFile == "" {
			return fmt.Errorf("the output file cannot be empty when the formatter is not console")
		}

		if Owner == "" || Repository == "" {
			return fmt.Errorf("neither Owner or Repository cannot be empty")
		}
		return nil
	},
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
