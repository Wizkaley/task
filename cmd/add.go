package cmd

import (
	"fmt"
	"strings"
	"task/db"

	"github.com/spf13/cobra"
)

var createTask = db.CreateTask

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a Task to the To-Do List",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("Add Called")
		if len(args) == 0 {
			fmt.Println("No argument provided with task")
			return
		}

		task := strings.Join(args, " ")
		_, err := createTask(task)
		if err != nil {
			fmt.Println("Something Went Wrong while Creating Task : ", err)
			return
		}
		fmt.Printf("Added \"%s\" to your To-Do List.", task)
	},
}

// Init ...
func init() {
	RootCmd.AddCommand(addCmd)
}
