module github.com/thomasfinstad/terraform-provider-vyos-rolling/tools/generate-vyos-infterface-definition-structs

go 1.23.1

replace github.com/thomasfinstad/terraform-provider-vyos-rolling => ../../

require (
	github.com/dave/dst v0.27.3
	github.com/gdexlab/go-render v1.0.1
	github.com/thomasfinstad/terraform-provider-vyos-rolling v1.6.202408210
	golang.org/x/text v0.19.0
)

require (
	github.com/fatih/color v1.13.0 // indirect
	github.com/hashicorp/go-hclog v1.5.0 // indirect
	github.com/hashicorp/terraform-plugin-log v0.9.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	golang.org/x/mod v0.17.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/tools v0.21.1-0.20240508182429-e35e4ccd0d2d // indirect
)
