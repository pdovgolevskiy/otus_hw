package main

import (
	"log"
	"os"
)

func main() {
	env, err := ReadDir(os.Args[1])
	if err != nil {
		log.Fatal(err.Error())
	}
	cmdArgs := os.Args[2:]
	rc := RunCmd(cmdArgs, env)
	os.Exit(rc)
}
