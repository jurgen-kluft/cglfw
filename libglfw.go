package main

import (
	pkg "github.com/jurgen-kluft/libglfw/package"
	"github.com/jurgen-kluft/xcode"
)

func main() {
	xcode.Init()
	xcode.Generate(pkg.GetPackage())
}
