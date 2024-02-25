// Code generated by /workspaces/terraform-provider-vyos/tools/build-vyos-infterface-definition-structs/main.go. DO NOT EDIT.

package vyosinterface

import (
	"encoding/xml"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/vyos/schema/interfacedefinition"
)

func system_timezone() interfacedefinition.InterfaceDefinition {
	return interfacedefinition.InterfaceDefinition{
		XMLName: xml.Name{
			Local: "interfaceDefinition",
		},
		Node: []*interfacedefinition.Node{{
			IsBaseNode: false,
			XMLName: xml.Name{
				Local: "node",
			},
			NodeNameAttr: "system",
			Children: []*interfacedefinition.Children{{
				XMLName: xml.Name{
					Local: "children",
				},
				LeafNode: []*interfacedefinition.LeafNode{{
					IsBaseNode: false,
					XMLName: xml.Name{
						Local: "leafNode",
					},
					NodeNameAttr: "time-zone",
					OwnerAttr:    "${vyos_conf_scripts_dir}/system_timezone.py",
					Properties: []*interfacedefinition.Properties{{
						XMLName: xml.Name{
							Local: "properties",
						},
						Help: []string{"Local time zone (default UTC)"},
						Constraint: []*interfacedefinition.Constraint{{
							XMLName: xml.Name{
								Local: "constraint",
							},
							Validator: []*interfacedefinition.Validator{{
								XMLName: xml.Name{
									Local: "validator",
								},
								NameAttr:     "timezone",
								ArgumentAttr: "--validate",
							}},
						}},
						CompletionHelp: []*interfacedefinition.CompletionHelp{{
							XMLName: xml.Name{
								Local: "completionHelp",
							},
							Script: []string{"timedatectl list-timezones"},
						}},
						Priority: []string{"100"},
					}},
				}},
			}},
		}},
	}
}
