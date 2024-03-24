package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"

	vyosinterfaces "github.com/thomasfinstad/terraform-provider-vyos/.build/vyosinterfaces"
	schemadefinition "github.com/thomasfinstad/terraform-provider-vyos/internal/vyos/schemadefinition"
)

type autogenTemplateInfo struct {
	PkgName        string
	PkgPath        string
	PkgConstructor string
}

func main() {
	args := os.Args[1:]
	rootOutputDirectory := args[0]
	selfImportRoot := args[1]
	skipDirAbsNames := strings.Split(args[2], ",")

	vyosInterfaces := vyosinterfaces.GetInterfaces()

	var packagesToGenerate []autogenTemplateInfo

	// Compile resources
	for _, vyosInterface := range vyosInterfaces {

		// Named (TagNode) resources
		TagNodes, ok := vyosInterface.TagNodes()
		if ok {
			for _, tagNode := range TagNodes {
				packagesToGenerate = append(packagesToGenerate, namedResources(tagNode, skipDirAbsNames, fmt.Sprintf("%s/named", rootOutputDirectory), "named", selfImportRoot)...)
			}
		}

		// Global (Node) resources
		Nodes, ok := vyosInterface.Nodes()
		if ok {
			for _, node := range Nodes {
				packagesToGenerate = append(packagesToGenerate, globalResources(node, skipDirAbsNames, fmt.Sprintf("%s/global", rootOutputDirectory), "global", selfImportRoot)...)
			}
		}
	}

	//sort.Slice(pkgs, func(i, j int) bool { return reflect.DeepEqual(pkgs[i], pkgs[j]) })
	//slices.Compact[[]autogenTemplateInfo](pkgs)
	// ugly ugly boy needs to remove duplicates of himself, but is too scared of the mirror
	var packagesToGenerateDeduped []autogenTemplateInfo
	for _, currentPkg := range packagesToGenerate {
		matchedConstructor := false
		matchedName := false
		matchedPath := false
		for _, preHandledPackage := range packagesToGenerateDeduped {
			if currentPkg.PkgConstructor == preHandledPackage.PkgConstructor {
				matchedConstructor = true
			}
			if currentPkg.PkgName == preHandledPackage.PkgName {
				matchedName = true
			}
			if currentPkg.PkgPath == preHandledPackage.PkgPath {
				matchedPath = true
			}
			if matchedConstructor || matchedName || matchedPath {
				if !(matchedConstructor && matchedName && matchedPath) {
					fmt.Println("Found duplicate, but not all fields match!")
					fmt.Printf("npkg: %#v \t opkg: %#v", preHandledPackage, currentPkg)
				}
				break
			}
		}
		if matchedConstructor || matchedName || matchedPath {
			continue
		}
		packagesToGenerateDeduped = append(packagesToGenerateDeduped, currentPkg)
	}

	// Create package resource inclusion function
	_, thisFilename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Did not get path info")
	}
	thisDir := filepath.Dir(thisFilename)
	outputFile := fmt.Sprintf("%s/package.go", rootOutputDirectory)
	fmt.Printf("Creating autogen resource include file: %s\n", outputFile)
	file, err := os.Create(outputFile)
	if err != nil {
		return
	}
	defer file.Close()

	// Render and write file
	t, err := template.New("package_generation").ParseFiles(filepath.Join(thisDir, "package-resource-include.gotmpl"))
	if err != nil {
		die(err)
	}
	err = t.ExecuteTemplate(file, "package", map[string]any{"pkgName": "autogen", "importRoot": strings.ToLower(selfImportRoot), "pkgs": packagesToGenerateDeduped})
	if err != nil {
		die(err)
	}

	// TODO autogenerate resource import doc
	//  milestone:2
	//  making it easier for users to import resources by following the documentation
	//  explain how the resource ID is built up by using a template
	//  something like policy__access-list__<access-list id>__rule__<rule id>

	// TODO look into marking values as sensitive
	//  milestone:3
	//  Is there any way to detect this from the shcema?
	//  If not a manual overwrite feature during code generation
	//  will be the next best thing.
	//  Example of sensitive value:
}

