package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/gdexlab/go-render/render"
	"github.com/go-git/go-git/v5"
	"github.com/hashicorp/go-version"
	"github.com/leodido/go-conventionalcommits"
)

const (
	rollingrelease bool = true
	DEBUG               = true
)

//  Ref: https://developer.hashicorp.com/terraform/plugin/best-practices/versioning#versioning-specification
//  Ref: https://developer.hashicorp.com/terraform/plugin/best-practices/versioning#changelog-specification
//  Ref: https://developer.hashicorp.com/terraform/cli/v1.5.x/commands/providers/schema

func main() {
	log.SetFlags(log.Lshortfile)

	if len(os.Args) != 3 && !DEBUG {
		log.Fatalf("Usage: %s [previous schema file] [new schema file]", os.Args[0])
	} else if DEBUG {
		os.Args = append(os.Args,
			getModuleRoot()+"data/provider-schema.json",
			getModuleRoot()+".build/provider-schema.json",
		)
	}

	changeLog := ChangeLog{}

	previousVersion, gitChgs := GenerateGitChanges()
	changeLog.previousVersion = *previousVersion
	for _, gitChg := range gitChgs {
		changeLog.addCc(gitChg)
	}

	var oldProvidersSchema ProviderSchemas
	var newProvidersSchema ProviderSchemas

	fileOld := os.Args[1]
	previousSchemaByte, err := os.ReadFile(fileOld)
	die(err)
	err = json.Unmarshal(previousSchemaByte, &oldProvidersSchema)
	die(err)

	fileNew := os.Args[2]
	newSchemaByte, err := os.ReadFile(fileNew)
	die(err)
	err = json.Unmarshal(newSchemaByte, &newProvidersSchema)
	die(err)

	if oldProvidersSchema.Count() != 1 {
		die(fmt.Errorf("old provider schema contains too many providers, expected exactly 1 got: %d", oldProvidersSchema.Count()))
	}
	if newProvidersSchema.Count() != 1 {
		die(fmt.Errorf("new provider schema contains too many providers, expected exactly 1 got: %d", newProvidersSchema.Count()))
	}

	providerName := oldProvidersSchema.ProviderNames()[0]

	oldProvider := oldProvidersSchema.Schemas[providerName]
	newProvider := newProvidersSchema.Schemas[providerName]

	providerDiff := ProviderSchemaChanges(oldProvider, newProvider)

	if DEBUG { // Dump the diff struct
		f, err := os.Create("../../.build/changelog-diff.go")
		die(err)
		defer f.Close()
		f.WriteString(`package tmp
		import "github.com/thomasfinstad/terraform-provider-vyos-rolling/tools/generate-changelog"
		var _ = `)
		_, err = f.WriteString(AddLineBreaks(render.AsCode(providerDiff)))
		die(err)
	}

	if DEBUG { // Dump the diff description
		f, err := os.Create("../../.build/changelog-description.md")
		die(err)
		defer f.Close()
		f.WriteString("# CHG log\n")

		var desc []string
		for _, d := range providerDiff {
			desc = append(desc, d.Description([]string{}...)...)
		}
		descs := strings.Join(desc, "\n")
		_, err = f.WriteString(descs)
		die(err)
	}

	changeLog.addSc(providerDiff)

	t, err := template.New("changelog-template").ParseFiles("CHANGELOG.md.gotmpl")
	die(err)

	f, err := os.Create("../../.build/CHANGELOG.md")
	die(err)
	defer f.Close()
	die(t.ExecuteTemplate(f, "changelog", changeLog))
}

// change contains either a git change or a schema change
type change struct {
	conventionalCommit *conventionalcommits.ConventionalCommit
	schemaChange       *SchemaChange
}

func (c change) IsBreakingChange() bool {
	if c.IsSchemaChange() {
		return c.schemaChange.EffectiveChangeSeverity() == SchemaChangeSeverityBreaking
	}
	return c.conventionalCommit.IsBreakingChange()
}

func (c change) IsNote() bool {
	if c.IsSchemaChange() {
		return c.schemaChange.EffectiveChangeSeverity() == SchemaChangeSeverityNote
	}
	return (!c.conventionalCommit.Ok() ||
		c.conventionalCommit.Type == "docs" ||
		c.conventionalCommit.Type == "refactor" ||
		c.conventionalCommit.Type == "ci")
}

func (c change) IsFeature() bool {
	if c.IsSchemaChange() {
		return c.schemaChange.EffectiveChangeSeverity() == SchemaChangeSeverityFeature
	}
	return c.conventionalCommit.IsFeat()
}

