package main

import (
	"encoding/json"
	"os"

	"github.com/wI2L/jsondiff"
)

func main() {
	args := os.Args[1:]

	previousSchemaByte, err := os.ReadFile(args[0])
	die(err)

	newSchemaByte, err := os.ReadFile(args[1])
	die(err)

	var previousSchema map[string]any
	err = json.Unmarshal(previousSchemaByte, &previousSchema)
	die(err)

	var newSchema map[string]any
	err = json.Unmarshal(newSchemaByte, &newSchema)
	die(err)

	patch, err := jsondiff.Compare(previousSchema, newSchema)
	die(err)

	b, err := json.MarshalIndent(patch, "", "    ")
	die(err)

	os.Stdout.Write(b)
}
