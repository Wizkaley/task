package db

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mitchellh/go-homedir"
)

var (
	dbErr     = errors.New("database not open")
	bucketErr = errors.New("Bucket not found")
)

func TestListTask(t *testing.T) {
	dir, _ := homedir.Dir()
	p := filepath.Join(dir, "tasks.db")
	db, err := Init(p)

	_, err = AllTasks()
	if err != nil {
		fmt.Println("Test Passed")
	}
	db.Close()
}

var testCreateTaskInput = []struct {
	input string
	err   error
}{
	{"task1", nil},
	{"dbClosed", dbErr},
	{"task2", nil},
	{"bucketErr", bucketErr},
	{"task3", nil},
}

func setDB() {
	dir, _ := homedir.Dir()
	p := filepath.Join(dir, "tasks.db")
	Init(p)
}
func TestCreateTask(t *testing.T) {

	for _, item := range testCreateTaskInput {

		if item.err != dbErr {
			setDB()
		}
		if item.err == bucketErr {
			taskBucket = []byte{}
		} else {
			taskBucket = []byte("tasks")
		}
		_, err := CreateTask(item.input)
		assert.Equalf(t, item.err, err, "Expected %q but got %q", item.err, err)
		db.Close()
	}
}

var testDeleteTaskInput = []struct {
	input    int
	expected error
}{
	{1, nil},
	{2, dbErr},
	{3, nil},
	{1, bucketErr},
	{2, nil},
}

func TestDeleteTask(t *testing.T) {
	//setDB()

	for _, item := range testDeleteTaskInput {
		if item.expected != dbErr {
			setDB()
		}
		if item.expected == bucketErr {
			taskBucket = []byte{}
		} else {
			taskBucket = []byte("tasks")
		}
		err := DeleteTask(item.input)
		assert.Equalf(t, item.expected, err, "Expected %v but got %v", item.expected, err)
		db.Close()
	}

}

// func TestDeleteTaskF(t *testing.T) {
// 	dir, _ := homedir.Dir()
// 	p := filepath.Join(dir, "tasks.db")
// 	db, err := Init(p)
// 	if err != nil {
// 		t.Errorf("Error while connecting to db : %v", err)
// 	}
// 	//getDb()
// 	err = DeleteTask(90900909009)
// 	fmt.Println(err)
// 	if err != nil {
// 		fmt.Println("Test Passed")
// 	}
// 	db.Close()

// }

// //TestInit ...
func TestInit(t *testing.T) {
	dir, _ := homedir.Dir()
	p := filepath.Join(dir, "tasks.db")

	db, err := Init(p)

	if err == nil && taskBucket != nil {
		fmt.Println("Connection Test Successfull \n Test Passed")

	} else {
		t.Errorf("Error While Iniatlising BoltDB : %v", err)
	}
	db.Close()

}
func TestInitF(t *testing.T) {

	p := filepath.Join("/", "tasks.db")

	_, err := Init(p)

	if err != nil {
		log.Println("Test Passed")
	}
}
