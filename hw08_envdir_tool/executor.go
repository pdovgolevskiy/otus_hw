package main

import (
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	for envKey, enVal := range env {
		if enVal.NeedRemove {
			os.Unsetenv(envKey)
			continue
		}
		os.Setenv(envKey, enVal.Value)
	}
	cmdC := exec.Command(cmd[0], cmd[1:]...)
	cmdC.Stdout = os.Stdout
	cmdC.Stderr = os.Stderr
	err := cmdC.Run()
	if exitErr, ok := err.(*exec.ExitError); ok {
		exitCode := exitErr.ExitCode()
		return exitCode
	}
	return cmdC.ProcessState.ExitCode()
}
