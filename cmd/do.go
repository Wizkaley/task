package cmd

import (
	"fmt"
	"strconv"
	"github.com/spf13/cobra"
)

var do = &cobra.Command{
	Use:   "do",
	Short: "Complete a Task",
	Run: func(cmd *cobra.Command, args []string) {
		var ids [] int
		for _,args := range(args){
			id,err := strconv.Atoi(args); if err!=nil{
				fmt.Printf("Could not convert : %v",args)
			}else{
				ids = append(ids,id)
			}
		}
		fmt.Println(ids)

	},
}

func init() {
	RootCmd.AddCommand(do)
}
