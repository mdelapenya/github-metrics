package cmd

import (
	"encoding/json"
	"log"

	"github.com/mdelapenya/github-metrics/http"
	"github.com/mdelapenya/github-metrics/types"
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
		getLabels()
	},
}

func getLabels() {
	req := http.Request{
		URL: "https://api.github.com/repos/" + Owner + "/" + Repository + "/labels",
	}
	response, err := http.Get(req)
	if err != nil {
		log.Fatalf("failed to proceed the request: %v", err)
	}

	var labels []types.Label
	json.Unmarshal(response, &labels)
	log.Println(labels)
}
