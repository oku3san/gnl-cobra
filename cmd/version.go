package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
    Use:   "version",
    Short: "Show version",
    Long:  `Show version`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("version 0.0.1")
    },
}

func init() {
    rootCmd.AddCommand(versionCmd)
}
