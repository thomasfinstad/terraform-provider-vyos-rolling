// Code generated by /workspaces/terraform-provider-vyos/tools/build-vyos-infterface-definition-structs/main.go. DO NOT EDIT.

package vyosinterface

import (
	"encoding/xml"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/vyos/schema/interfacedefinition"
)

func system_proxy() interfacedefinition.InterfaceDefinition {
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
					NodeNameAttr: "proxy",
					OwnerAttr:    "${vyos_conf_scripts_dir}/system_proxy.py",
					Properties: []*interfacedefinition.Properties{{
						XMLName: xml.Name{
							Local: "properties",
						},
						Help: []string{"Sets a proxy for system wide use"},
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
							NodeNameAttr: "url",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Proxy URL"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Regex: []string{"http(s)?:\\/\\/[a-z0-9-\\.]+"},
								}},
							}},
						}, {
							IsBaseNode: false,
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "port",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Port number used by connection"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Validator: []*interfacedefinition.Validator{{
										XMLName: xml.Name{
											Local: "validator",
										},
										NameAttr:     "numeric",
										ArgumentAttr: "--range 1-65535",
									}},
								}},
								ValueHelp: []*interfacedefinition.ValueHelp{{
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "u32:1-65535",
									Description: "Numeric IP port",
								}},
								ConstraintErrorMessage: []string{"Port number must be in range 1 to 65535"},
							}},
						}, {
							IsBaseNode: false,
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "username",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Username used for authentication"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Regex: []string{"[[:ascii:]]{1,128}"},
								}},
								ValueHelp: []*interfacedefinition.ValueHelp{{
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "txt",
									Description: "Username",
								}},
								ConstraintErrorMessage: []string{"Username is limited to ASCII characters only, with a total length of 128"},
							}},
						}, {
							IsBaseNode: false,
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "password",
							Properties: []*interfacedefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Password used for authentication"},
								Constraint: []*interfacedefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Regex: []string{"[[:ascii:]]{1,128}"},
								}},
								ValueHelp: []*interfacedefinition.ValueHelp{{
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "txt",
									Description: "Password",
								}},
								ConstraintErrorMessage: []string{"Password is limited to ASCII characters only, with a total length of 128"},
							}},
						}},
					}},
				}},
			}},
		}},
	}
}
