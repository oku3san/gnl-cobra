/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
    "log"
    "os"
    "os/exec"

    "github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
    Use:   "run",
    Short: "A brief description of your command",
    Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
