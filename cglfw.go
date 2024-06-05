package main

import (
	"github.com/jurgen-kluft/ccode"
	pkg "github.com/jurgen-kluft/cglfw/package"
)

func main() {
	ccode.Init()
	ccode.GenerateSpecificFiles(ccode.CLANGFORMAT | ccode.GITIGNORE)
	ccode.Generate(pkg.GetPackage())
}
