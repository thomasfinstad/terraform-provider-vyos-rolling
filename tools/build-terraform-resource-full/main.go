package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/vyos/schema/interfacedefinition"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/vyos/vyosinterface"
)

func main() {
	args := os.Args[1:]
	rootOutputDirectory := args[0]
	rootPkgName := args[1]
	selfImportRoot := args[2]
	skipDirAbsNames := strings.Split(args[3], ",")

	_, thisFilename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Did not get path info")
	}
	thisDir := filepath.Dir(thisFilename)

	vyosInterfaces := vyosinterface.GetInterfaces()

	// Compile named TagNode based resources
	for _, vyosInterface := range vyosInterfaces {
		rootTagNodes, ok := vyosInterface.BaseTagNodes()
		if !ok {
			continue
		}

		for _, rootTagNode := range rootTagNodes {

			fmt.Printf("BaseTagNode: %s\n", rootTagNode.AbsName())

			absDirName := strings.Join(rootTagNode.AbsName(), "-")
			shouldSkip := false
			for _, skipDirAbsName := range skipDirAbsNames {
				if absDirName == skipDirAbsName {
					shouldSkip = true
					break
				}
			}
			if shouldSkip {
				fmt.Printf("\nWARNING: [%s] Has been marked for skipping\n", rootTagNode.AbsName())
				continue
			}

			// Create output dir
			resourceOutputDir := fmt.Sprintf("%s/%s", rootOutputDirectory, absDirName)
			if err := os.MkdirAll(resourceOutputDir, os.ModePerm); err != nil {
				log.Fatal(err)
			}

			// Compile template
			t, err := template.New("resource_generation").ParseFiles(filepath.Join(thisDir, "template.gotmpl"))
			if err != nil {
				die(err)
			}

			resourceModelSubDir := "resourcemodel"

			// Key/Value map of template-name = template-data
			templateRuns := map[string]any{
				"validate":                    rootTagNode,
				"resource-tagnode-based-full": rootTagNode,
				"metadata":                    rootTagNode,
				"schema":                      rootTagNode,
				"crud":                        rootTagNode,
			}

			for templateName, data := range templateRuns {
				outputFile := fmt.Sprintf(
					"%s/%s.go",
					resourceOutputDir,
					strings.Split(templateName, "-")[0],
				)
				fmt.Printf("Creating: %s\n", outputFile)
				file, err := os.Create(outputFile)
				if err != nil {
					return
				}
				defer file.Close()

				// Write Package
				err = t.ExecuteTemplate(file, "package", map[string]string{"caller": thisFilename, "pkg": strings.ToLower(rootPkgName + rootTagNode.BaseNameCG())})
				if err != nil {
					die(err)
				}

				// Write Imports
				var importExtra []string
				if templateName == "resource-tagnode-based-full" {
					importExtra = []string{
						fmt.Sprintf("%s/%s/%s", selfImportRoot, resourceOutputDir, resourceModelSubDir),
					}
				}
				err = t.ExecuteTemplate(file, "imports", importExtra)
				if err != nil {
					die(err)
				}

				// Write template data
				err = t.ExecuteTemplate(file, templateName, data)
				if err != nil {
					die(err)
				}
			}

			// Create Resource model in subdir
			resourceModelOutputDir := fmt.Sprintf("%s/%s", resourceOutputDir, resourceModelSubDir)
			if err := os.MkdirAll(resourceModelOutputDir, os.ModePerm); err != nil {
				log.Fatal(err)
			}

			var resourceModelGeneration func(node interfacedefinition.NodeParent)
			resourceModelGeneration = func(node interfacedefinition.NodeParent) {
				outputFile := fmt.Sprintf("%s/%s.go", resourceModelOutputDir, strings.Join(node.AbsName(), "-"))
				fmt.Printf("Creating: %s\n", outputFile)
				file, err := os.Create(outputFile)
				if err != nil {
					return
				}
				defer file.Close()

				// Write Package
				err = t.ExecuteTemplate(file, "package", map[string]string{"caller": thisFilename, "pkg": resourceModelSubDir})
				if err != nil {
					die(err)
				}

				// Write Imports
				err = t.ExecuteTemplate(file, "imports", nil)
				if err != nil {
					die(err)
				}

				// Print template data
				// for _, d := range node.GetChildren().LeafNodes() {
				// 	fmt.Printf("[%T]: %v :::: %v\n", node, node, d)
				// }

				// Write the resource model
				err = t.ExecuteTemplate(file, "resource-model", node)
				if err != nil {
					die(err)
				}

				// Recurse
				c := node.GetChildren()
				for _, n := range c.TagNodes() {
					resourceModelGeneration(n)
				}
				for _, n := range c.Nodes() {
					resourceModelGeneration(n)
				}
			}

			resourceModelGeneration(rootTagNode)

			fmt.Printf("Done...\n\n")
		}
	}

	// TODO Create Node based global resources
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
