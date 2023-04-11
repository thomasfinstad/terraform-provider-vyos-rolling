package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/vyos/vyosinterface"
)

func main() {
	args := os.Args[1:]
	outputDirectory := args[0]
	pkgName := args[1]

	_, thisFilename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Did not get path info")
	}
	thisDir := filepath.Dir(thisFilename)

	vyosInterfaces := vyosinterface.GetInterfaces()

	for _, vyosInterface := range vyosInterfaces {

		//rootNode := vyosInterface.GetRootNode()
		baseNodes := vyosInterface.BaseTagNodes()
		if baseNodes == nil {
			continue
		}

		for _, baseNode := range baseNodes {
			fmt.Printf("BaseTagNode: %s", baseNode.BaseName())

			nodeAncestroy := vyosInterface.GetAncestory(baseNode)
			var nodeAncestroyNames []string
			for _, n := range nodeAncestroy {
				fmt.Printf(" < %s", n.BaseName())
				nodeAncestroyNames = append(nodeAncestroyNames, n.BaseName())
			}

			// Create output file
			outputFile := fmt.Sprintf("%s/autogen-%s.go", outputDirectory, strings.Join(nodeAncestroyNames, "-"))

			fmt.Printf(" Creating: %s ", outputFile)

			file, err := os.Create(outputFile)
			if err != nil {
				return
			}
			defer file.Close()

			baseNode.SetBaseName(strings.Join(nodeAncestroyNames, " "))

			// Compile template
			t, err := template.New("resource_generation").ParseFiles(filepath.Join(thisDir, "template.gotmpl"))
			if err != nil {
				die(err)
			}

			// Write package
			err = t.ExecuteTemplate(file, "package", pkgName)
			if err != nil {
				die(err)
			}

			// Write imports
			err = t.ExecuteTemplate(file, "imports", baseNode)
			if err != nil {
				die(err)
			}

			// Validate resource
			err = t.ExecuteTemplate(file, "validate", baseNode)
			if err != nil {
				die(err)
			}

			// Write resource
			err = t.ExecuteTemplate(file, "resource", baseNode)
			if err != nil {
				die(err)
			}

			// Write resource model
			err = t.ExecuteTemplate(file, "resourcemodel", baseNode)
			if err != nil {
				die(err)
			}

			// Write metadata
			err = t.ExecuteTemplate(file, "metadata", baseNode)
			if err != nil {
				die(err)
			}

			// Method to initiate new resource
			err = t.ExecuteTemplate(file, "instantiation", baseNode)
			if err != nil {
				die(err)
			}

			// Write schema
			err = t.ExecuteTemplate(file, "schema", baseNode)
			if err != nil {
				die(err)
			}

			// CRUD Functions
			err = t.ExecuteTemplate(file, "crud", baseNode)
			if err != nil {
				die(err)
			}

			fmt.Println("Done")
		}
	}
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
