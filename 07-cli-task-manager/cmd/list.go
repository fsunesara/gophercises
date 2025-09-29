package cmd

import (
	"fmt"
	"task/db"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use: "list",
	Run: list,
}

func list(cmd *cobra.Command, args []string) {
	tasks := db.ListTasks()
	if len(tasks) == 0 {
		fmt.Println("You have no tasks.")
		return
	}
	fmt.Println("You have the following tasks:")
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task.Text)
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
}
