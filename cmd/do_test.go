package cmd

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"task/db"
	"testing"

	"github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/assert"
)

var testDoInput = []struct {
	testCase string
	cmd      []string
	expected string
}{
	{"tc1", []string{"1"}, "Marked Task "},
	{"tc2", []string{"999"}, "Invalid Task Number : 999"},
	{"tc3", []string{"parsingError"}, "Invalid Option : "},
	{"tc4", []string{"2"}, "Could Not mark Task as Complete : "},
	//{"tc5", []string{"parsingError"}, "Invalid Option : "},
}

func TestDoCmd(t *testing.T) {
	tempAllTask := AllTasks
	tempDelTask := DelTask
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	//fmt.Println(dbPath)

	dbc, err := db.Init(dbPath)
	defer dbc.Close()
	if err != nil {
		t.Errorf("Error : %v", err)
	}

	f, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)

	oldStdout := os.Stdout
	os.Stdout = f

	for _, item := range testDoInput {

		if len(item.cmd) > 0 {
			if item.cmd[0] == "999" {
				AllTasks = func() ([]db.Task, error) {
					return nil, errors.New("Something went Wrong : ")
				}
			} else if item.cmd[0] == "2" {
				db.CreateTask("task1")
				db.CreateTask("task2")
				db.CreateTask("task3")
				DelTask = func(k int) error {
					return errors.New("Could Not mark Task as Complete : ")
				}
			}
		}

		do.Run(do, item.cmd)
		f.Seek(0, 0)

		content, err := ioutil.ReadAll(f)
		if err != nil {
			t.Errorf("Error while reading contents of file : %v", err)
		}

		output := string(content)
		val := strings.Contains(output, item.expected)
		assert.Equal(t, true, val, "Must be Equal", item.expected, output)

		f.Truncate(0)
		f.Seek(0, 0)
		AllTasks = tempAllTask
		DelTask = tempDelTask

	}

	os.Stdout = oldStdout
	f.Close()

}
