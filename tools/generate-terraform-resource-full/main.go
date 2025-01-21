package main

import (
	"cmp"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
	"text/template"

	schemadefinition "github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/vyos/schemadefinition"
	vyosinterfaces "github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/vyos/vyosinterfaces"
)

type autogenTemplateInfo struct {
	PkgName        string
	PkgPath        string
	PkgConstructor string
}

func main() {
	args := os.Args[1:]
	gitRootDirectory := args[0]
	outputDirectory := args[1]
	selfImportRoot := args[2]
	skipDirAbsNames := strings.Split(args[3], ",")

	vyosInterfaces := vyosinterfaces.GetInterfaces()

	var packagesToGenerate []autogenTemplateInfo

	// Compile resources
	for _, vyosInterface := range vyosInterfaces {
		{ // Helpful for troubleshooting
			n, err := vyosInterface.GetRootNode()
			die(err)
			fmt.Printf("Generating from interface: %s\n", n.NodeNameAttr)
		}

		// Named (TagNode) resources
		TagNodes, ok := vyosInterface.BaseTagNodes()
		if ok {
			for _, tagNode := range TagNodes {
				packagesToGenerate = append(packagesToGenerate, namedResources(gitRootDirectory, tagNode, skipDirAbsNames, fmt.Sprintf("%s/named", outputDirectory), "named", selfImportRoot)...)
			}
		}

		// Global (Node) resources
		Nodes, ok := vyosInterface.BaseNodes()
		if ok {
			for _, node := range Nodes {
				packagesToGenerate = append(packagesToGenerate, globalResources(gitRootDirectory, node, skipDirAbsNames, fmt.Sprintf("%s/global", outputDirectory), "global", selfImportRoot)...)
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

	slices.SortFunc(packagesToGenerateDeduped, func(a, b autogenTemplateInfo) int { return cmp.Compare(a.PkgName, b.PkgName) })

	// Create package resource inclusion function
	_, callerFileName, _, ok := runtime.Caller(0)
	if !ok {
		panic("Did not get path info")
	}
	thisDir := filepath.Dir(callerFileName)
	packageOutputFile, err := filepath.Abs(filepath.FromSlash(fmt.Sprintf("%s/%s/package.go", gitRootDirectory, outputDirectory)))
	die(err)
	fmt.Printf("Creating autogen resource include file: %s\n", packageOutputFile)
	file, err := os.Create(packageOutputFile)
	if err != nil {
		return
	}
	defer file.Close()

	// Render and write file
	t, err := template.New("package_generation").ParseFiles(filepath.Join(thisDir, "templates", "package.gotmpl"))
	if err != nil {
		die(err)
	}
	err = t.ExecuteTemplate(file, "package", map[string]any{"pkgName": "autogen", "importRoot": strings.ToLower(selfImportRoot), "pkgs": packagesToGenerateDeduped})
	if err != nil {
		die(err)
	}

	// TODO autogenerate resource import doc
	//  milestone: 2
	//  making it easier for users to import resources by following the documentation
	//  explain how the resource ID is built up by using a template
	//  something like policy__access-list__<access-list id>__rule__<rule id>

	// TODO look into marking values as sensitive
	//  milestone: 3
	//  Is there any way to detect this from the shcema?
	//  If not a manual overwrite feature during code generation
	//  will be the next best thing.
	//  Example of sensitive value:
}

func namedResources(gitRootDirectory string, tagNode *schemadefinition.TagNode, skipDirAbsNames []string, outputDirectory string, rootPkgName string, selfImportRoot string) (pkgs []autogenTemplateInfo) {
	_, callerFileName, _, ok := runtime.Caller(0)
	if !ok {
		panic("Did not get path info")
	}
	callerFileName, err := filepath.Abs(callerFileName)
	die(err)
	thisDir, err := filepath.Abs(filepath.Dir(callerFileName))
	die(err)
	callerFileName = strings.TrimLeft(strings.TrimPrefix(callerFileName, gitRootDirectory), string(filepath.Separator))

	// fmt.Printf("BaseTagNode: %s\n", tagNode.AbsName())

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
	resourceOutputDir := strings.Join([]string{gitRootDirectory, outputDirectory, absDirName}, "/")
	if err := os.MkdirAll(resourceOutputDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	// Compile template
	templateFiles, err := filepath.Glob(filepath.Join(thisDir, "templates", "resources", "*", "*.gotmpl"))
	die(err)
	t, err := template.New("resource_generation").ParseFiles(templateFiles...)
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
		pkgs = append(pkgs, namedResourceGeneration(gitRootDirectory, resourceOutputDir, templateName, callerFileName, rootPkgName, tagNode, selfImportRoot, resourceModelSubDir, t, data))
	}

	// Create Resource model in subdir
	resourceModelOutputDir := strings.Join([]string{resourceOutputDir, resourceModelSubDir}, "/")
	if err := os.MkdirAll(resourceModelOutputDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	namedResourceModelGeneration(resourceModelOutputDir, tagNode, t, callerFileName, resourceModelSubDir)

	return pkgs
}

func globalResources(gitRootDirectory string, node *schemadefinition.Node, skipDirAbsNames []string, outputDirectory string, rootPkgName string, selfImportRoot string) (pkgs []autogenTemplateInfo) {
	_, callerFileName, _, ok := runtime.Caller(0)
	if !ok {
		panic("Did not get path info")
	}
	callerFileName, err := filepath.Abs(callerFileName)
	die(err)

	thisDir, err := filepath.Abs(filepath.Dir(callerFileName))
	die(err)

	callerFileName = strings.TrimLeft(strings.TrimPrefix(callerFileName, gitRootDirectory), string(filepath.Separator))

	// fmt.Printf("BaseNode: %s\n", node.AbsName())

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
	resourceOutputDir := strings.Join([]string{gitRootDirectory, outputDirectory, absDirName}, "/")
	if err := os.MkdirAll(resourceOutputDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	// Compile template
	templateFiles, err := filepath.Glob(filepath.Join(thisDir, "templates", "resources", "*", "*.gotmpl"))
	die(err)
	t, err := template.New("resource_generation").ParseFiles(templateFiles...)
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
		pkgs = append(pkgs, globalResourceGeneration(gitRootDirectory, resourceOutputDir, templateName, callerFileName, rootPkgName, node, selfImportRoot, resourceModelSubDir, t, data))
	}

	// Create Resource model in subdir
	resourceModelOutputDir := strings.Join([]string{resourceOutputDir, resourceModelSubDir}, "/")
	if err := os.MkdirAll(resourceModelOutputDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	globalResourceModelGeneration(resourceModelOutputDir, node, t, callerFileName, resourceModelSubDir)

	return pkgs
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func globalResourceGeneration(gitRootDirectory string, resourceOutputDir string, templateName string, callerFileName string, rootPkgName string, rootNode *schemadefinition.Node, selfImportRoot string, resourceModelSubDir string, t *template.Template, data any) (pkg autogenTemplateInfo) {
	// Format output file name
	outputFile := fmt.Sprintf(
		"%s.go",
		strings.Join(
			[]string{resourceOutputDir, strings.Split(templateName, "-")[0]},
			"/",
		),
	)

	// fmt.Printf("Creating global resource file: %s\n", outputFile)
	file, err := os.Create(outputFile)
	if err != nil {
		return
	}
	defer file.Close()

	pkg.PkgPath = strings.TrimLeft(strings.ReplaceAll(strings.TrimPrefix(filepath.FromSlash(resourceOutputDir), filepath.FromSlash(gitRootDirectory)), string(filepath.Separator)+string(filepath.Separator), string(filepath.Separator)), string(filepath.Separator))
	pkg.PkgName = strings.ToLower(rootPkgName + rootNode.BaseNameCG())
	pkg.PkgConstructor = "New" + rootNode.BaseNameCG()

	// Write Package
	err = t.ExecuteTemplate(file, "package", map[string]string{"caller": callerFileName, "pkg": pkg.PkgName})
	if err != nil {
		die(err)
	}

	// Write Imports
	var importExtra []string
	if templateName == "resource-node-based-full" {
		importExtra = []string{
			strings.ReplaceAll(
				filepath.FromSlash(
					fmt.Sprintf("%s/%s/%s",
						selfImportRoot,
						strings.TrimPrefix(resourceOutputDir, gitRootDirectory),
						resourceModelSubDir)),

				string(filepath.Separator)+string(filepath.Separator),
				string(filepath.Separator)),
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

func namedResourceGeneration(gitRootDirectory string, resourceOutputDir string, templateName string, callerFileName string, rootPkgName string, rootTagNode *schemadefinition.TagNode, selfImportRoot string, resourceModelSubDir string, t *template.Template, data any) (pkg autogenTemplateInfo) {
	// Format output file name
	outputFile := fmt.Sprintf(
		"%s.go",
		strings.Join(
			[]string{resourceOutputDir, strings.Split(templateName, "-")[0]},
			"/",
		),
	)
	// fmt.Printf("Creating named resource file: %s\n", outputFile)
	file, err := os.Create(outputFile)
	if err != nil {
		return
	}
	defer file.Close()

	pkg.PkgPath = strings.TrimLeft(strings.ReplaceAll(strings.TrimPrefix(filepath.FromSlash(resourceOutputDir), filepath.FromSlash(gitRootDirectory)), string(filepath.Separator)+string(filepath.Separator), string(filepath.Separator)), string(filepath.Separator))
	pkg.PkgName = strings.ToLower(rootPkgName + rootTagNode.BaseNameCG())
	pkg.PkgConstructor = "New" + rootTagNode.BaseNameCG()

	// Write Package
	err = t.ExecuteTemplate(file, "package", map[string]string{"caller": callerFileName, "pkg": pkg.PkgName})
	if err != nil {
		die(err)
	}

	// Write Imports
	var importExtra []string
	if templateName == "resource-tagnode-based-full" {
		importExtra = []string{
			strings.ReplaceAll(
				filepath.FromSlash(
					fmt.Sprintf("%s/%s/%s",
						selfImportRoot,
						strings.TrimPrefix(resourceOutputDir, gitRootDirectory),
						resourceModelSubDir)),

				string(filepath.Separator)+string(filepath.Separator),
				string(filepath.Separator)),
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

func namedResourceModelGeneration(resourceModelOutputDir string, node schemadefinition.NodeParent, t *template.Template, callerFileName string, resourceModelSubDir string) {
	outputFile := fmt.Sprintf(
		"%s.go",
		strings.Join(
			[]string{resourceModelOutputDir, strings.Join(node.AbsName(), "-")},
			"/",
		),
	)

	// fmt.Printf("Creating named resource model file: %s\n", outputFile)
	file, err := os.Create(outputFile)
	if err != nil {
		return
	}
	defer file.Close()

	// Write Package
	err = t.ExecuteTemplate(file, "package", map[string]string{"caller": callerFileName, "pkg": resourceModelSubDir})
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
		if n.GetIsBaseNode() {
			panic(fmt.Sprintf("A basenode of type Node should not be possible here as it is below a basenode of type TagNode, offending Node: %v", n.AbsName()))
		}
		namedResourceModelGeneration(resourceModelOutputDir, n, t, callerFileName, resourceModelSubDir)
	}
	for _, n := range c.TagNodes() {
		if n.GetIsBaseNode() {
			continue
		}
		namedResourceModelGeneration(resourceModelOutputDir, n, t, callerFileName, resourceModelSubDir)
	}
}

func globalResourceModelGeneration(resourceModelOutputDir string, node schemadefinition.NodeParent, t *template.Template, callerFileName string, resourceModelSubDir string) {
	outputFile := fmt.Sprintf(
		"%s.go",
		strings.Join(
			[]string{resourceModelOutputDir, strings.Join(node.AbsName(), "-")},
			"/",
		),
	)

	// fmt.Printf("Creating global resource model file: %s\n", outputFile)
	file, err := os.Create(outputFile)
	if err != nil {
		return
	}
	defer file.Close()

	// Write Package
	err = t.ExecuteTemplate(file, "package", map[string]string{"caller": callerFileName, "pkg": resourceModelSubDir})
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
		if n.GetIsBaseNode() {
			continue
		}
		globalResourceModelGeneration(resourceModelOutputDir, n, t, callerFileName, resourceModelSubDir)
	}
	for _, n := range c.TagNodes() {
		if n.GetIsBaseNode() {
			continue
		}
		globalResourceModelGeneration(resourceModelOutputDir, n, t, callerFileName, resourceModelSubDir)
	}
}
