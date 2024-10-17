package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"text/template"

	vyosinterfaces "github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/vyos/vyosinterfaces"
)

type resourceInfo struct {
	Name      string
	ImportStr string
}

func main() {
	args := os.Args[1:]
	outputDirectory, err := filepath.Abs(args[0])
	die(err)

	_, callerFileName, _, ok := runtime.Caller(0)
	if !ok {
		panic("Did not get path info")
	}
	thisDir := filepath.Dir(callerFileName)

	vyosInterfaces := vyosinterfaces.GetInterfaces()

	// Compile resources
	for _, vyosInterface := range vyosInterfaces {
		{ // Helpful for troubleshooting
			n, err := vyosInterface.GetRootNode()
			die(err)
			fmt.Printf("Generating from interface: %s\n", n.NodeNameAttr)
		}

		// Named (TagNode) resources
		if TagNodes, ok := vyosInterface.BaseTagNodes(); ok {
			tn, err := template.New("import_doc").ParseFiles(filepath.Join(thisDir, "templates", "named.gotmpl"))
			if err != nil {
				die(err)
			}

			// Named resources: Render and write file
			for _, TagNode := range TagNodes {
				resName := "vyos_" + TagNode.AbsNameR()
				res := resourceInfo{
					Name:      resName,
					ImportStr: TagNode.ImportStr(),
				}

				err = os.MkdirAll(filepath.Join(outputDirectory, resName), os.FileMode(int(0755)))
				die(err)
				outputFile, err := os.Create(filepath.Join(outputDirectory, resName, "import.sh"))
				die(err)
				defer outputFile.Close()

				err = tn.ExecuteTemplate(outputFile, "import-named", res)
				die(err)
			}
		}

		// Global (Node) resources
		if Nodes, ok := vyosInterface.BaseNodes(); ok {
			tg, err := template.New("import_doc").ParseFiles(filepath.Join(thisDir, "templates", "global.gotmpl"))
			if err != nil {
				die(err)
			}

			// Global resources: Render and write file
			for _, Node := range Nodes {
				resName := "vyos_" + Node.AbsNameR()
				res := resourceInfo{
					Name:      resName,
					ImportStr: Node.ImportStr(),
				}

				err = os.MkdirAll(filepath.Join(outputDirectory, resName), os.FileMode(int(0755)))
				die(err)
				outputFile, err := os.Create(filepath.Join(outputDirectory, resName, "import.sh"))
				die(err)
				defer outputFile.Close()

				err = tg.ExecuteTemplate(outputFile, "import-global", res)
				die(err)
			}
		}
	}
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
