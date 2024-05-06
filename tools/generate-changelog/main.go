package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/gdexlab/go-render/render"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/hashicorp/go-version"
	"github.com/leodido/go-conventionalcommits"
	"github.com/leodido/go-conventionalcommits/parser"
	"github.com/qri-io/deepdiff"
)

const rollingrelease bool = true

//  Ref: https://developer.hashicorp.com/terraform/plugin/best-practices/versioning#versioning-specification
//  Ref: https://developer.hashicorp.com/terraform/plugin/best-practices/versioning#changelog-specification
//  Ref: https://developer.hashicorp.com/terraform/cli/v1.5.x/commands/providers/schema

// change contains either a git change or a schema change
type change struct {
	conventionalCommit *conventionalcommits.ConventionalCommit
	schemaChange       *schemaChg
}

func (c change) IsBreakingChange() bool {
	if c.FromSchema() {
		return c.schemaChange.chgLvl == schemaChgLvlBreak
	}
	return c.conventionalCommit.IsBreakingChange()
}

func (c change) IsNote() bool {
	if c.FromSchema() {
		return c.schemaChange.chgLvl == schemaChgLvlNote
	}
	return (!c.conventionalCommit.Ok() ||
		c.conventionalCommit.Type == "docs" ||
		c.conventionalCommit.Type == "refactor")
}

func (c change) IsFeature() bool {
	if c.FromSchema() {
		return c.schemaChange.chgLvl == schemaChgLvlFeat
	}
	return c.conventionalCommit.IsFeat()
}

func (c change) IsFix() bool {
	if c.FromSchema() {
		return c.schemaChange.chgLvl == schemaChgLvlFix
	}
	return c.conventionalCommit.IsFix()
}

/*
For provider development typically the "subsystem" is the resource or data source affected e.g. resource/load_balancer,
or provider if the change affects whole provider (e.g. authentication logic).
Each bullet also references the corresponding pull request number that contained the code changes,
in the format of [GH-####] (for HashiCorp released plugins, this will be automatically updated on release).

* resource/load_balancer: Add `ATTRIBUTE` argument (support X new functionality) [GH-12]
* resource/subnet: Now better [GH-22, GH-32]
*/
func (c change) Description() string {
	if c.FromSchema() {
		if c.schemaChange.isMeta {
			return fmt.Sprintf(
				"%s/`%s` %s attribute `%s` field `%s`",
				string(c.schemaChange.schemaType),
				c.schemaChange.field[0],
				c.schemaChange.changeType,
				strings.Join(c.schemaChange.field[1:len(c.schemaChange.field)-2], "."),
				strings.Join(c.schemaChange.field[len(c.schemaChange.field)-1:], "."),
			)
		}

		return fmt.Sprintf(
			"%s/`%s` %s attribute `%s`",
			string(c.schemaChange.schemaType),
			c.schemaChange.field[0],
			c.schemaChange.changeType,
			strings.Join(c.schemaChange.field[1:], "."),
		)
	}
	return c.conventionalCommit.Description
}

func (c change) FromSchema() bool {
	return c.conventionalCommit == nil
}

type schemaType string

const (
	schemaTypeResource   schemaType = "Resource"
	schemaTypeDataSource schemaType = "Data Source"
	schemaTypeProvider   schemaType = "Provider Config"
)

type schemaChgType string

const (
	schemaChgTypeAdd    schemaChgType = "Added"
	schemaChgTypeChange schemaChgType = "Changed"
	schemaChgTypeDel    schemaChgType = "Removed"
)

type schemaChgLvl int

const (
	schemaChgLvlBreak schemaChgLvl = iota
	schemaChgLvlNote
	schemaChgLvlFeat
	schemaChgLvlFix
)

type schemaChg struct {
	schemaType schemaType
	field      []string
	isMeta     bool
	changeType schemaChgType
	chgLvl     schemaChgLvl
}

type chglog struct {
	msgs            []change
	previousVersion version.Version
}

func (c chglog) Version() *version.Version {
	var err error

	major := c.previousVersion.Segments()[0]
	minor := c.previousVersion.Segments()[1]
	patch := c.previousVersion.Segments()[2]

	isMajor := false
	isMinor := false

	var newVersion *version.Version

	for _, m := range c.msgs {
		if m.IsBreakingChange() {
			isMajor = true
			isMinor = false
			break
		}

		if m.IsFeature() {
			isMinor = true
		}
	}

	if isMajor {
		major = major + 1
	} else if isMinor {
		minor = minor + 1
	} else if rollingrelease {
		b, err := os.ReadFile("../../data/vyos-1x-info.txt")
		die(err)
		patch, err = strconv.Atoi(strings.ReplaceAll(strings.Split(string(b), "T")[0], "-", ""))
		die(err)
	} else {
		patch = patch + 1
	}

	newVersion, err = version.NewVersion(fmt.Sprintf("%d.%d.%d", major, minor, patch))
	die(err)
	return newVersion
}

func (c chglog) Date() string {
	return time.Now().UTC().Format("2006-01-02 15:04:05 MST")
}

