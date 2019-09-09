package main

import (
	"fmt"
	"task/db"
	"path/filepath"
	"task/cmd"
	//"github.com/spf13/cobra"
	"github.com/mitchellh/go-homedir"
)

func main(){
	home,_ := homedir.Dir()
	dbPath := filepath.Join(home,"tasks.db")
	err := db.Init(dbPath); if err!=nil{
		panic(err)
	}
	fmt.Println("DB Conn")
	cmd.RootCmd.Execute()
}
