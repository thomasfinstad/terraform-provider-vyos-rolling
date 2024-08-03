package main

import (
	"fmt"
	"log"
	"reflect"
	"slices"
	"strings"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/zclconf/go-cty/cty"
	"golang.org/x/exp/maps"
)

type SchemaChangeTo string

const (
	SchemaChangeToProvider   SchemaChangeTo = "PROVIDER"
	SchemaChangeToResource   SchemaChangeTo = "RESOURCE"
	SchemaChangeToDataSource SchemaChangeTo = "DATASOURCE"
	SchemaChangeToAttribute  SchemaChangeTo = "ATTRIBUTE"
)

type SchemaChangeOperation string

const (
	SchemaChangeOperationAdd SchemaChangeOperation = "Added"
	SchemaChangeOperationDel SchemaChangeOperation = "Removed"
	SchemaChangeOperationChg SchemaChangeOperation = "Changed"
	SchemaChangeOperationSub SchemaChangeOperation = "<sub>"
)

type SchemaChangeSeverityInt int

const (
	SchemaChangeSeverityUNKNOWN SchemaChangeSeverityInt = iota
	SchemaChangeSeverityFix
	SchemaChangeSeverityEnhancement
	SchemaChangeSeverityNote
	SchemaChangeSeverityFeature
	SchemaChangeSeverityBreaking
)

type SchemaChangeSeverityStr string

var SchemaChangeSeverity = map[SchemaChangeSeverityStr]SchemaChangeSeverityInt{
	"UNKNOWN":     SchemaChangeSeverityUNKNOWN,
	"FIX":         SchemaChangeSeverityFix,
	"ENHANCEMENT": SchemaChangeSeverityEnhancement,
	"NOTE":        SchemaChangeSeverityNote,
	"FEATURE":     SchemaChangeSeverityFeature,
	"BREAKING":    SchemaChangeSeverityBreaking,
}

type SchemaChange struct {
	SubChanges     SchemaChanges
	ChangeTo       SchemaChangeTo
	Operation      SchemaChangeOperation
	Address        string
	ChangeSeverity SchemaChangeSeverityInt
	CustomMsg      string // Optional custom message
}

func (c SchemaChange) EffectiveChangeSeverity() SchemaChangeSeverityInt {
	t := c.ChangeSeverity
	for _, s := range c.SubChanges {
		st := s.EffectiveChangeSeverity()
		if st > t {
			t = st
		}
	}
	return t
}

func (c SchemaChange) Description(parentAddress ...string) (descriptionLines []string) {

	parentAddress = slices.DeleteFunc(parentAddress, func(e string) bool { return e == "" })

	switch c.ChangeTo {
	case SchemaChangeToProvider:
		for _, s := range c.SubChanges {
			descriptionLines = append(descriptionLines, fmt.Sprintf("%s", s.Description(parentAddress...)))
		}
		return descriptionLines
	case SchemaChangeToResource:
		switch c.Operation {
		case SchemaChangeOperationAdd:
			descriptionLines = append(descriptionLines, fmt.Sprintf("* New Resource `%s`", c.Address))
			return descriptionLines
		case SchemaChangeOperationDel:
			descriptionLines = append(descriptionLines, fmt.Sprintf("* **Removed Resource** `%s`", c.Address))
			return descriptionLines
		case SchemaChangeOperationSub:
			// desc := c.SubChanges[0].Description(parentAddress)
			// if c.SubChanges.Count() == 1 { // Converge to 1 line if only 1 subchange
			// 	descriptionLines = append(descriptionLines, fmt.Sprintf("* Modified Resource `%s` %s", c.Address, strings.Join(, " ")))
			// 	return descriptionLines
			// }

			descriptionLines = append(descriptionLines, fmt.Sprintf("* Modified Resource `%s`", c.Address))
			for _, s := range c.SubChanges {
				for _, l := range s.Description(parentAddress...) {
					descriptionLines = append(descriptionLines, fmt.Sprintf("\t%s", l))
				}
			}
			return descriptionLines
		default:
			log.Fatalf("Unknown resource ChangeOperation type: %#v", c)
		}
	// case ChangeToDataSource:
	// if c.SubChanges.Count() == 0 {
	// 	descriptionLines = fmt.Sprintf("* Data Source `%s` has been %s", c.Address, c.Operation)
	// } else if c.SubChanges.Count() == 1 {
	// 	descriptionLines = fmt.Sprintf("* Data Source `%s` has changed %s", c.Address, c.SubChanges[0].Description())
	// } else {
	// 	descriptionLines = fmt.Sprintf("* Data Source `%s` has the following changes:\n", c.Address)
	// 	for _, s := range c.SubChanges {
	// 		descriptionLines += fmt.Sprintf("\t%s\n", s.Description())
	// 	}
	// }
	case SchemaChangeToAttribute:
		switch c.Operation {

		case SchemaChangeOperationAdd:
			descriptionLines = []string{fmt.Sprintf("* New attribute `%s`", strings.Join(append(parentAddress, c.Address), "."))}
			return descriptionLines
		case SchemaChangeOperationDel:
			descriptionLines = []string{fmt.Sprintf("* **Removed attribute** `%s`", strings.Join(append(parentAddress, c.Address), "."))}
			return descriptionLines
		case SchemaChangeOperationChg:
			descriptionLines = []string{fmt.Sprintf("* Modified attribute `%s` %s", strings.Join(append(parentAddress, c.Address), "."), c.CustomMsg)}
			return descriptionLines
		case SchemaChangeOperationSub:
			if c.SubChanges.Count() == 1 {
				descriptionLines = append(descriptionLines, c.SubChanges[0].Description(append(parentAddress, c.Address)...)...)
			} else {
				descriptionLines = []string{fmt.Sprintf("* Modified attributes under: `%s` %s", strings.Join(append(parentAddress, c.Address), "."), c.CustomMsg)}
				for _, s := range c.SubChanges {
					for _, l := range s.Description([]string{}...) {
						descriptionLines = append(descriptionLines, fmt.Sprintf("\t%s", l))
					}
				}
			}
			return descriptionLines
		}
	default:
		log.Fatalf("Unknown ChangeTo type: %#v", c)
	}
	return nil
}

