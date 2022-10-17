package cmd

import (
	"encoding/json"
	"log"

	"github.com/mdelapenya/github-metrics/github"
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
	response, err := github.Labels(Owner, Repository)
	if err != nil {
		log.Fatalf("failed to proceed the request: %v", err)
	}

	var labels []types.Label
	json.Unmarshal(response, &labels)

	for _, label := range labels {
		response, err := github.IssuesCountByLabel(Owner, Repository, label.Name)
		if err != nil {
			log.Printf("failed to get issues count for '%s' label. Continuing. error: %v", label.Name, err)
			continue
		}

		var issuesSearch types.IssuesSearch
		json.Unmarshal(response, &issuesSearch)
		log.Printf("Number of Issues for %s: %d\n", label.Name, issuesSearch.TotalCount)
	}
}
