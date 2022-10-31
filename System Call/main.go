package main

// Calling shell command

// cd to a given dir
// call ls -al
// print its output

import "fmt"
import "os"
// To make a system call, we use import "os/exec"
import "os/exec"

func main() {

	var dirToRun = "/Users/xah/web/"
	var err = os.Chdir(dirToRun)
	if err != nil {
		panic(err)
	}

	var cmdName = "ls"

	var cmd = exec.Command(cmdName, "-a", "-l")
	
	// We run it, wait and get output
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", string(output))

}
