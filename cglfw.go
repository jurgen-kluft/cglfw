package main

import (
	"github.com/jurgen-kluft/ccode"
	pkg "github.com/jurgen-kluft/cglfw/package"
)

func main() {
	if ccode.Init() {
		p := pkg.GetPackage()
		ccode.GenerateClangFormat()
		ccode.GenerateFiles(p)
		ccode.Generate(p)
	}
}
