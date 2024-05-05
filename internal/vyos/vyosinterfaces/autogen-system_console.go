// Code generated by /repo/tools/build-vyos-infterface-definition-structs/main.go. DO NOT EDIT.

package vyosinterface

import (
	"encoding/xml"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/vyos/schemadefinition"
)

func system_console() schemadefinition.InterfaceDefinition {
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
					NodeNameAttr: "console",
					OwnerAttr:    "${vyos_conf_scripts_dir}/system_console.py",
					Properties: []*schemadefinition.Properties{{
						XMLName: xml.Name{
							Local: "properties",
						},
						Help:     []string{"Serial console configuration"},
						Priority: []string{"100"},
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
							NodeNameAttr: "device",
							Properties: []*schemadefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Serial console device name"},
								Constraint: []*schemadefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Regex: []string{"(ttyS[0-9]+|hvc[0-9]+|usb[0-9]+b.&)"},
								}},
								ValueHelp: []*schemadefinition.ValueHelp{{
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "ttySN",
									Description: "TTY device name, regular serial port",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "usbNbXpY",
									Description: "TTY device name, USB based",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "hvcN",
									Description: "Xen console",
								}},
								CompletionHelp: []*schemadefinition.CompletionHelp{{
									XMLName: xml.Name{
										Local: "completionHelp",
									},
									Script: []string{"ls -1 /dev | grep -e ttyS -e hvc", "if [ -d /dev/serial/by-bus ]; then ls -1 /dev/serial/by-bus; fi"},
								}},
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
									NodeNameAttr: "speed",
									DefaultValue: []string{"115200"},
									Properties: []*schemadefinition.Properties{{
										XMLName: xml.Name{
											Local: "properties",
										},
										Help: []string{"Console baud rate"},
										Constraint: []*schemadefinition.Constraint{{
											XMLName: xml.Name{
												Local: "constraint",
											},
											Regex: []string{"(1200|2400|4800|9600|19200|38400|57600|115200)"},
										}},
										ValueHelp: []*schemadefinition.ValueHelp{{
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "1200",
											Description: "1200 bps",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "2400",
											Description: "2400 bps",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "4800",
											Description: "4800 bps",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "9600",
											Description: "9600 bps",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "19200",
											Description: "19200 bps",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "38400",
											Description: "38400 bps",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "57600",
											Description: "57600 bps",
										}, {
											XMLName: xml.Name{
												Local: "valueHelp",
											},
											Format:      "115200",
											Description: "115200 bps",
										}},
										CompletionHelp: []*schemadefinition.CompletionHelp{{
											XMLName: xml.Name{
												Local: "completionHelp",
											},
											List: []string{"1200 2400 4800 9600 19200 38400 57600 115200"},
										}},
									}},
								}},
							}},
						}},
						LeafNode: []*schemadefinition.LeafNode{{
							IsBaseNode: false,
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "powersave",
							Properties: []*schemadefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Enable screen blank powersaving on VGA console"},
								Valueless: []*schemadefinition.Valueless{{
									XMLName: xml.Name{
										Local: "valueless",
									},
								}},
							}},
						}},
					}},
				}},
			}},
		}},
	}
}