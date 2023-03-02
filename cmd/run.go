package cmd

import (
    "log"
    "os"
    "os/exec"

    "github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
    Use:   "run",
    Short: "Run git add, commit and push",
    Long:  "Run git add, commit and push",
    Run: func(cmd *cobra.Command, args []string) {
        command := exec.Command("git", "add", "-A")
        command.Stdout = os.Stdout
        err := command.Run()
        if err != nil {
            log.Fatal("git add の実行に失敗しました")
        }

        command = exec.Command("git", "commit", "-m", "Update")
        command.Stdout = os.Stdout
        err = command.Run()
        if err != nil {
            log.Fatal("git commit の実行に失敗しました")
        }

        command = exec.Command("git", "push", "origin", "main")
        command.Stdout = os.Stdout
        err = command.Run()
        if err != nil {
            log.Fatal("git push の実行に失敗しました")
        }
    },
}

func init() {
    rootCmd.AddCommand(runCmd)

    // Here you will define your flags and configuration settings.

    // Cobra supports Persistent Flags which will work for this command
    // and all subcommands, e.g.:
    // runCmd.PersistentFlags().String("foo", "", "A help for foo")

    // Cobra supports local flags which will only run when this command
    // is called directly, e.g.:
    // runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
