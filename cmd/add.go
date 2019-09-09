package cmd

import(
	"github.com/spf13/cobra"
	"fmt"
	"strings"	
)

var addCmd = &cobra.Command{
	Use: "add",
	Short: "Adds a Task to the To-Do List",
	Run: func(cmd *cobra.Command, args [] string){
		// fmt.Println("Add Called")
		task := strings.Join(args," ")
		fmt.Printf("Added \"%s\" to your To-Do List.",task)
	},
}


// Init ...
func init(){
	RootCmd.AddCommand(addCmd)
}