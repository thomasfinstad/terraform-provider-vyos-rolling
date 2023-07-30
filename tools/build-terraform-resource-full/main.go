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
		TagNodes, ok := vyosInterface.TagNodes()
		if !ok {
			continue
		}

		for _, tagNode := range TagNodes {

			fmt.Printf("BaseTagNode: %s\n", tagNode.AbsName())

			// Check if blacklisted
			absDirName := fmt.Sprintf("%s/%s", tagNode.AbsName()[0], strings.Join(tagNode.AbsName()[1:], "-"))
			shouldSkip := false
			for _, skipDirAbsName := range skipDirAbsNames {
				if absDirName == skipDirAbsName {
					shouldSkip = true
					break
				}
			}
			if shouldSkip {
				fmt.Printf("\nWARNING: [%s] Has been marked for skipping\n", tagNode.AbsName())
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
				"validate":                    tagNode,
				"resource-tagnode-based-full": tagNode,
				"metadata":                    tagNode,
				"schema":                      tagNode,
				"crud":                        tagNode,
			}

			for templateName, data := range templateRuns {
				resourceGeneration(resourceOutputDir, templateName, thisFilename, rootPkgName, tagNode, selfImportRoot, resourceModelSubDir, t, data)
			}

			// Create Resource model in subdir
			resourceModelOutputDir := fmt.Sprintf("%s/%s", resourceOutputDir, resourceModelSubDir)
			if err := os.MkdirAll(resourceModelOutputDir, os.ModePerm); err != nil {
				log.Fatal(err)
			}

			resourceModelGeneration(resourceModelOutputDir, tagNode, t, thisFilename, resourceModelSubDir)

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

func resourceGeneration(resourceOutputDir string, templateName string, thisFilename string, rootPkgName string, rootTagNode *interfacedefinition.TagNode, selfImportRoot string, resourceModelSubDir string, t *template.Template, data any) {
	// Format outpout file name
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

func resourceModelGeneration(resourceModelOutputDir string, node interfacedefinition.NodeParent, t *template.Template, thisFilename string, resourceModelSubDir string) {
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
	// for _, n := range c.TagNodes() {
	// 	resourceModelGeneration(resourceModelOutputDir, n, t, thisFilename, resourceModelSubDir)
	// }
	for _, n := range c.Nodes() {
		resourceModelGeneration(resourceModelOutputDir, n, t, thisFilename, resourceModelSubDir)
	}
}
