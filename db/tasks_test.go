package db

import (
	"path/filepath"
	"github.com/mitchellh/go-homedir"
	"testing"
	"fmt"
	"github.ibm.com/dash/dash_utils/dashtest"
)



func TestListTask(t * testing.T){
	dir,_ := homedir.Dir()
	p := filepath.Join(dir,"tasks.db")
	Init(p)

	_,err := AllTasks(); if err!=nil{
		fmt.Println("Test Passed")
	}
	db.Close()
}


func TestCreateTask(t * testing.T){
	Init("C:\\Users\\GS-1709\\tasks.db")
	
	str := "Ignore"

	 _,err := CreateTask(str); if err==nil{
		fmt.Println("Test Passed")
	 }else{
		 t.Errorf("Test Failed : %v",err)
	 }
	 db.Close()

}


func TestDeleteTask(t * testing.T){
	Init("C:\\Users\\GS-1709\\tasks.db")
	err := DeleteTask(2); if err!=nil{
		t.Errorf("Test Failed : %v",err)

	}else{
		fmt.Println("Test Passed")
	}
	db.Close()
}

func TestDeleteTaskF(t * testing.T){
	Init("C:\\Users\\GS-1709\\tasks.db")
	//getDb()
	err := DeleteTask(1000);if err !=nil{
		fmt.Println("Test Passed")
	}
	db.Close()
}


//TestInit ...
func TestInit(t *testing.T){
	dir,_ := homedir.Dir()
	p := filepath.Join(dir,"tasks.db")

	err := Init(p)

	if err == nil && taskBucket != nil{
		fmt.Print("Connection Test Successfull \n Test Passed")
		
	}else{
		t.Errorf("Error While Iniatlising BoltDB : %v",err)
	}

}

func TestInitF(t * testing.T){
	
	p := filepath.Join("C:/gocode/src/task","tasks.db")

	err := Init(p); if err != nil {
		t.Errorf("Test Passed : %v",err)
	}else{
		fmt.Println("Connection successfull \n Test Failed")
	}
}



func TestMain(M *testing.M){
	dashtest.ControlCoverage(M)
}