package cmd

import (
	"fmt"
	"strconv"
	"task/db"

	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm [task]",
	Short: "Remove a task.",
	Run:   rm,
	Args:  cobra.MinimumNArgs(1),
}

func rm(cmd *cobra.Command, args []string) {
	taskId, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Error: %s is not a valid integer", args[0])
		return
	}
	tasks := db.ListTasks(false)
	if taskId < 1 || taskId > len(tasks)+2 {
		fmt.Printf("Error: task %d not found", taskId)
		return
	}
	db.RemoveTask(tasks[taskId-1].Id)
	fmt.Printf("You have deleted the %q task.", tasks[taskId-1].Text)
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
