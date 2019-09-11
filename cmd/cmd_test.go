package cmd

import(
	"testing"
	//"github.com/stretchr/testify/assert"
	"github.com/rendon/testcli"
	"github.com/spf13/cobra"
	"github.ibm.com/dash/dash_utils/dashtest"
	//"bytes"
	//"fmt"
)


func emptyRun(*cobra.Command,[] string){ }


// func TestListCommand(t *testing.T){
// 	c := testcli.Command("task","--help")
	
// 	c.Run()
	
// 	// if !c.Success(){
// 	// 	t.Fatalf("Expected to Succeed, but failed : %s",testcli.Error())
// 	// }
// 	if !c.StdoutContains("Task is a Simple To-Do Task Manager"){
// 		t.Fatalf("Expected %q to contain %q",testcli.Stdout(),"Task is a Simple To-Do Task Manager")
// 	}

// }

func TestListWithArg(t *testing.T){
	c := testcli.Command("task","list")
	//fmt.Println()
	c.Run()
	//fmt.Println(g)
	if !c.StdoutContains("1. Ignore KEY=0\n"){
		t.Fatalf("Expected %q to contain %q",c.Stdout(),"1. Ignore KEY=0\n")
	}
}

func TestAddWithArg(t *testing.T){
	c := testcli.Command("task","add","look up")
	c.Run()
	if !c.StdoutContains("Added"){
		t.Fatalf("Expected %q to contain %q",c.Stdout(),"Added")
	}

}

func TestDoWithArgs(t * testing.T){
	c := testcli.Command("task","do", "4")
	c.Run()
	if !c.StdoutContains("Marked Task"){
		t.Fatalf("Expected %q to contain %q",c.Stdout(),"Marked Task")
	}
}
	


func TestMain(m *testing.M){
	dashtest.ControlCoverage(m)
}