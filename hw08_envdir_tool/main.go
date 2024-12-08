package main

import (
	"log"
	"os"
)

func main() {
	const dirStr = "./testdata/env"
	env, err := ReadDir(dirStr)
	if err != nil {
		log.Fatal(err.Error())
	}
	rc := RunCmd(os.Args, env)
	os.Exit(rc)
}
