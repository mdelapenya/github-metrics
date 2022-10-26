package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/mdelapenya/github-metrics/formatters"
	"github.com/mdelapenya/github-metrics/github"
	"github.com/mdelapenya/github-metrics/log"
	"github.com/mdelapenya/github-metrics/types"
	"github.com/spf13/cobra"
)

var Query string

func init() {
	csCmd.PersistentFlags().StringVarP(&Query, "query", "Q", "", "Search query")

	rootCmd.AddCommand(csCmd)
}

var csCmd = &cobra.Command{
	Use:   "cs",
	Short: "Return the number of occurrences for a search string",
	Long:  "Return the number of occurrences for a search string, scrapping the HTML outputted by cs.github.com, locating the HTML element with the total count",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if Query == "" {
			return fmt.Errorf("search query cannot be empty")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		getCount()
	},
}

func getCount() {
	response, err := github.CsSearch(Query)
	if err != nil {
		log.Fatal("failed to proceed the request: %v", err)
	}

	var result types.CodeCountResponse
	json.Unmarshal(response, &result)

	formatter := formatters.Get(Format, OutputFile)

	lr := &types.MetricResponse{
		Message: "Total files",
		Count:   result.Count,
	}

	formatter.Format(lr)
}
