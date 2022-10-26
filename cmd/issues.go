package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/mdelapenya/github-metrics/formatters"
	"github.com/mdelapenya/github-metrics/github"
	"github.com/mdelapenya/github-metrics/log"
	"github.com/mdelapenya/github-metrics/types"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var Owner string
var Repository string

func init() {
	issuesCmd.PersistentFlags().StringVarP(&Owner, "owner", "O", "", "Github owner (organization or user)")
	issuesCmd.PersistentFlags().StringVarP(&Repository, "repository", "R", "", "Github repository")

	rootCmd.AddCommand(issuesCmd)

	issuesCmd.AddCommand(labelsCmd)
}

var issuesCmd = &cobra.Command{
	Use:   "issues",
	Short: "Subcommand for issues",
	Long:  "Subcommand to plug-in other subcommands related to issues",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if Owner == "" || Repository == "" {
			return fmt.Errorf("neither Owner or Repository cannot be empty")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		//
	},
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
		log.Fatal("failed to proceed the request: %v", err)
	}

	var labels []types.Label
	json.Unmarshal(response, &labels)

	log.Info("Processing labels", zap.Int("count", len(labels)))
	failedLabels := []string{}

	formatter := formatters.Get(Format, OutputFile)

	for _, label := range labels {
		lr, err := getLabel(label.Name)
		if err != nil {
			log.Warn("failed to get issues count. Will be retried later", zap.String("label", label.Name), zap.Error(err))
			failedLabels = append(failedLabels, label.Name)
			continue
		}
		formatter.Format(lr)
	}

	if len(failedLabels) > 0 {
		log.Info("Retrying failed labels", zap.Strings("labels", failedLabels))
		for _, failedLabel := range failedLabels {
			lr, err := getLabel(failedLabel)
			if err != nil {
				log.Warn("failed to get issues count. Continuing.", zap.String("label", failedLabel), zap.Error(err))
				continue
			}
			formatter.Format(lr)
		}
	}
}

func getLabel(label string) (*types.MetricResponse, error) {
	response, err := github.IssuesCountByLabel(Owner, Repository, label)
	if err != nil {
		return nil, err
	}

	var issuesSearch types.IssuesSearch
	json.Unmarshal(response, &issuesSearch)
	return &types.MetricResponse{Title: label, Count: issuesSearch.TotalCount}, nil
}
