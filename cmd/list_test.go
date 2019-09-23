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

func TestListCmd(t *testing.T) {

	var tstListInput = []struct {
		testCase string
		//input 	[]string
		expected  string
		isDeleted bool
	}{
		{"tc1", "Ignore", true},
		{"tc2", "Something Went Wrong : ", false},
	}

	tempAllTasks := AllTasks

	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")

	dbc, err := db.Init(dbPath)
	defer dbc.Close()
	if err != nil {
		t.Errorf("Error while connecting to db : %v", err)
	}

	f, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)

	oldStdout := os.Stdout
	os.Stdout = f

	for _, item := range tstListInput {

		if item.testCase == "tc2" {
			AllTasks = func() ([]db.Task, error) {
				return nil, errors.New("Something Went Wrong : ")
			}
		}
		list.Run(list, []string{})

		f.Seek(0, 0)

		content, err := ioutil.ReadAll(f)
		if err != nil {
			t.Errorf("Error While Reading file : %v", err)
		}
		output := string(content)

		val := strings.Contains(output, item.expected)

		assert.Equal(t, true, val, "They Must be Equal", item.isDeleted, output)

		f.Truncate(0)
		f.Seek(0, 0)
		AllTasks = tempAllTasks
	}
	os.Stdout = oldStdout

	f.Close()

}