type SchemaChanges []SchemaChange

func (c SchemaChanges) Count() int {
	return len(c)
}

func (c SchemaChanges) TotalCount() int {
	i := c.Count()
	for _, s := range c {
		i += s.SubChanges.TotalCount()
	}
	return i
}

func (c *SchemaChanges) Add(chg ...SchemaChange) {
	*c = append(*c, chg...)
}

func (c SchemaChanges) HasChangeTo(s SchemaChangeTo) bool {
	for _, chg := range c {
		if chg.ChangeTo == s {
			return true
		}
	}
	return false
}

func (c SchemaChanges) GetChangeTo(s SchemaChangeTo) SchemaChanges {
	var chgs SchemaChanges
	for _, chg := range c {
		if chg.ChangeTo == s {
			chgs = append(chgs, chg)
		}
	}
	return chgs
}

func (c SchemaChanges) HasChangeSeverity(s SchemaChangeSeverityStr) bool {
	for _, chg := range c {
		if chg.EffectiveChangeSeverity() == SchemaChangeSeverity[SchemaChangeSeverityStr(strings.ToUpper(string(s)))] {
			return true
		}
	}
	return false
}

func (c SchemaChanges) GetChangeSeverity(s SchemaChangeSeverityStr) SchemaChanges {
	var chgs SchemaChanges
	for _, chg := range c {
		if chg.EffectiveChangeSeverity() == SchemaChangeSeverity[SchemaChangeSeverityStr(strings.ToUpper(string(s)))] {
			chgs = append(chgs, chg)
		}
	}
	return chgs
}

type ProviderSchemas tfjson.ProviderSchemas

func (p ProviderSchemas) Count() int {
	return len(p.Schemas)
}

func (p ProviderSchemas) ProviderNames() []string {
	return maps.Keys[map[string]*tfjson.ProviderSchema](p.Schemas)
}

func ProviderSchemaChanges(old, new *tfjson.ProviderSchema) (changes SchemaChanges) {
	//-----
	// Provider config
	changes.Add(GenerateSchemaChanges(old.ConfigSchema, new.ConfigSchema)...)

	//-----
	// Resources
	for k, v := range old.ResourceSchemas {
		// Deleted resources
		if _, ok := new.ResourceSchemas[k]; !ok {
			changes.Add(SchemaChange{
				ChangeTo:       SchemaChangeToResource,
				Operation:      SchemaChangeOperationDel,
				ChangeSeverity: SchemaChangeSeverityBreaking,
				Address:        k,
			})
			continue
		}

		// Changed resource
		if chg := GenerateSchemaChanges(v, new.ResourceSchemas[k]); chg.TotalCount() > 0 {
			changes.Add(SchemaChange{
				ChangeTo:       SchemaChangeToResource,
				Operation:      SchemaChangeOperationSub,
				ChangeSeverity: SchemaChangeSeverityNote,
				Address:        k,
				SubChanges:     chg,
			})
		}
	}

	// Added resource
	for k := range new.ResourceSchemas {
		if _, ok := old.ResourceSchemas[k]; !ok {
			changes.Add(SchemaChange{
				ChangeTo:       SchemaChangeToResource,
				Operation:      SchemaChangeOperationAdd,
				ChangeSeverity: SchemaChangeSeverityFeature,
				Address:        k,
			})
		}
	}

	//-----
	// Data sources
	for k, v := range old.DataSourceSchemas {
		// Deleted data source
		if _, ok := new.DataSourceSchemas[k]; !ok {
			changes.Add(SchemaChange{
				ChangeTo:       SchemaChangeToDataSource,
				Operation:      SchemaChangeOperationDel,
				ChangeSeverity: SchemaChangeSeverityBreaking,
				Address:        k,
			})
			continue
		}

		// Changed data source
		if chg := GenerateSchemaChanges(v, new.DataSourceSchemas[k]); chg.TotalCount() > 0 {
			changes.Add(SchemaChange{
				ChangeTo:       SchemaChangeToDataSource,
				Operation:      SchemaChangeOperationSub,
				ChangeSeverity: SchemaChangeSeverityNote,
				Address:        k,
				SubChanges:     chg,
			})
		}
	}

	// Added data source
	for k := range new.DataSourceSchemas {
		if _, ok := old.DataSourceSchemas[k]; !ok {
			changes.Add(SchemaChange{
				ChangeTo:       SchemaChangeToDataSource,
				Operation:      SchemaChangeOperationAdd,
				ChangeSeverity: SchemaChangeSeverityFeature,
				Address:        k,
			})
		}
	}

	return changes
}

