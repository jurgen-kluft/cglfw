package main

import (
	"github.com/jurgen-kluft/ccode"
	pkg "github.com/jurgen-kluft/cglfw/package"
)

func main() {
	if ccode.Init() {
		ccode.GenerateClangFormat()
		ccode.GenerateGitIgnore()
		ccode.Generate(pkg.GetPackage())
	}
}
