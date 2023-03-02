package cmd

import (
    "log"
    "os"
    "os/exec"

    "github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
    Use:   "status",
    Short: "Run git status command",
    Long:  "Run git status command",
    Run: func(cmd *cobra.Command, args []string) {
        command := exec.Command("git", "status")
        command.Stdout = os.Stdout
        err := command.Run()
        if err != nil {
            log.Fatal("git status の実行に失敗しました")
        }
    },
}

func init() {
    rootCmd.AddCommand(statusCmd)
}
