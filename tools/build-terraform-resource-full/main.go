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
	skipFileBaseNames := strings.Split(args[2], ",")

	_, thisFilename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Did not get path info")
	}
	thisDir := filepath.Dir(thisFilename)

	vyosInterfaces := vyosinterface.GetInterfaces()

	// Compile named TagNode based resources
	for _, vyosInterface := range vyosInterfaces {

		//rootNode := vyosInterface.GetRootNode()
		baseNodes, ok := vyosInterface.BaseTagNodes()
		if !ok {
			continue
		}

		for _, baseNode := range baseNodes {
			fmt.Printf("BaseTagNode: %s", baseNode.AbsName())

			nodeAncestroy := vyosInterface.GetAncestory(baseNode)
			var nodeAncestroyNames []string
			for _, n := range nodeAncestroy {
				fmt.Printf(" < %s", n.BaseName())
				nodeAncestroyNames = append(nodeAncestroyNames, n.BaseName())
			}

			baseNode.MutateWithAncestors(nodeAncestroy)
			baseNode.SetBaseName(strings.Join(nodeAncestroyNames, " "))

			fileBaseName := strings.ReplaceAll(baseNode.BaseName(), " ", "-")
			shouldSkip := false
			for _, skipFileBaseName := range skipFileBaseNames {
				if fileBaseName == skipFileBaseName {
					shouldSkip = true
					break
				}
			}
			if shouldSkip {
				fmt.Printf("\nWARNING: [%s] Has been marked for skipping\n", baseNode.AbsName())
				continue
			}

			// Create output file
			outputFile := fmt.Sprintf("%s/autogen-resource-%s.go", outputDirectory, fileBaseName)

			fmt.Printf(" Creating: %s ", outputFile)

			file, err := os.Create(outputFile)
			if err != nil {
				return
			}
			defer file.Close()

			// Compile template
			t, err := template.New("resource_generation").ParseFiles(filepath.Join(thisDir, "template.gotmpl"))
			if err != nil {
				die(err)
			}

			// Write package
			err = t.ExecuteTemplate(file, "package", map[string]string{"caller": thisFilename, "pkg": pkgName})
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
			err = t.ExecuteTemplate(file, "tagNodeResourcemodel", baseNode)
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
			err = t.ExecuteTemplate(file, "tagNodeSchema", baseNode)
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

	// TODO Create Node based global resources
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
