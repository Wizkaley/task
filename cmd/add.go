package cmd

import(
	"github.com/spf13/cobra"
	"fmt"
	"strings"	
	"task/db"
)

var addCmd = &cobra.Command{
	Use: "add",
	Short: "Adds a Task to the To-Do List",
	Run: func(cmd *cobra.Command, args [] string){
		// fmt.Println("Add Called")
		task := strings.Join(args," ")
		_,err := db.CreateTask(task)
		if err!= nil{
			fmt.Println("Something Went Wrong while Creating Task : ",err)
			return
		}
		fmt.Printf("Added \"%s\" to your To-Do List.",task)
	},
}


// Init ...
func init(){
	RootCmd.AddCommand(addCmd)
}