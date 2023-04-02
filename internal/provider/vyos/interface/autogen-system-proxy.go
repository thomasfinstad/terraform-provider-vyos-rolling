// Code generated by /workspaces/terraform-provider-vyos/tools/build-resource-terraform-schemas/main.go. DO NOT EDIT.

package vyosinterface

import (
	"encoding/xml"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/provider/vyos/schema/interfacedefinition"
)

func SystemProxy() interfacedefinition.InterfaceDefinition {
	return interfacedefinition.InterfaceDefinition{
		XMLName: xml.Name{
			Local: `interfaceDefinition`},
		Node: []*interfacedefinition.Node{
			{
				XMLName: xml.Name{
					Local: `node`},
				Children: []*interfacedefinition.Children{
					{
						XMLName: xml.Name{
							Local: `children`},
						Node: []*interfacedefinition.Node{
							{
								XMLName: xml.Name{
									Local: `node`},
								OwnerAttr: `${vyos_conf_scripts_dir}/system-proxy.py`,
								Properties: []*interfacedefinition.Properties{
									{
										XMLName: xml.Name{
											Local: `properties`},
										Help: []string{
											`Sets a proxy for system wide use`},
										KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}},
								Children: []*interfacedefinition.Children{
									{
										XMLName: xml.Name{
											Local: `children`},
										LeafNode: []*interfacedefinition.LeafNode{
											{
												XMLName: xml.Name{
													Local: `leafNode`},
												Properties: []*interfacedefinition.Properties{
													{
														XMLName: xml.Name{
															Local: `properties`},
														Help: []string{
															`Proxy URL`},
														Constraint: []*interfacedefinition.Constraint{
															{
																XMLName: xml.Name{
																	Local: `constraint`},
																Regex: []string{
																	`http(s)?:\\/\\/[a-z0-9-\\.]+`},
																Validator: []*interfacedefinition.Validator(nil)}},
														KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}}},
											{
												XMLName: xml.Name{
													Local: `leafNode`},
												Properties: []*interfacedefinition.Properties{
													{
														XMLName: xml.Name{
															Local: `properties`},
														Help: []string{
															`Port number used by connection`},
														Constraint: []*interfacedefinition.Constraint{
															{
																XMLName: xml.Name{
																	Local: `constraint`},
																Validator: []*interfacedefinition.Validator{
																	{
																		XMLName: xml.Name{
																			Local: `validator`},
																		NameAttr:     `numeric`,
																		ArgumentAttr: `--range 1-65535`}}}},
														ValueHelp: []*interfacedefinition.ValueHelp{
															{
																XMLName: xml.Name{
																	Local: `valueHelp`},
																Format:      `u32:1-65535`,
																Description: `Numeric IP port`}},
														ConstraintErrorMessage: []string{
															`Port number must be in range 1 to 65535`},
														KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}}},
											{
												XMLName: xml.Name{
													Local: `leafNode`},
												Properties: []*interfacedefinition.Properties{
													{
														XMLName: xml.Name{
															Local: `properties`},
														Help: []string{
															`Username used for authentication`},
														Constraint: []*interfacedefinition.Constraint{
															{
																XMLName: xml.Name{
																	Local: `constraint`},
																Regex: []string{
																	`[[:ascii:]]{1,
128}`},
																Validator: []*interfacedefinition.Validator(nil)}},
														ValueHelp: []*interfacedefinition.ValueHelp{
															{
																XMLName: xml.Name{
																	Local: `valueHelp`},
																Format:      `txt`,
																Description: `Username`}},
														ConstraintErrorMessage: []string{
															`Username is limited to ASCII characters only,
 with a total length of 128`},
														KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}}},
											{
												XMLName: xml.Name{
													Local: `leafNode`},
												Properties: []*interfacedefinition.Properties{
													{
														XMLName: xml.Name{
															Local: `properties`},
														Help: []string{
															`Password used for authentication`},
														Constraint: []*interfacedefinition.Constraint{
															{
																XMLName: xml.Name{
																	Local: `constraint`},
																Regex: []string{
																	`[[:ascii:]]{1,
128}`},
																Validator: []*interfacedefinition.Validator(nil)}},
														ValueHelp: []*interfacedefinition.ValueHelp{
															{
																XMLName: xml.Name{
																	Local: `valueHelp`},
																Format:      `txt`,
																Description: `Password`}},
														ConstraintErrorMessage: []string{
															`Password is limited to ASCII characters only,
 with a total length of 128`},
														KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}}}}}}}},
						LeafNode: []*interfacedefinition.LeafNode(nil)}}}}}
}