func namedResources(tagNode *schemadefinition.TagNode, skipDirAbsNames []string, rootOutputDirectory string, rootPkgName string, selfImportRoot string) (pkgs []autogenTemplateInfo) {
	_, thisFilename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Did not get path info")
	}
	thisDir := filepath.Dir(thisFilename)

	fmt.Printf("BaseTagNode: %s\n", tagNode.AbsName())

	// Absolute dir name
	absDirNameComponents := []string{tagNode.AbsName()[0], strings.Join(tagNode.AbsName()[1:], "-")}
	for i := len(absDirNameComponents) - 1; i > 0; i-- {
		if absDirNameComponents[i] == "" {
			absDirNameComponents = append(absDirNameComponents[:i], absDirNameComponents[i+1:]...)
		}
	}
	absDirName := strings.Join(absDirNameComponents, "/")

	// Check if blacklisted
	shouldSkip := false
	for _, skipDirAbsName := range skipDirAbsNames {
		if absDirName == skipDirAbsName {
			shouldSkip = true
			break
		}
	}
	if shouldSkip {
		fmt.Printf("\nWARNING: [%s] Has been marked for skipping\n", tagNode.AbsName())
		return
	}

	// Create output dir
	resourceOutputDir := strings.Join([]string{rootOutputDirectory, absDirName}, "/")
	if err := os.MkdirAll(resourceOutputDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	// Compile template
	t, err := template.New("resource_generation").ParseFiles(filepath.Join(thisDir, "named-template.gotmpl"))
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
		pkgs = append(pkgs, namedResourceGeneration(resourceOutputDir, templateName, thisFilename, rootPkgName, tagNode, selfImportRoot, resourceModelSubDir, t, data))
	}

	// Create Resource model in subdir
	resourceModelOutputDir := strings.Join([]string{resourceOutputDir, resourceModelSubDir}, "/")
	if err := os.MkdirAll(resourceModelOutputDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	namedResourceModelGeneration(resourceModelOutputDir, tagNode, t, thisFilename, resourceModelSubDir)

	fmt.Printf("Done...\n\n")

	return pkgs
}

