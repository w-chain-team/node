package main

import (
	_ "embed"

	"github.com/w-chain-team/node/command/root"
	"github.com/w-chain-team/node/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
