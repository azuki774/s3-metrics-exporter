package main

import (
	"log/slog"

	"github.com/s3-metrics-exporter/internal/client"
	"github.com/s3-metrics-exporter/internal/service"
	"github.com/spf13/cobra"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		slog.Info("fetch called")

		s3client, err := client.NewS3Client()
		if err != nil {
			slog.Error("failed to create s3 client", err)
			return err
		}

		slog.Info("create s3 client")

		fetcher := service.Fetcher{
			S3Client: s3client,
		}

		fetcher.Fetch()
		if err != nil {
			slog.Error("failed to run fetcher", err)
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
