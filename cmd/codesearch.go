package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mdelapenya/github-metrics/formatters"
	"github.com/mdelapenya/github-metrics/github"
	"github.com/mdelapenya/github-metrics/log"
	"github.com/mdelapenya/github-metrics/types"
	"github.com/spf13/cobra"
)

var Query string
var Title string

func init() {
	csCmd.PersistentFlags().StringVarP(&Query, "query", "Q", "", "Search query")
	csCmd.PersistentFlags().StringVarP(&Title, "title", "T", "Total files", "Title to be used in the output message")

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

	queryMap := map[string]string{}
	queryTokens := strings.Fields(Query)
	for i, token := range queryTokens {
		queryMap[fmt.Sprintf("query.%d", i)] = token
	}

	lr := &types.MetricResponse{
		Title:    Title,
		Count:    result.Count,
		Metadata: queryMap,
		Query:    Query,
	}

	formatter.Format(lr)
}
