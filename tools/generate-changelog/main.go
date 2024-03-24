package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/qri-io/deepdiff"
)

func main() {
	var err error
	//args := os.Args[1:]
	// args := []string{
	// 	"../../data/provider-schema/2024.03.21.json",
	// 	"../../data/provider-schema/2024.03.22.json",
	// }

	dir := "../../data/provider-schema"
	dirEntry, err := os.ReadDir(dir)
	die(err)

	var files []string
	for _, e := range dirEntry {
		i, err := e.Info()
		die(err)

		files = append(files, fmt.Sprintf("%s/%s", dir, i.Name()))
	}

	for i := range len(files) - 1 {
		fmt.Println("\n====================================")

		// from := args[0]
		from := files[i]
		dateFrom := strings.TrimSuffix(filepath.Base(from), ".json")

		// to := args[1]
		to := files[i+1]
		dateTo := strings.TrimSuffix(filepath.Base(to), ".json")

		fmt.Println(dateFrom, "->", dateTo+"\n")

		//----------------
		// pre := `package main
		// import (jd "github.com/josephburnett/jd/lib"
		// "github.com/wI2L/jsondiff")
		// var _ = `

		// jsonDiffFile := "json-diff-" + dateFrom + "-" + dateTo + ".go"
		// err = os.WriteFile(jsonDiffFile, []byte(pre+tryJsondiff(from, to)), 0655)
		// die(err)

		//--------------------
		// pre := "# Changelog\n\n"
		// jdFile := "jd-" + dateFrom + "-" + dateTo + ".md"
		// os.WriteFile(jdFile, []byte(pre+tryJd(from, to)), 0644)
		// die(err)

		//--------------------
		pre := fmt.Sprintf("## Version %s\n", dateTo)
		jdFile := "tmp/deepdiff-" + dateFrom + "-" + dateTo + ".md"
		os.WriteFile(jdFile, []byte(pre+tryDeepDiff(from, to)), 0644)
		die(err)

		//--------------------
		// jdFile := "tmp/deepdiff-" + dateFrom + "-" + dateTo + ".json"
		// os.WriteFile(jdFile, []byte(tryDeepDiff(from, to)), 0644)
		// die(err)

	}

}

func tryDeepDiff(fileA, fileB string) string {
	previousSchemaByte, err := os.ReadFile(fileA)
	die(err)

	newSchemaByte, err := os.ReadFile(fileB)
	die(err)

	var previousSchema map[string]any
	err = json.Unmarshal(previousSchemaByte, &previousSchema)
	die(err)

	var newSchema map[string]any
	err = json.Unmarshal(newSchemaByte, &newSchema)
	die(err)

	dd := deepdiff.New(func(cfg *deepdiff.Config) { cfg.CalcChanges = true })

	rootDeltas, err := dd.Diff(context.TODO(), previousSchema, newSchema)
	die(err)

	var removeNoopDeltas func(d deepdiff.Deltas) deepdiff.Deltas
	removeNoopDeltas = func(d deepdiff.Deltas) deepdiff.Deltas {
		for i := len(d) - 1; i >= 0; i-- {
			d[i].Deltas = removeNoopDeltas(d[i].Deltas)
			if d[i].Deltas.Len() == 0 && d[i].Type == deepdiff.DTContext {
				d = slices.Delete(d, i, i+1)
			}
		}
		return d
	}
	rootDeltas = removeNoopDeltas(rootDeltas)

	var changes = make(map[string]map[string][]string)
	var organizeChanges func(d deepdiff.Deltas, path, kind string)
	organizeChanges = func(d deepdiff.Deltas, path, kind string) {
		for _, dd := range d {

			if dd.Path.String() == "provider_schemas" {
				organizeChanges(dd.Deltas[0].Deltas, path, kind)
				continue
			}

			subpath := path

			switch dd.Path.String() {
			case "resource_schemas":
				kind = "Resources"
			case "data_source_schemas":
				kind = "Data Sources"
			case "provider":
				kind = "Provider Config"
			case "block", "attributes", "nested_type":
				// Noop
			default:
				if path == "" {
					subpath = dd.Path.String()
				} else {
					subpath = subpath + "." + dd.Path.String()
				}
			}

			if changes[kind] == nil {
				changes[kind] = make(map[string][]string)
			}

			if dd.Type == deepdiff.DTContext {
				organizeChanges(dd.Deltas, subpath, kind)
				continue
			}

			switch dd.Type {
			case deepdiff.DTDelete:
				changes[kind]["Remove"] = append(changes[kind]["Remove"], subpath)
			case deepdiff.DTInsert:
				changes[kind]["Add"] = append(changes[kind]["Add"], subpath)
			case deepdiff.DTUpdate:
				changes[kind]["Change"] = append(changes[kind]["Change"], subpath)
			default:
				fmt.Println(dd)
				panic("unknown deepdiff type")
			}
		}
	}

	organizeChanges(rootDeltas, "", "")

	r := ""

	for kind, v := range changes {
		r += fmt.Sprintf("### %s\n", kind)
		for op, paths := range v {
			r += fmt.Sprintf("#### %s\n", op)
			for _, path := range paths {
				r += fmt.Sprintf("* %s\n\n", path)
			}
		}
	}

	return r
}

// AddLineBreaks is a very naive way to add line breaks
// to a .go file string. It came about as a solution when
// regex replace were too slow to be feasible
func AddLineBreaks(r string) string {
	rr := []rune(r)
	var pa rune
	var insideQuote bool
	var quote rune
	for pos, char := range rr {
		if insideQuote {
			if char == quote {
				if char == '`' {
					insideQuote = false
				} else if pa != '\\' {
					insideQuote = false
				}
			}
		} else {
			if char == '"' || char == '\'' || char == '`' {
				quote = char
				insideQuote = true
			} else {
				if char == ',' {
					// TODO change replacement to appending
					//  milestone:6
					//  this will need to be changed to a list
					//  of indexes where we wish to act, then loop over
					//  that list backwards after it is populated
					//  as we do not with to append to the slice while
					//  iterating over it in a forwards fashion
					rr[pos+1] = '\n'
				}
			}
		}
		pa = char
	}
	return string(rr)
}