func (c *chglog) addCc(m conventionalcommits.Message) {
	if mc, ok := m.(*conventionalcommits.ConventionalCommit); ok {
		c.msgs = append(c.msgs, change{conventionalCommit: mc})
	} else {
		die(fmt.Errorf("failed to assert as conventionalcommits.ConventionalCommit: %v", m))
	}
}

func (c *chglog) addSc(s schemaChg) {
	c.msgs = append(c.msgs, change{schemaChange: &s})
}

func (c chglog) HasBreaking() bool {
	return len(c.Breaking()) > 0
}

func (c chglog) Breaking() []change {
	r := make([]change, 0)
	for _, m := range c.msgs {
		if m.IsBreakingChange() {
			r = append(r, m)
		}
	}

	return r
}

func (c chglog) HasNotes() bool {
	return len(c.Notes()) > 0

}

func (c chglog) Notes() []change {
	r := make([]change, 0)
	for _, m := range c.msgs {
		if m.IsNote() {
			r = append(r, m)
		}
	}
	return r
}

func (c chglog) HasFeatures() bool {
	return len(c.Features()) > 0
}

func (c chglog) Features() []change {
	r := make([]change, 0)
	for _, m := range c.msgs {
		if m.IsFeature() {
			r = append(r, m)
		}
	}
	return r
}

func (c chglog) HasEnhancements() bool {
	return len(c.Enhancements()) > 0
}

func (c chglog) Enhancements() []change {
	// Not implemented
	return make([]change, 0)
}

func (c chglog) HasFixes() bool {
	return len(c.Fixes()) > 0

}

func (c chglog) Fixes() []change {
	r := make([]change, 0)
	for _, m := range c.msgs {
		if m.IsFix() {
			r = append(r, m)
		}
	}
	return r
}

func main() {
	log.SetFlags(log.Lshortfile)

	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s [previous schema file] [new schema file]", os.Args[0])
	}

	fileFrom := os.Args[1]
	fileTo := os.Args[2]

	chgs := chglog{}

	previousVersion, gitChgs := gitChanges()
	chgs.previousVersion = *previousVersion
	for _, gitChg := range gitChgs {
		chgs.addCc(gitChg)
	}

	for _, schemaChg := range schemaChanges(fileFrom, fileTo) {
		chgs.addSc(schemaChg)
	}

	t, err := template.New("changelog-template").ParseFiles("CHANGELOG.md.gotmpl")
	die(err)

	f, err := os.Create("../../.build/CHANGELOG.md")
	die(err)
	defer f.Close()
	die(t.ExecuteTemplate(f, "changelog", chgs))
}

func gitChanges() (previousVersion *version.Version, commitsSinceLastVersion []conventionalcommits.Message) {
	repo, err := git.PlainOpen("../..")
	die(err)

	previousVersion, err = version.NewVersion("0.0.0")
	die(err)

	itr, err := repo.Tags()
	die(err)
	err = itr.ForEach(func(ref *plumbing.Reference) error {
		fmt.Println("Comparing versions: ", previousVersion.Original(), ref.Name().Short())
		v, err := version.NewVersion(ref.Name().Short())
		die(err)
		if v.GreaterThan(previousVersion) {
			previousVersion = v
		}
		return nil
	})
	die(err)

	fmt.Println("Previous version: ", previousVersion.Original())
	hash, err := repo.ResolveRevision(plumbing.Revision(previousVersion.Original()))
	die(err)
	releaseCommit, err := repo.CommitObject(*hash)
	die(err)

	commits, err := repo.Log(&git.LogOptions{
		Since: &releaseCommit.Author.When,
		Order: git.LogOrderCommitterTime,
	})
	die(err)

	chgs := make([]conventionalcommits.Message, 0)
	ccm := parser.NewMachine(parser.WithTypes(conventionalcommits.TypesConventional), parser.WithBestEffort())
	err = commits.ForEach(func(commit *object.Commit) error {
		cc, _ := ccm.Parse([]byte(commit.Message))

		chgs = append(chgs, cc)
		return nil
	})
	die(err)

	return previousVersion, chgs
}

func schemaChanges(fileFrom, fileTo string) []schemaChg {
	var changes = make([]schemaChg, 0)

	fmt.Println(filepath.Base(fileFrom), "->", filepath.Base(fileTo)+"\n")

	changes = append(changes, schemaJSONDeepDiff(fileFrom, fileTo)...)
	fmt.Println("total changes:", len(changes))

	return changes

}

func schemaJSONDeepDiff(fileA, fileB string) []schemaChg {
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

	return parseDiffs(rootDeltas)
}

