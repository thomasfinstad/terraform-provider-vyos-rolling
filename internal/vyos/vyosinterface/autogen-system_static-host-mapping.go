// Code generated by /workspaces/terraform-provider-vyos/tools/build-vyos-infterface-definition-structs/main.go. DO NOT EDIT.

package vyosinterface

import (
	"encoding/xml"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/vyos/schema/interfacedefinition"
)

func system_statichostmapping() interfacedefinition.InterfaceDefinition {
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
				Node: []*interfacedefinition.Node{{
					IsBaseNode: false,
					XMLName: xml.Name{
						Local: "node",
					},
					NodeNameAttr: "static-host-mapping",
					OwnerAttr:    "${vyos_conf_scripts_dir}/system_host-name.py",
					Properties: []*interfacedefinition.Properties{{
						XMLName: xml.Name{
							Local: "properties",
						},
						Help:     []string{"Map host names to addresses"},
						Priority: []string{"400"},
					}},
					Children: []*interfacedefinition.Children{{
						XMLName: xml.Name{
							Local: "children",
						},
						TagNode: []*interfacedefinition.TagNode{{
							IsBaseNode: false,
							XMLName: xml.Name{
								Local: "tagNode",
							},
							NodeNameAttr: "host-name",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Host name for static address mapping"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Regex: []string{"[A-Za-z0-9][-.A-Za-z0-9]&[A-Za-z0-9]"},
								}},
								ConstraintErrorMessage: []string{"Host-name must be alphanumeric and can contain hyphens"},
							}},
							Children: []*interfacedefinition.Children{{
								XMLName: xml.Name{
									Local: "children",
								},
								LeafNode: []*interfacedefinition.LeafNode{{
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "alias",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Alias for this address"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Regex: []string{".{1,63}"},
										}},
										ConstraintErrorMessage: []string{"invalid alias hostname, needs to be between 1 and 63 charactes"},
										Multi: []*interfacedefinition.Multi{{
											XMLName: xml.Name{
												Local: "multi",
											},
										}},
									}},
								}, {
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "inet",
									Properties: []*interfacedefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"IP Address"},
										Constraint: []*interfacedefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Validator: []*interfacedefinition.Validator{{
												XMLName: xml.Name{
													Local: "validator",
												},
												NameAttr: "ip-address",
											}},
										}},
										ValueHelp: []*interfacedefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "ipv4",
											Description: "IPv4 address",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "ipv6",
											Description: "IPv6 address",
										}},
										Multi: []*interfacedefinition.Multi{{
											XMLName: xml.Name{
												Local: "multi",
											},
										}},
									}},
								}},
							}},
						}},
					}},
				}},
			}},
		}},
	}
}