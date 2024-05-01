// Code generated by /repo/tools/build-vyos-infterface-definition-structs/main.go. DO NOT EDIT.

package vyosinterface

import (
	"encoding/xml"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/vyos/schemadefinition"
)

func system_statichostmapping() schemadefinition.InterfaceDefinition {
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
				Node: []*schemadefinition.Node{{
					IsBaseNode: false,
					XMLName: xml.Name{
						Local: "node",
					},
					NodeNameAttr: "static-host-mapping",
					OwnerAttr:    "${vyos_conf_scripts_dir}/system_host-name.py",
					Properties: []*schemadefinition.Properties{{
						XMLName: xml.Name{
							Local: "properties",
						},
						Help:     []string{"Map host names to addresses"},
						Priority: []string{"400"},
					}},
					Children: []*schemadefinition.Children{{
						XMLName: xml.Name{
							Local: "children",
						},
						TagNode: []*schemadefinition.TagNode{{
							IsBaseNode: false,
							XMLName: xml.Name{
								Local: "tagNode",
							},
							NodeNameAttr: "host-name",
							Properties: []*schemadefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Host name for static address mapping"},
								Constraint: []*schemadefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Regex: []string{"[A-Za-z0-9][-.A-Za-z0-9]&[A-Za-z0-9]"},
								}},
								ConstraintErrorMessage: []string{"Host-name must be alphanumeric and can contain hyphens"},
							}},
							Children: []*schemadefinition.Children{{
								XMLName: xml.Name{
									Local: "children",
								},
								LeafNode: []*schemadefinition.LeafNode{{
									IsBaseNode: false,
									XMLName: xml.Name{
										Local: "leafNode",
									},
									NodeNameAttr: "alias",
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Alias for this address"},
										Constraint: []*schemadefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Regex: []string{".{1,63}"},
										}},
										ConstraintErrorMessage: []string{"invalid alias hostname, needs to be between 1 and 63 charactes"},
										Multi: []*schemadefinition.Multi{{
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
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"IP Address"},
										Constraint: []*schemadefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Validator: []*schemadefinition.Validator{{
												XMLName: xml.Name{
													Local: "validator",
												},
												NameAttr: "ip-address",
											}},
										}},
										ValueHelp: []*schemadefinition.ValueHelp{{
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
										Multi: []*schemadefinition.Multi{{
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