func GenerateSchemaChanges(old, new *tfjson.Schema) (changes SchemaChanges) {
	changes.Add(SchemaBlockChanges(old.Block, new.Block)...)

	return changes
}

func SchemaBlockChanges(old, new *tfjson.SchemaBlock) (changes SchemaChanges) {
	//--------
	// Attributes
	for k, v := range old.Attributes {
		// Deleted attrs
		if _, ok := new.Attributes[k]; !ok {
			changes.Add(SchemaChange{
				ChangeTo:       SchemaChangeToAttribute,
				Operation:      SchemaChangeOperationDel,
				ChangeSeverity: SchemaChangeSeverityBreaking,
				Address:        k,
			})
			continue
		}

		// Changed attrs
		if attrChanges := SchemaAttributeChanges(v, new.Attributes[k]); attrChanges.TotalCount() > 0 {
			changes.Add(SchemaChange{
				ChangeTo:       SchemaChangeToAttribute,
				Operation:      SchemaChangeOperationSub,
				ChangeSeverity: SchemaChangeSeverityNote,
				Address:        k,
				SubChanges:     attrChanges,
			})
		}
	}

	// Added attrs
	for k := range new.Attributes {
		if _, ok := old.Attributes[k]; !ok {
			changes.Add(SchemaChange{
				ChangeTo:       SchemaChangeToAttribute,
				Operation:      SchemaChangeOperationAdd,
				ChangeSeverity: SchemaChangeSeverityFeature,
				Address:        k,
			})
		}
	}

	//--------
	// Block types
	if len(old.NestedBlocks) > 0 || len(new.NestedBlocks) > 0 {
		log.Fatal("Changelog generation does not support block_types yet")
	}

	return changes
}

