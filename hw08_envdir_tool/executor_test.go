package main

import (
	"testing"

	//nolint:depguard
	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	const envDir = "./testdata/env"
	env, err := ReadDir(envDir)
	require.Nil(t, err)
	testCmds := []string{"/bin/bash", "./testdata/echo.sh", "arg1=1", "arg2=2"} //"$(pwd)/testdata/env"
	rc := RunCmd(testCmds, env)
	require.Equal(t, 0, rc)
}
