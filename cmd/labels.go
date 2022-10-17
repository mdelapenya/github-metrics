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

	failedLabels := []string{}

	for _, label := range labels {
		lr, err := getLabel(label.Name)
		if err != nil {
			log.Printf("failed to get issues count for '%s' label. Will be retried later. error: %v\n", label.Name, err)
			failedLabels = append(failedLabels, label.Name)
			continue
		}
		log.Printf("Number of Issues for %s: %d\n", lr.Label, lr.Count)
	}

	if len(failedLabels) > 0 {
		log.Printf("retrying failed labels: %v", failedLabels)
		for _, failedLabel := range failedLabels {
			lr, err := getLabel(failedLabel)
			if err != nil {
				log.Printf("failed to get issues count for '%s' label. Continuing. error: %v\n", failedLabel, err)
				continue
			}
			log.Printf("Number of Issues for %s: %d\n", lr.Label, lr.Count)
		}
	}
}

func getLabel(label string) (*types.LabelResponse, error) {
	response, err := github.IssuesCountByLabel(Owner, Repository, label)
	if err != nil {
		return nil, err
	}

	var issuesSearch types.IssuesSearch
	json.Unmarshal(response, &issuesSearch)
	return &types.LabelResponse{Label: label, Count: issuesSearch.TotalCount}, nil
}
