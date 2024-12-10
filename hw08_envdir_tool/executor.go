package main

import (
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {

	//fmt.Println("cmd start")
	for envKey, enVal := range env {
		if enVal.NeedRemove {
			os.Unsetenv(envKey)
			continue
		}
		os.Setenv(envKey, enVal.Value)
	}
	cmdC := exec.Command(cmd[0], cmd[1:]...)
	//cmdC.Env = os.Environ()
	cmdC.Stdout = os.Stdout
	cmdC.Stderr = os.Stderr
	// Place your code here.
	err := cmdC.Run()
	if exitErr, ok := err.(*exec.ExitError); ok {
		exitCode := exitErr.ExitCode()
		return exitCode
	}
	return cmdC.ProcessState.ExitCode()
}
