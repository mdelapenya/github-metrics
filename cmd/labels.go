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

	failedLabels := []types.Label{}

	for _, label := range labels {
		err := getLabel(label.Name)
		if err != nil {
			log.Printf("failed to get issues count for '%s' label. Will be retried later. error: %v", label.Name, err)
			failedLabels = append(failedLabels, label)
			continue
		}
	}

	for _, failedLabel := range failedLabels {
		err := getLabel(failedLabel.Name)
		if err != nil {
			log.Printf("failed to get issues count for '%s' label. Continuing. error: %v", failedLabel.Name, err)
			continue
		}
	}
}

func getLabel(label string) error {
	response, err := github.IssuesCountByLabel(Owner, Repository, label)
	if err != nil {
		return err
	}

	var issuesSearch types.IssuesSearch
	json.Unmarshal(response, &issuesSearch)
	log.Printf("Number of Issues for %s: %d\n", label, issuesSearch.TotalCount)
	return nil
}
