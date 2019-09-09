
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var list = &cobra.Command{
	Use:   "list",
	Short: "Lists All the Task To-Do",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list.go called")
	},
}

func init() {
	RootCmd.AddCommand(list)
}
