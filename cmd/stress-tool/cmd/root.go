package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"stress-tool/internal/adapters/cli"
	"stress-tool/internal/services"
)

var url string
var requests int
var concurrency int

var rootCmd = &cobra.Command{
	Use:   "stress-tool run",
	Short: "stress-tool is a tool to stress test a web server",
	Long:  "stress-tool is a tool to stress test a web server, it can be used to test the performance of a web server by simulating multiple users accessing the server at the same time.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cli.NewRootAdapter(services.NewStressService()).Run(url, requests, concurrency)
	},
}

func Execute() {
	rootCmd.Flags().StringVarP(&url, "url", "u", "", "The URL of the server to test")
	rootCmd.MarkFlagRequired("url")
	rootCmd.Flags().IntVarP(&requests, "requests", "r", 100, "The number of requests to make to the server")
	rootCmd.Flags().IntVarP(&concurrency, "concurrency", "c", 10, "The number of concurrent requests to make to the server")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
