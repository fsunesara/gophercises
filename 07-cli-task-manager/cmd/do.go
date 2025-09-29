package cmd

import (
	"fmt"
	"strconv"
	"task/db"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:  "do [task]",
	Run:  do,
	Args: cobra.ExactArgs(1),
}

func do(cmd *cobra.Command, args []string) {
	taskId, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Error: %s is not a valid integer", args[0])
		return
	}
	tasks := db.ListTasks()
	if taskId < 1 || taskId > len(tasks)+2 {
		fmt.Printf("Error: task %d not found", taskId)
		return
	}
	db.CompleteTask(tasks[taskId-1].Id)
	fmt.Printf("You have completed the %q task.", tasks[taskId-1].Text)
}

func init() {
	rootCmd.AddCommand(doCmd)
}
