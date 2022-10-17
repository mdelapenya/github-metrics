package main

import (
	"github.com/mdelapenya/github-metrics/cmd"
	"github.com/mdelapenya/github-metrics/log"
)

func main() {
	log.New()
	defer log.Sync() // flushes buffer, if any

	cmd.Execute()
}
