package main

import (
	"testing"

	//nolint:depguard
	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	const envDir = "./testdata/env"
	env, err := ReadDir(envDir)
	require.Equal(t, err, nil)
	require.Equal(t, env["HELLO"].Value, `"hello"`)
	require.Equal(t, env["BAR"].Value, "bar")
	require.Equal(t, env["FOO"].Value, `   foo
with new line`)
	require.Equal(t, env["EMPTY"].Value, "")
	require.True(t, env["UNSET"].NeedRemove)

	const unsupportedEnvDir = "./testdata/unsupportedEnv"
	_, err2 := ReadDir(unsupportedEnvDir)
	require.Equal(t, err2, ErrUnsupported)
}
