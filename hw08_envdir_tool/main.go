package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("not enought aruguments")
	}
	envs, err := ReadDir(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(RunCmd(os.Args[2:], envs))
}
