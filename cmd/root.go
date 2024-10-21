package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "selkieman-compose",
		Short: "Manage pods, containers and images",
		RunE:  root,
	}
)

func root(cmd *cobra.Command, _ []string) error {
	fmt.Println("test1")
	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
