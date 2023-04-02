// Code generated by /workspaces/terraform-provider-vyos/tools/build-resource-terraform-schemas/main.go. DO NOT EDIT.

package vyosinterface

import (
	"encoding/xml"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/provider/vyos/schema/interfacedefinition"
)

func Dhcpv6Relay() interfacedefinition.InterfaceDefinition {
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
								OwnerAttr: `${vyos_conf_scripts_dir}/dhcpv6_relay.py`,
								Properties: []*interfacedefinition.Properties{
									{
										XMLName: xml.Name{
											Local: `properties`},
										Help: []string{
											`DHCPv6 Relay Agent parameters`},
										Priority: []string{
											`900`},
										KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}},
								Children: []*interfacedefinition.Children{
									{
										XMLName: xml.Name{
											Local: `children`},
										TagNode: []*interfacedefinition.TagNode{
											{
												XMLName: xml.Name{
													Local: `tagNode`},
												Properties: []*interfacedefinition.Properties{
													{
														XMLName: xml.Name{
															Local: `properties`},
														Help: []string{
															`Interface for DHCPv6 Relay Agent to listen for requests`},
														CompletionHelp: []*interfacedefinition.CompletionHelp{
															{
																XMLName: xml.Name{
																	Local: `completionHelp`},
																Script: []string{
																	`${vyos_completion_dir}/list_interfaces`}}},
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
																			`IPv6 address on listen-interface listen for requests on`},
																		Constraint: []*interfacedefinition.Constraint{
																			{
																				XMLName: xml.Name{
																					Local: `constraint`},
																				Validator: []*interfacedefinition.Validator{
																					{
																						XMLName: xml.Name{
																							Local: `validator`},
																						NameAttr:     `ipv6-address`,
																						ArgumentAttr: ``}}}},
																		ValueHelp: []*interfacedefinition.ValueHelp{
																			{
																				XMLName: xml.Name{
																					Local: `valueHelp`},
																				Format:      `ipv6`,
																				Description: `IPv6 address on listen interface`}},
																		KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}}}}}}},
											{
												XMLName: xml.Name{
													Local: `tagNode`},
												Properties: []*interfacedefinition.Properties{
													{
														XMLName: xml.Name{
															Local: `properties`},
														Help: []string{
															`Interface for DHCPv6 Relay Agent forward requests out`},
														CompletionHelp: []*interfacedefinition.CompletionHelp{
															{
																XMLName: xml.Name{
																	Local: `completionHelp`},
																Script: []string{
																	`${vyos_completion_dir}/list_interfaces`}}},
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
																			`IPv6 address to forward requests to`},
																		Constraint: []*interfacedefinition.Constraint{
																			{
																				XMLName: xml.Name{
																					Local: `constraint`},
																				Validator: []*interfacedefinition.Validator{
																					{
																						XMLName: xml.Name{
																							Local: `validator`},
																						NameAttr:     `ipv6-address`,
																						ArgumentAttr: ``}}}},
																		ValueHelp: []*interfacedefinition.ValueHelp{
																			{
																				XMLName: xml.Name{
																					Local: `valueHelp`},
																				Format:      `ipv6`,
																				Description: `IPv6 address of the DHCP server`}},
																		Multi: []*interfacedefinition.Multi{
																			{
																				XMLName: xml.Name{
																					Local: `multi`}}},
																		KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}}}}}}}},
										LeafNode: []*interfacedefinition.LeafNode{
											{
												XMLName: xml.Name{
													Local: `leafNode`},
												DefaultValue: []string{
													`10`},
												Properties: []*interfacedefinition.Properties{
													{
														XMLName: xml.Name{
															Local: `properties`},
														Help: []string{
															`Maximum hop count for which requests will be processed`},
														Constraint: []*interfacedefinition.Constraint{
															{
																XMLName: xml.Name{
																	Local: `constraint`},
																Validator: []*interfacedefinition.Validator{
																	{
																		XMLName: xml.Name{
																			Local: `validator`},
																		NameAttr:     `numeric`,
																		ArgumentAttr: `--range 1-255`}}}},
														ValueHelp: []*interfacedefinition.ValueHelp{
															{
																XMLName: xml.Name{
																	Local: `valueHelp`},
																Format:      `u32:1-255`,
																Description: `Hop count`}},
														ConstraintErrorMessage: []string{
															`max-hop-count must be a value between 1 and 255`},
														KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}}},
											{
												XMLName: xml.Name{
													Local: `leafNode`},
												Properties: []*interfacedefinition.Properties{
													{
														XMLName: xml.Name{
															Local: `properties`},
														Help: []string{
															`Option to set DHCPv6 interface-ID option`},
														Valueless: []*interfacedefinition.Valueless{
															{
																XMLName: xml.Name{
																	Local: `valueless`}}},
														KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}}}}}}}},
						LeafNode: []*interfacedefinition.LeafNode(nil)}}}}}
}
