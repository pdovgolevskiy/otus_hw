package main

import (
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {

	for envKey, enVal := range env {
		if enVal.NeedRemove == true {
			os.Unsetenv(envKey)
			continue
		}
		os.Setenv(envKey, enVal.Value)
	}
	cmdC := exec.Command(cmd[0], cmd[0:]...)
	// Place your code here.
	err := cmdC.Run()
	if exitErr, ok := err.(*exec.ExitError); ok {
		exitCode := exitErr.ExitCode()
		return exitCode
		//fmt.Println(exitCode)
	}
	return -1
}
