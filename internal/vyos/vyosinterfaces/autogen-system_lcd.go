// Code generated by tools/build-vyos-infterface-definition-structs/main.go. DO NOT EDIT.

package vyosinterface

import (
	"encoding/xml"

	"github.com/thomasfinstad/terraform-provider-vyos-rolling/internal/vyos/schemadefinition"
)

func system_lcd() schemadefinition.InterfaceDefinition {
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
					NodeNameAttr: "lcd",
					OwnerAttr:    "${vyos_conf_scripts_dir}/system_lcd.py",
					Properties: []*schemadefinition.Properties{{
						XMLName: xml.Name{
							Local: "properties",
						},
						Help:     []string{"System LCD display"},
						Priority: []string{"100"},
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
							NodeNameAttr: "model",
							Properties: []*schemadefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Model of the display attached to this system"},
								Constraint: []*schemadefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Regex: []string{"(cfa-533|cfa-631|cfa-633|cfa-635|hd44780|sdec)"},
								}},
								ValueHelp: []*schemadefinition.ValueHelp{{
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "cfa-533",
									Description: "Crystalfontz CFA-533",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "cfa-631",
									Description: "Crystalfontz CFA-631",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "cfa-633",
									Description: "Crystalfontz CFA-633",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "cfa-635",
									Description: "Crystalfontz CFA-635",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "hd44780",
									Description: "Hitachi HD44780, Caswell Appliances",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "sdec",
									Description: "Lanner, Watchguard, Nexcom NSA, Sophos UTM appliances",
								}},
								CompletionHelp: []*schemadefinition.CompletionHelp{{
									XMLName: xml.Name{
										Local: "completionHelp",
									},
									List: []string{"cfa-533 cfa-631 cfa-633 cfa-635 hd44780 sdec"},
								}},
							}},
						}, {
							IsBaseNode: false,
							XMLName: xml.Name{
								Local: "leafNode",
							},
							NodeNameAttr: "device",
							Properties: []*schemadefinition.Properties{{
								XMLName: xml.Name{
									Local: "properties",
								},
								Help: []string{"Physical device used by LCD display"},
								Constraint: []*schemadefinition.Constraint{{
									XMLName: xml.Name{
										Local: "constraint",
									},
									Regex: []string{"(ttyS[0-9]+|usb[0-9]+b.&)"},
								}},
								ValueHelp: []*schemadefinition.ValueHelp{{
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "ttySXX",
									Description: "TTY device name, regular serial port",
								}, {
									XMLName: xml.Name{
										Local: "valueHelp",
									},
									Format:      "usbNbXpY",
									Description: "TTY device name, USB based",
								}},
								CompletionHelp: []*schemadefinition.CompletionHelp{{
									XMLName: xml.Name{
										Local: "completionHelp",
									},
									Script: []string{"ls -1 /dev | grep ttyS", "if [ -d /dev/serial/by-bus ]; then ls -1 /dev/serial/by-bus; fi"},
								}},
							}},
						}},
					}},
				}},
			}},
		}},
	}
}
