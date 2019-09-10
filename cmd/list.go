
package cmd

import (
	"task/db"
	"fmt"

	"github.com/spf13/cobra"
)

var list = &cobra.Command{
	Use:   "list",
	Short: "Lists All the Task To-Do",
	Run: func(cmd *cobra.Command, args []string) {
	
		tasks, err := db.AllTasks()

		if err!= nil{
			fmt.Println("Something Went Wrong : ",err)
		}

		if len(tasks) == 0{
			fmt.Println("You Have no remaining tasks !!!!!!")
		}
		for k,v := range(tasks){
			fmt.Printf("%d. %s KEY=%d\n ",k+1,v.Value,k)
		} 
	},
}

func init() {
	RootCmd.AddCommand(list)
}
