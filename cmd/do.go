package cmd

import (
	"fmt"
	"strconv"
	"task/db"

	"github.com/spf13/cobra"
)

var (
	AllTasks = db.AllTasks
	DelTask  = db.DeleteTask
)

var do = &cobra.Command{
	Use:   "do",
	Short: "Complete a Task",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int

		for _, args := range args {

			id, err := strconv.Atoi(args)
			if err != nil {
				fmt.Printf("Invalid Option : %v", args)
			} else {
				ids = append(ids, id)
			}
		}

		tasks, err := AllTasks()
		if err != nil {
			fmt.Println("Something went Wrong : ", err)
		}

		for _, id := range ids {
			// if id ==  {

			// }
			if id <= 0 || id > len(tasks) {
				fmt.Printf("Invalid Task Number : %d", id)
				continue
			}

			task := tasks[id-1]

			err := DelTask(task.Key)
			if err != nil {
				fmt.Println("Could Not mark Task as Complete : ", err)
			}
			fmt.Printf("Marked Task \"%s\" as Complete.", task.Value)
		}

	},
}

func init() {
	RootCmd.AddCommand(do)
}
