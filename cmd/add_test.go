package cmd

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"task/db"
	"testing"

	"github.ibm.com/dash/dash_utils/dashtest"

	"github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/assert"
)

// func TestMain(m *testing.M) {
// 	dashtest.ControlCoverage(m)
// }

func SetUpDb() {
	db.Init("/home/wiz/tasks.db")
}

var test = []struct {
	cmd      []string
	expected string
}{
	{[]string{}, "No argument provided with task"},
	{[]string{"walk the dog"}, "Added \"walk the dog\" to your To-Do List."},
	{[]string{"errorTask"}, "Something Went Wrong while Creating Task"},
	// {[]string{"Learn go lang"}, "Added \"Learn go lang\" to youryour To-Do List."},
}

func TestAddCmd(t *testing.T) {
	temp := createTask
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")

	//Open File for comparison
	f, err := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)

	if err != nil {
		t.Errorf("Error in Initializing Database : %v", err)
	}

	//Initialize Database
	dbc, _ := db.Init(dbPath)
	//SetUpDb()
	//shift stdout
	oldStdout := os.Stdout
	os.Stdout = f
	//run command

	//a := []string{"Walk the Dog"}

	for _, item := range test {

		if len(item.cmd) > 0 {
			if item.cmd[0] == "errorTask" {
				createTask = func(str string) (int, error) {
					return -1, errors.New("Something Went Wrong while Creating Task")
				}
			}
		}
		addCmd.Run(addCmd, item.cmd)

		//seek file

		f.Seek(0, 0)
		//Read file content
		content, err := ioutil.ReadAll(f)
		if err != nil {
			t.Errorf("Error while reading contents of file : %v", err)
		}
		//check if output contains
		output := string(content)
		val := strings.Contains(output, item.expected)

		//assert.Equals(t, true, val, "Must Contain Substr")
		assert.Equal(t, true, val, "They Should be Equal")
		f.Truncate(0)
		f.Seek(0, 0)
		createTask = temp
	}

	os.Stdout = oldStdout
	f.Close()
	dbc.Close()
}

func TestMain(m *testing.M) {
	dashtest.ControlCoverage(m)
}
