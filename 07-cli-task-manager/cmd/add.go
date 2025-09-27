package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [desc]",
	Short: "Add a task to the task list.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("wewewe")
	},
}

// func Execute() {
// 	taskText := strings.Join(addCmd.Args, " ")
// }
