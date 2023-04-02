// Code generated by /workspaces/terraform-provider-vyos/tools/build-resource-terraform-schemas/main.go. DO NOT EDIT.

package vyosinterface

import (
	"encoding/xml"

	"github.com/thomasfinstad/terraform-provider-vyos/internal/provider/vyos/schema/interfacedefinition"
)

func Lldp() interfacedefinition.InterfaceDefinition {
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
								OwnerAttr: `${vyos_conf_scripts_dir}/lldp.py`,
								Properties: []*interfacedefinition.Properties{
									{
										XMLName: xml.Name{
											Local: `properties`},
										Help: []string{
											`LLDP settings`},
										Priority: []string{
											`985`},
										KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}},
								Children: []*interfacedefinition.Children{
									{
										XMLName: xml.Name{
											Local: `children`},
										Node: []*interfacedefinition.Node{
											{
												XMLName: xml.Name{
													Local: `node`},
												Properties: []*interfacedefinition.Properties{
													{
														XMLName: xml.Name{
															Local: `properties`},
														Help: []string{
															`Legacy (vendor specific) protocols`},
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
																			`Listen for CDP for Cisco routers/switches`},
																		Valueless: []*interfacedefinition.Valueless{
																			{
																				XMLName: xml.Name{
																					Local: `valueless`}}},
																		KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}}},
															{
																XMLName: xml.Name{
																	Local: `leafNode`},
																Properties: []*interfacedefinition.Properties{
																	{
																		XMLName: xml.Name{
																			Local: `properties`},
																		Help: []string{
																			`Listen for EDP for Extreme routers/switches`},
																		Valueless: []*interfacedefinition.Valueless{
																			{
																				XMLName: xml.Name{
																					Local: `valueless`}}},
																		KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}}},
															{
																XMLName: xml.Name{
																	Local: `leafNode`},
																Properties: []*interfacedefinition.Properties{
																	{
																		XMLName: xml.Name{
																			Local: `properties`},
																		Help: []string{
																			`Listen for FDP for Foundry routers/switches`},
																		Valueless: []*interfacedefinition.Valueless{
																			{
																				XMLName: xml.Name{
																					Local: `valueless`}}},
																		KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}}},
															{
																XMLName: xml.Name{
																	Local: `leafNode`},
																Properties: []*interfacedefinition.Properties{
																	{
																		XMLName: xml.Name{
																			Local: `properties`},
																		Help: []string{
																			`Listen for SONMP for Nortel routers/switches`},
																		Valueless: []*interfacedefinition.Valueless{
																			{
																				XMLName: xml.Name{
																					Local: `valueless`}}},
																		KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}}}}}}},
											{
												XMLName: xml.Name{
													Local: `node`},
												Properties: []*interfacedefinition.Properties{
													{
														XMLName: xml.Name{
															Local: `properties`},
														Help: []string{
															`SNMP parameters for LLDP`},
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
																			`Enable SNMP queries of the LLDP database`},
																		Valueless: []*interfacedefinition.Valueless{
																			{
																				XMLName: xml.Name{
																					Local: `valueless`}}},
																		KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}}}}}}}},
										TagNode: []*interfacedefinition.TagNode{
											{
												XMLName: xml.Name{
													Local: `tagNode`},
												Properties: []*interfacedefinition.Properties{
													{
														XMLName: xml.Name{
															Local: `properties`},
														Help: []string{
															`Location data for interface`},
														ValueHelp: []*interfacedefinition.ValueHelp{
															{
																XMLName: xml.Name{
																	Local: `valueHelp`},
																Format:      `all`,
																Description: `Location data all interfaces`},
															{
																XMLName: xml.Name{
																	Local: `valueHelp`},
																Format:      `txt`,
																Description: `Location data for a specific interface`}},
														CompletionHelp: []*interfacedefinition.CompletionHelp{
															{
																XMLName: xml.Name{
																	Local: `completionHelp`},
																List: []string{
																	`all`},
																Script: []string{
																	`${vyatta_sbindir}/vyatta-interfaces.pl --show all`}}},
														KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}},
												Children: []*interfacedefinition.Children{
													{
														XMLName: xml.Name{
															Local: `children`},
														Node: []*interfacedefinition.Node{
															{
																XMLName: xml.Name{
																	Local: `node`},
																Properties: []*interfacedefinition.Properties{
																	{
																		XMLName: xml.Name{
																			Local: `properties`},
																		Help: []string{
																			`LLDP-MED location data`},
																		KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}},
																Children: []*interfacedefinition.Children{
																	{
																		XMLName: xml.Name{
																			Local: `children`},
																		Node: []*interfacedefinition.Node{
																			{
																				XMLName: xml.Name{
																					Local: `node`},
																				Properties: []*interfacedefinition.Properties{
																					{
																						XMLName: xml.Name{
																							Local: `properties`},
																						Help: []string{
																							`Coordinate based location`},
																						KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}},
																				Children: []*interfacedefinition.Children{
																					{
																						XMLName: xml.Name{
																							Local: `children`},
																						LeafNode: []*interfacedefinition.LeafNode{
																							{
																								XMLName: xml.Name{
																									Local: `leafNode`},
																								DefaultValue: []string{
																									`0`},
																								Properties: []*interfacedefinition.Properties{
																									{
																										XMLName: xml.Name{
																											Local: `properties`},
																										Help: []string{
																											`Altitude in meters`},
																										Constraint: []*interfacedefinition.Constraint{
																											{
																												XMLName: xml.Name{
																													Local: `constraint`},
																												Validator: []*interfacedefinition.Validator{
																													{
																														XMLName: xml.Name{
																															Local: `validator`},
																														NameAttr:     `numeric`,
																														ArgumentAttr: ``}}}},
																										ValueHelp: []*interfacedefinition.ValueHelp{
																											{
																												XMLName: xml.Name{
																													Local: `valueHelp`},
																												Format:      `0`,
																												Description: `No altitude`},
																											{
																												XMLName: xml.Name{
																													Local: `valueHelp`},
																												Format:      `[+-]<meters>`,
																												Description: `Altitude in meters`}},
																										ConstraintErrorMessage: []string{
																											`Altitude should be a positive or negative number`},
																										KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}}},
																							{
																								XMLName: xml.Name{
																									Local: `leafNode`},
																								DefaultValue: []string{
																									`WGS84`},
																								Properties: []*interfacedefinition.Properties{
																									{
																										XMLName: xml.Name{
																											Local: `properties`},
																										Help: []string{
																											`Coordinate datum type`},
																										Constraint: []*interfacedefinition.Constraint{
																											{
																												XMLName: xml.Name{
																													Local: `constraint`},
																												Regex: []string{
																													`(WGS84|NAD83|MLLW)`},
																												Validator: []*interfacedefinition.Validator(nil)}},
																										ValueHelp: []*interfacedefinition.ValueHelp{
																											{
																												XMLName: xml.Name{
																													Local: `valueHelp`},
																												Format:      `WGS84`,
																												Description: `WGS84`},
																											{
																												XMLName: xml.Name{
																													Local: `valueHelp`},
																												Format:      `NAD83`,
																												Description: `NAD83`},
																											{
																												XMLName: xml.Name{
																													Local: `valueHelp`},
																												Format:      `MLLW`,
																												Description: `NAD83/MLLW`}},
																										ConstraintErrorMessage: []string{
																											`Datum should be WGS84,
 NAD83,
 or MLLW`},
																										CompletionHelp: []*interfacedefinition.CompletionHelp{
																											{
																												XMLName: xml.Name{
																													Local: `completionHelp`},
																												List: []string{
																													`WGS84 NAD83 MLLW`},
																												Script: []string(nil)}},
																										KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}}},
																							{
																								XMLName: xml.Name{
																									Local: `leafNode`},
																								Properties: []*interfacedefinition.Properties{
																									{
																										XMLName: xml.Name{
																											Local: `properties`},
																										Help: []string{
																											`Latitude`},
																										Constraint: []*interfacedefinition.Constraint{
																											{
																												XMLName: xml.Name{
																													Local: `constraint`},
																												Regex: []string{
																													`(\\d+)(\\.\\d+)?[nNsS]`},
																												Validator: []*interfacedefinition.Validator(nil)}},
																										ValueHelp: []*interfacedefinition.ValueHelp{
																											{
																												XMLName: xml.Name{
																													Local: `valueHelp`},
																												Format:      `<latitude>`,
																												Description: `Latitude (example \'37.524449N\')`}},
																										ConstraintErrorMessage: []string{
																											`Latitude should be a number followed by S or N`},
																										KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}}},
																							{
																								XMLName: xml.Name{
																									Local: `leafNode`},
																								Properties: []*interfacedefinition.Properties{
																									{
																										XMLName: xml.Name{
																											Local: `properties`},
																										Help: []string{
																											`Longitude`},
																										Constraint: []*interfacedefinition.Constraint{
																											{
																												XMLName: xml.Name{
																													Local: `constraint`},
																												Regex: []string{
																													`(\\d+)(\\.\\d+)?[eEwW]`},
																												Validator: []*interfacedefinition.Validator(nil)}},
																										ValueHelp: []*interfacedefinition.ValueHelp{
																											{
																												XMLName: xml.Name{
																													Local: `valueHelp`},
																												Format:      `<longitude>`,
																												Description: `Longitude (example \'122.267255W\')`}},
																										ConstraintErrorMessage: []string{
																											`Longiture should be a number followed by E or W`},
																										KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}}}}}}}},
																		LeafNode: []*interfacedefinition.LeafNode{
																			{
																				XMLName: xml.Name{
																					Local: `leafNode`},
																				Properties: []*interfacedefinition.Properties{
																					{
																						XMLName: xml.Name{
																							Local: `properties`},
																						Help: []string{
																							`ECS ELIN (Emergency location identifier number)`},
																						Constraint: []*interfacedefinition.Constraint{
																							{
																								XMLName: xml.Name{
																									Local: `constraint`},
																								Regex: []string{
																									`[0-9]{10,
25}`},
																								Validator: []*interfacedefinition.Validator(nil)}},
																						ValueHelp: []*interfacedefinition.ValueHelp{
																							{
																								XMLName: xml.Name{
																									Local: `valueHelp`},
																								Format:      `u32:0-9999999999`,
																								Description: `Emergency Call Service ELIN number (between 10-25 numbers)`}},
																						ConstraintErrorMessage: []string{
																							`ELIN number must be between 10-25 numbers`},
																						KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}}}}}}}},
														LeafNode: []*interfacedefinition.LeafNode{
															{
																XMLName: xml.Name{
																	Local: `leafNode`},
																Properties: []*interfacedefinition.Properties{
																	{
																		XMLName: xml.Name{
																			Local: `properties`},
																		Help: []string{
																			`Disable instance`},
																		Valueless: []*interfacedefinition.Valueless{
																			{
																				XMLName: xml.Name{
																					Local: `valueless`}}},
																		KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}}}}}}}},
										LeafNode: []*interfacedefinition.LeafNode{
											{
												XMLName: xml.Name{
													Local: `leafNode`},
												Properties: []*interfacedefinition.Properties{
													{
														XMLName: xml.Name{
															Local: `properties`},
														Help: []string{
															`Management IP Address`},
														Constraint: []*interfacedefinition.Constraint{
															{
																XMLName: xml.Name{
																	Local: `constraint`},
																Validator: []*interfacedefinition.Validator{
																	{
																		XMLName: xml.Name{
																			Local: `validator`},
																		NameAttr:     `ip-address`,
																		ArgumentAttr: ``}}}},
														ValueHelp: []*interfacedefinition.ValueHelp{
															{
																XMLName: xml.Name{
																	Local: `valueHelp`},
																Format:      `ipv4`,
																Description: `IPv4 Management Address`},
															{
																XMLName: xml.Name{
																	Local: `valueHelp`},
																Format:      `ipv6`,
																Description: `IPv6 Management Address`}},
														CompletionHelp: []*interfacedefinition.CompletionHelp{
															{
																XMLName: xml.Name{
																	Local: `completionHelp`},
																Script: []string{
																	`${vyos_completion_dir}/list_local_ips.sh --both`}}},
														Multi: []*interfacedefinition.Multi{
															{
																XMLName: xml.Name{
																	Local: `multi`}}},
														KeepChildOrder: []*interfacedefinition.KeepChildOrder(nil)}}}}}}}},
						LeafNode: []*interfacedefinition.LeafNode(nil)}}}}}
}
