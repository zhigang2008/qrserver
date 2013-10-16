package main

import (
	"log"
	"os"
	"os/exec"
	"path"
)

func main() {
	file, _ := os.Getwd()
	log.Println("current path:", file)

	file, _ = exec.LookPath(os.Args[0])
	log.Println("exec path:", file)

	dir, _ := path.Split(file)
	log.Println("exec folder relative path:", dir)

	os.Chdir(dir)
	wd, _ := os.Getwd()
	log.Println("exec folder absolute path:", wd)
}
