package main

import (
	"fmt"
	"os"
	"path/filepath"
	"task/cmd"
	"task/db"

	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	db.Init(dbPath)
	fmt.Println(dbPath)
	cmd.RootCmd.Execute()
}

func throw(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
