package main

import (
	"os"

	"github.com/kelsonic-networks/kelca/internal/cli"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "kelca",
		Short: "Kelsonic Networks certificate authority tool",
		Long:  `A Public Key Infrastructure (PKI) management tool for creating and managing certificates.`,
	}

	cli.RegisterCommands(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
