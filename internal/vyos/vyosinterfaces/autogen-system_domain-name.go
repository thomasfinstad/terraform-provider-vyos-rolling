// Code generated by /repo/tools/build-vyos-infterface-definition-structs/main.go. DO NOT EDIT.

package vyosinterface

import (
	"encoding/xml"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/vyos/schemadefinition"
)

func system_domainname() schemadefinition.InterfaceDefinition {
	return schemadefinition.InterfaceDefinition{
		XMLName: xml.Name{
			Local: "interfaceDefinition",
		},
		Node: []*schemadefinition.Node{{
			IsBaseNode: false,
			XMLName: xml.Name{
				Local: "node",
			},
			NodeNameAttr: "system",
			Children: []*schemadefinition.Children{{
				XMLName: xml.Name{
					Local: "children",
				},
				LeafNode: []*schemadefinition.LeafNode{{
					IsBaseNode: false,
					XMLName: xml.Name{
						Local: "leafNode",
					},
					NodeNameAttr: "domain-name",
					OwnerAttr:    "${vyos_conf_scripts_dir}/system_host-name.py",
					Properties: []*schemadefinition.Properties{{
						XMLName: xml.Name{
							Local: "properties",
						},
						Help: []string{"System domain name"},
						Constraint: []*schemadefinition.Constraint{{
							XMLName: xml.Name{
								Local: "constraint",
							},
							Validator: []*schemadefinition.Validator{{
								XMLName: xml.Name{
									Local: "validator",
								},
								NameAttr: "fqdn",
							}},
						}},
						Priority: []string{"6"},
					}},
				}},
			}},
		}},
	}
}