func globalResources(node *schemadefinition.Node, skipDirAbsNames []string, rootOutputDirectory string, rootPkgName string, selfImportRoot string) (pkgs []autogenTemplateInfo) {
	_, thisFilename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Did not get path info")
	}
	thisDir := filepath.Dir(thisFilename)

	fmt.Printf("BaseNode: %s\n", node.AbsName())

	// Absolute dir name
	absDirNameComponents := []string{node.AbsName()[0], strings.Join(node.AbsName()[1:], "-")}
	for i := len(absDirNameComponents) - 1; i > 0; i-- {
		if absDirNameComponents[i] == "" {
			absDirNameComponents = append(absDirNameComponents[:i], absDirNameComponents[i+1:]...)
		}
	}
	absDirName := strings.Join(absDirNameComponents, "/")

	// Check if blacklisted
	shouldSkip := false
	for _, skipDirAbsName := range skipDirAbsNames {
		if absDirName == skipDirAbsName {
			shouldSkip = true
			break
		}
	}
	if shouldSkip {
		fmt.Printf("\nWARNING: [%s] Has been marked for skipping\n", node.AbsName())
		return
	}

	// Create output dir
	resourceOutputDir := strings.Join([]string{rootOutputDirectory, absDirName}, "/")
	if err := os.MkdirAll(resourceOutputDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	// Compile template
	t, err := template.New("resource_generation").ParseFiles(filepath.Join(thisDir, "global-template.gotmpl"))
	if err != nil {
		die(err)
	}

	resourceModelSubDir := "resourcemodel"

	// Key/Value map of template-name = template-data
	templateRuns := map[string]any{
		"validate":                 node,
		"resource-node-based-full": node,
		"metadata":                 node,
		"schema":                   node,
		"crud":                     node,
	}

	for templateName, data := range templateRuns {
		pkgs = append(pkgs, globalResourceGeneration(resourceOutputDir, templateName, thisFilename, rootPkgName, node, selfImportRoot, resourceModelSubDir, t, data))
	}

	// Create Resource model in subdir
	resourceModelOutputDir := strings.Join([]string{resourceOutputDir, resourceModelSubDir}, "/")
	if err := os.MkdirAll(resourceModelOutputDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	globalResourceModelGeneration(resourceModelOutputDir, node, t, thisFilename, resourceModelSubDir)

	fmt.Printf("Done...\n\n")

	return pkgs
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func globalResourceGeneration(resourceOutputDir string, templateName string, thisFilename string, rootPkgName string, rootNode *schemadefinition.Node, selfImportRoot string, resourceModelSubDir string, t *template.Template, data any) (pkg autogenTemplateInfo) {
	// Format output file name
	outputFile := fmt.Sprintf(
		"%s.go",
		strings.Join(
			[]string{resourceOutputDir, strings.Split(templateName, "-")[0]},
			"/",
		),
	)
	fmt.Printf("Creating global resource: %s\n", outputFile)
	file, err := os.Create(outputFile)
	if err != nil {
		return
	}
	defer file.Close()

	pkg.PkgPath = resourceOutputDir
	pkg.PkgName = strings.ToLower(rootPkgName + rootNode.BaseNameCG())
	pkg.PkgConstructor = "New" + rootNode.BaseNameCG()

	// Write Package
	err = t.ExecuteTemplate(file, "package", map[string]string{"caller": thisFilename, "pkg": pkg.PkgName})
	if err != nil {
		die(err)
	}

	// Write Imports
	var importExtra []string
	if templateName == "resource-node-based-full" {
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

	return pkg
}

func namedResourceGeneration(resourceOutputDir string, templateName string, thisFilename string, rootPkgName string, rootTagNode *schemadefinition.TagNode, selfImportRoot string, resourceModelSubDir string, t *template.Template, data any) (pkg autogenTemplateInfo) {
	// Format output file name
	outputFile := fmt.Sprintf(
		"%s.go",
		strings.Join(
			[]string{resourceOutputDir, strings.Split(templateName, "-")[0]},
			"/",
		),
	)
	fmt.Printf("Creating named resource: %s\n", outputFile)
	file, err := os.Create(outputFile)
	if err != nil {
		return
	}
	defer file.Close()

	pkg.PkgPath = resourceOutputDir
	pkg.PkgName = strings.ToLower(rootPkgName + rootTagNode.BaseNameCG())
	pkg.PkgConstructor = "New" + rootTagNode.BaseNameCG()

	// Write Package
	err = t.ExecuteTemplate(file, "package", map[string]string{"caller": thisFilename, "pkg": pkg.PkgName})
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

	return pkg
}

func namedResourceModelGeneration(resourceModelOutputDir string, node schemadefinition.NodeParent, t *template.Template, thisFilename string, resourceModelSubDir string) {
	outputFile := fmt.Sprintf(
		"%s.go",
		strings.Join(
			[]string{resourceModelOutputDir, strings.Join(node.AbsName(), "-")},
			"/",
		),
	)
	fmt.Printf("Creating named resource model: %s\n", outputFile)
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

	// Write the resource model
	err = t.ExecuteTemplate(file, "resource-model", node)
	if err != nil {
		die(err)
	}

	// Recurse
	c := node.GetChildren()
	for _, n := range c.Nodes() {
		namedResourceModelGeneration(resourceModelOutputDir, n, t, thisFilename, resourceModelSubDir)
	}
}

func globalResourceModelGeneration(resourceModelOutputDir string, node schemadefinition.NodeParent, t *template.Template, thisFilename string, resourceModelSubDir string) {
	outputFile := fmt.Sprintf(
		"%s.go",
		strings.Join(
			[]string{resourceModelOutputDir, strings.Join(node.AbsName(), "-")},
			"/",
		),
	)
	fmt.Printf("Creating global resource model: %s\n", outputFile)
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

	// Write the resource model
	err = t.ExecuteTemplate(file, "resource-model", node)
	if err != nil {
		die(err)
	}
}