func parseDiffs(d deepdiff.Deltas) []schemaChg {

	{ // useful for debugging
		f, err := os.Create("../../.build/schema-deepdiff.go")
		die(err)
		defer f.Close()
		f.WriteString(`package tmp
		import "github.com/qri-io/deepdiff"
		var _ = `)
		f.WriteString(AddLineBreaks(render.AsCode(d)))
	}

	var changes = make([]schemaChg, 0)

	if len(d) == 0 || len(d[0].Deltas) == 0 || len(d[0].Deltas[0].Deltas) == 0 {
		fmt.Println("No changes")
		return []schemaChg{}
	}

	d = d[0].Deltas[0].Deltas

	for _, dd := range d {

		// Recurse into provider schema, there should only be one provider in schema
		switch dd.Path.String() {
		case "provider":
			for _, c := range parseDiffsSchemaRepresentation(dd.Deltas[0].Deltas) {
				c.schemaType = schemaTypeProvider
				changes = append(changes, c)
			}
		case "resource_schemas":
			for _, r := range dd.Deltas {
				for _, c := range parseDiffsSchemaRepresentation(r.Deltas) {
					c.schemaType = schemaTypeResource
					c.field = append([]string{r.Path.String()}, c.field...)
					changes = append(changes, c)
				}
			}
		case "data_source_schemas":
			log.Fatalf("%+v", dd)
		default:
			log.Fatalf("%+v", dd)
		}
	}

	return changes
}

func parseDiffsSchemaRepresentation(d deepdiff.Deltas) []schemaChg {
	var changes = make([]schemaChg, 0)

	for _, dd := range d {
		if dd.Path.String() == "block" {
			changes = append(changes, parseDiffsBlock(dd.Deltas)...)
		}
	}
	for _, chg := range changes {
		chg.schemaType = schemaTypeProvider
	}
	return changes
}

func parseDiffsBlock(d deepdiff.Deltas) []schemaChg {
	var changes = make([]schemaChg, 0)

	for _, dd := range d {
		if dd.Path.String() == "attributes" {
			changes = append(changes, parseDiffsAttrs(dd.Deltas)...)
			for _, c := range changes {
				c.field = append([]string{dd.Path.String()}, c.field...)
			}
		}

		if dd.Path.String() == "block_types" {
			log.Fatalf("%+v", dd)
		}
	}

	// 	// "block_types" describes any nested blocks that appear directly
	// 	// inside the block.
	// 	// Keys in this map are the names of the block_type.
	// 	"block_types": {
	// 	  "example_block_name": {
	// 		// "nesting_mode" describes the nesting mode for the
	// 		// child block, and can be one of the following:
	// 		//    single
	// 		//    list
	// 		//    set
	// 		//    map
	// 	  "nesting_mode": "list",
	// 	  "block": <block-representation>,

	// 	  // "min_items" and "max_items" set lower and upper
	// 	  // limits on the number of child blocks allowed for
	// 	  // the list and set modes. These are
	// 	  // omitted for other modes.
	// 	  "min_items": 1,
	// 	  "max_items": 3
	// 	}
	//   }

	return changes

}

func parseDiffsAttrs(d deepdiff.Deltas) []schemaChg {
	var changes = make([]schemaChg, 0)

	for _, dd := range d {
		if dd.Type == deepdiff.DTInsert {
			changes = append(changes, schemaChg{
				changeType: schemaChgTypeAdd,
				field:      []string{dd.Path.String()},
				isMeta:     false,
				chgLvl:     schemaChgLvlFeat,
			})
			continue
		}
		if dd.Type == deepdiff.DTUpdate {
			log.Fatalf("%+v", dd)
		}
		if dd.Type == deepdiff.DTDelete {
			changes = append(changes, schemaChg{
				changeType: schemaChgTypeDel,
				field:      []string{dd.Path.String()},
				isMeta:     false,
				chgLvl:     schemaChgLvlBreak,
			})
			continue
		}

		if dd.Type == deepdiff.DTContext {
			for _, ddd := range dd.Deltas {

				if ddd.Path.String() == "nested_type" {
					if len(ddd.Deltas) != 1 || ddd.Deltas[0].Path.String() != "attributes" {
						log.Fatalf("%+v", ddd)
					}
					changes = append(changes, parseDiffsAttrs(ddd.Deltas[0].Deltas)...)
					continue
				}

				chg := schemaChg{
					field:  []string{ddd.Path.String()},
					isMeta: true,
				}

				switch ddd.Path.String() {
				case "type":
					chg.chgLvl = schemaChgLvlBreak
				case "description":
					chg.chgLvl = schemaChgLvlNote
				case "required":
					if ddd.Value == true {
						chg.chgLvl = schemaChgLvlBreak
					} else {
						chg.chgLvl = schemaChgLvlFix
					}
				case "optional":
					if ddd.Value == true {
						chg.chgLvl = schemaChgLvlFix
					} else {
						chg.chgLvl = schemaChgLvlBreak
					}
				case "computed":
					chg.chgLvl = schemaChgLvlNote
				case "sensitive":
					chg.chgLvl = schemaChgLvlNote
				default:
					log.Fatalf("%+v", ddd)
				}

				var chgT schemaChgType
				switch ddd.Type {
				case deepdiff.DTInsert:
					chgT = schemaChgTypeAdd
				case deepdiff.DTUpdate:
					chgT = schemaChgTypeChange
				case deepdiff.DTDelete:
					chgT = schemaChgTypeDel
				default:
					log.Fatalf("%+v", ddd)
				}
				chg.changeType = chgT
			}
		}

	}

	return changes
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
					//  milestone: 6
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
