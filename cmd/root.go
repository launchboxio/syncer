package cmd

import (
	"fmt"
	"github.com/launchboxio/syncer/pkg/syncer"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "syncer",
	Short: "Sync data between Launchbox and a Kubernetes cluster",
	Run: func(cmd *cobra.Command, args []string) {
		// Start an HTTP server for metrics service
		// Initiate the syncer service

		config := &syncer.Config{}
		sync := &syncer.Syncer{Config: config}

		if err := sync.Run(); err != nil {
			log.Fatal(err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
