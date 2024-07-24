package main

import (
	"log"
	"os/exec"
	"path/filepath"
)

var reset = "\033[0m"

// var red = "\033[31m"
// var green = "\033[32m"
var yellow = "\033[33m"

// var blue = "\033[34m"
// var purple = "\033[35m"
// var cyan = "\033[36m"
// var gray = "\033[37m"
// var white = "\033[97m"

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getModuleRoot() (root string) {
	cmd := exec.Command("go", "env", "GOMOD")
	out, err := cmd.Output()
	die(err)
	return filepath.Dir(string(out)) + "/"
}