func SchemaAttributeChanges(old, new *tfjson.SchemaAttribute) (changes SchemaChanges) {
	if old.AttributeType != new.AttributeType {
		if new.AttributeType != (cty.Type{}) {
			changes.Add(SchemaChange{
				ChangeTo:       SchemaChangeToAttribute,
				Operation:      SchemaChangeOperationAdd,
				ChangeSeverity: SchemaChangeSeverityEnhancement,
				CustomMsg:      "type changed to " + new.AttributeType.GoString(),
			})
		} else {
			changes.Add(SchemaChange{
				ChangeTo:       SchemaChangeToAttribute,
				Operation:      SchemaChangeOperationDel,
				ChangeSeverity: SchemaChangeSeverityBreaking,
				CustomMsg:      "type information removed",
			})
		}
	}

	// AttributeNestedType
	if !reflect.DeepEqual(old.AttributeNestedType, new.AttributeNestedType) {
		//--------
		// Attributes
		for k, v := range old.AttributeNestedType.Attributes {
			// Deleted attrs
			if _, ok := new.AttributeNestedType.Attributes[k]; !ok {
				changes.Add(SchemaChange{
					ChangeTo:       SchemaChangeToAttribute,
					Operation:      SchemaChangeOperationDel,
					ChangeSeverity: SchemaChangeSeverityBreaking,
					Address:        k,
				})
				continue
			}

			// Changed attrs
			attrChanges := SchemaAttributeChanges(v, new.AttributeNestedType.Attributes[k])
			if attrChanges.TotalCount() > 0 {
				changes.Add(SchemaChange{
					ChangeTo:       SchemaChangeToAttribute,
					Operation:      SchemaChangeOperationSub,
					ChangeSeverity: SchemaChangeSeverityNote,
					Address:        k,
					SubChanges:     attrChanges,
				})
			}
		}

		// Added attrs
		for k := range new.AttributeNestedType.Attributes {
			if _, ok := old.AttributeNestedType.Attributes[k]; !ok {
				changes.Add(SchemaChange{
					ChangeTo:       SchemaChangeToAttribute,
					Operation:      SchemaChangeOperationAdd,
					ChangeSeverity: SchemaChangeSeverityFeature,
					Address:        k,
				})
			}
		}
		//--------
		// Nesting mode
		if old.AttributeNestedType.NestingMode != new.AttributeNestedType.NestingMode {
			changes.Add(SchemaChange{
				ChangeTo:       SchemaChangeToAttribute,
				Operation:      SchemaChangeOperationChg,
				ChangeSeverity: SchemaChangeSeverityBreaking,
				CustomMsg:      "nested attribute mode changed to " + string(new.AttributeNestedType.NestingMode),
			})
		}

		//--------
		// Min items
		if old.AttributeNestedType.MinItems != new.AttributeNestedType.MinItems {
			changes.Add(SchemaChange{
				ChangeTo:       SchemaChangeToAttribute,
				Operation:      SchemaChangeOperationAdd,
				ChangeSeverity: SchemaChangeSeverityEnhancement,
				CustomMsg:      "minimum count changed to",
			})
		}

		//--------
		// Max items
		if old.AttributeNestedType.MaxItems != new.AttributeNestedType.MaxItems {
			changes.Add(SchemaChange{
				ChangeTo:       SchemaChangeToAttribute,
				Operation:      SchemaChangeOperationAdd,
				ChangeSeverity: SchemaChangeSeverityEnhancement,
				CustomMsg:      "maximum count changed to",
			})
		}
	}

	if old.Description != new.Description {
		var op SchemaChangeOperation
		if old.Description == "" {
			op = SchemaChangeOperationAdd
		} else if new.Description == "" {
			op = SchemaChangeOperationDel
		} else {
			op = SchemaChangeOperationChg
		}
		changes.Add(SchemaChange{
			ChangeTo:       SchemaChangeToAttribute,
			Operation:      op,
			ChangeSeverity: SchemaChangeSeverityEnhancement,
			CustomMsg:      "changed description",
		})
	}

	if old.DescriptionKind != new.DescriptionKind {
		changes.Add(SchemaChange{
			ChangeTo:       SchemaChangeToAttribute,
			Operation:      SchemaChangeOperationChg,
			ChangeSeverity: SchemaChangeSeverityEnhancement,
			CustomMsg:      "description kind changed to " + string(new.DescriptionKind),
		})
	}

	if old.Deprecated != new.Deprecated && new.Deprecated {
		changes.Add(SchemaChange{
			ChangeTo:       SchemaChangeToAttribute,
			Operation:      SchemaChangeOperationChg,
			ChangeSeverity: SchemaChangeSeverityNote,
			CustomMsg:      "marked as deprecated",
		})
	}

	if old.Required != new.Required {
		if new.Required {
			changes.Add(SchemaChange{
				ChangeTo:       SchemaChangeToAttribute,
				Operation:      SchemaChangeOperationAdd,
				ChangeSeverity: SchemaChangeSeverityBreaking,
				CustomMsg:      "is new required",
			})

		} else {
			changes.Add(SchemaChange{
				ChangeTo:       SchemaChangeToAttribute,
				Operation:      SchemaChangeOperationDel,
				ChangeSeverity: SchemaChangeSeverityFix,
				CustomMsg:      "is no longer required",
			})
		}
	}

	if old.Optional != new.Optional {
		if old.Optional {
			changes.Add(SchemaChange{
				ChangeTo:       SchemaChangeToAttribute,
				Operation:      SchemaChangeOperationDel,
				ChangeSeverity: SchemaChangeSeverityBreaking,
				CustomMsg:      "is no longer optional",
			})
		} else {
			changes.Add(SchemaChange{
				ChangeTo:       SchemaChangeToAttribute,
				Operation:      SchemaChangeOperationDel,
				ChangeSeverity: SchemaChangeSeverityEnhancement,
				CustomMsg:      "is now optional",
			})
		}
	}

	if old.Computed != new.Computed {
		changes.Add(SchemaChange{
			ChangeTo:       SchemaChangeToAttribute,
			Operation:      SchemaChangeOperationChg,
			ChangeSeverity: SchemaChangeSeverityEnhancement,
			Address:        "changed computed status of attribute",
		})
	}

	if old.Sensitive != new.Sensitive {
		changes.Add(SchemaChange{
			ChangeTo:       SchemaChangeToAttribute,
			Operation:      SchemaChangeOperationChg,
			ChangeSeverity: SchemaChangeSeverityEnhancement,
			CustomMsg:      "changed sensitive status of attribute",
		})
	}

	return changes
}