func (c change) IsEnhancement() bool {
	if c.IsSchemaChange() {
		return c.schemaChange.EffectiveChangeSeverity() == SchemaChangeSeverityEnhancement
	}
	return false
}

func (c change) IsFix() bool {
	if c.IsSchemaChange() {
		return c.schemaChange.EffectiveChangeSeverity() == SchemaChangeSeverityFix
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
	if c.IsSchemaChange() {
		return strings.Join(c.schemaChange.Description([]string{}...), "\n")
	}
	return c.conventionalCommit.Description
}

func (c change) IsSchemaChange() bool {
	return c.conventionalCommit == nil
}

type ChangeLog struct {
	msgs            []change
	previousVersion version.Version
}

func (c ChangeLog) Version() *version.Version {
	var err error

	major := c.previousVersion.Segments()[0]
	minor := c.previousVersion.Segments()[1]
	patch := c.previousVersion.Segments()[2]

	repo, err := git.PlainOpen("../..")
	die(err)

	isMajor := false
	isMinor := false

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
	} else {
		patch = patch + 1
	}

	if rollingrelease {
		b, err := os.ReadFile("../../data/vyos-1x-info.txt")
		die(err)
		patch, err = strconv.Atoi(strings.ReplaceAll(strings.Split(string(b), "T")[0], "-", ""))
		die(err)
	}

	var newVersion *version.Version
	for i := 0; i < 10; i++ {
		patch_nr := patch*10 + i
		newVersion, err = version.NewVersion(fmt.Sprintf("%d.%d.%d", major, minor, patch_nr))
		die(err)
		_, err := repo.Tag("v" + newVersion.Original())
		if err != nil {
			break
		}
	}

	return newVersion
}

func (c ChangeLog) Date() string {
	return time.Now().UTC().Format("2006-01-02 15-04-05 MST")
}

func (c *ChangeLog) addCc(m conventionalcommits.Message) {
	if mc, ok := m.(*conventionalcommits.ConventionalCommit); ok {
		c.msgs = append(c.msgs, change{conventionalCommit: mc})
	} else {
		die(fmt.Errorf("failed to assert as conventionalcommits.ConventionalCommit: %v", m))
	}
}

func (c *ChangeLog) addSc(s SchemaChanges) {
	for _, chg := range s {
		c.msgs = append(c.msgs, change{schemaChange: &chg})
	}
}

func (c ChangeLog) HasGitChanges() bool {
	for _, m := range c.msgs {
		if !m.IsSchemaChange() {
			return true
		}
	}

	return false
}

func (c ChangeLog) GitChanges() ChangeLog {
	r := make([]change, 0)
	for _, m := range c.msgs {
		if !m.IsSchemaChange() {
			r = append(r, m)
		}
	}

	return ChangeLog{
		previousVersion: c.previousVersion,
		msgs:            r,
	}
}

func (c ChangeLog) HasSchemaChanges() bool {
	for _, m := range c.msgs {
		if m.IsSchemaChange() {
			return true
		}
	}

	return false
}

func (c ChangeLog) SchemaChanges() SchemaChanges {
	r := SchemaChanges{}
	for _, m := range c.msgs {
		if m.IsSchemaChange() {
			r = append(r, *m.schemaChange)
		}
	}

	return r
}

func (c ChangeLog) HasBreaking() bool {
	return len(c.Breaking()) > 0
}

func (c ChangeLog) Breaking() []change {
	r := make([]change, 0)
	for _, m := range c.msgs {
		if m.IsBreakingChange() {
			r = append(r, m)
		}
	}

	return r
}

func (c ChangeLog) HasNotes() bool {
	return len(c.Notes()) > 0

}

func (c ChangeLog) Notes() []change {
	r := make([]change, 0)
	for _, m := range c.msgs {
		if m.IsNote() {
			r = append(r, m)
		}
	}
	return r
}

func (c ChangeLog) HasFeatures() bool {
	return len(c.Features()) > 0
}

func (c ChangeLog) Features() []change {
	r := make([]change, 0)
	for _, m := range c.msgs {
		if m.IsFeature() {
			r = append(r, m)
		}
	}
	return r
}

func (c ChangeLog) HasEnhancements() bool {
	return len(c.Enhancements()) > 0
}

func (c ChangeLog) Enhancements() []change {
	r := make([]change, 0)
	for _, m := range c.msgs {
		if m.IsEnhancement() {
			r = append(r, m)
		}
	}
	return r
}

func (c ChangeLog) HasFixes() bool {
	return len(c.Fixes()) > 0

}

func (c ChangeLog) Fixes() []change {
	r := make([]change, 0)
	for _, m := range c.msgs {
		if m.IsFix() {
			r = append(r, m)
		}
	}
	return r
}
