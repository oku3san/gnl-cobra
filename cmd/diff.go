package cmd

import (
    "log"
    "os"
    "os/exec"

    "github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
    Use:   "diff",
    Short: "Run git diff command",
    Long:  "Run git status command",
    Run: func(cmd *cobra.Command, args []string) {
        command := exec.Command("git", "diff")
        command.Stdout = os.Stdout
        err := command.Run()
        if err != nil {
            log.Fatal("git diff の実行に失敗しました")
        }
    },
}

func init() {
    rootCmd.AddCommand(diffCmd)
}
