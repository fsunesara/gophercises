package cmd

import (
	"fmt"
	"strings"
	"task/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [desc]",
	Short: "Add a task to the task list.",
	Run:   add,
	Args:  cobra.MinimumNArgs(1),
}

func add(cmd *cobra.Command, args []string) {
	taskText := strings.Join(args, " ")
	db.AddTask(db.Task{Text: taskText})
	fmt.Printf("Added %q to your task list.", taskText)
}

func init() {
	rootCmd.AddCommand(addCmd)
}
