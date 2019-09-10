package main

import (
	"fmt"
	"task/db"
	"os"
	"path/filepath"
	"task/cmd"
	"github.com/mitchellh/go-homedir"
)

func main(){
	home,_ := homedir.Dir()
	dbPath := filepath.Join(home,"tasks.db")
	throw(db.Init(dbPath))
	fmt.Println("DB Conn")
	throw(cmd.RootCmd.Execute())
}

func throw(err error){
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
}
