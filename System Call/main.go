package main

// example of calling shell command

// cd to a given dir
// call ls -al
// print its output

import "fmt"
import "os"
import "os/exec"

func main() {

	var dirToRun = "/Users/xah/web/"
	var err = os.Chdir(dirToRun)
	if err != nil {
		panic(err)
	}

	var cmdName = "ls"

	var cmd = exec.Command(cmdName, "-a", "-l")

	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", string(output))

}
