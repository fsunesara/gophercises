package cmd

import (
	"fmt"
	"task/db"

	"github.com/spf13/cobra"
)

var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "View tasks completed within the past 24 hours.",
	Run:   completed,
}

func completed(cmd *cobra.Command, args []string) {
	tasks := db.ListTasks(true)
	if len(tasks) == 0 {
		fmt.Println("You have not completed any tasks in the past 24 hours.")
		return
	}
	fmt.Println("You have completed the following tasks within the past 24 hours:")
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task.Text)
	}
}

func init() {
	rootCmd.AddCommand(completedCmd)
}
