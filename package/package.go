package cglfw

import (
	cbase "github.com/jurgen-kluft/cbase/package"
	"github.com/jurgen-kluft/ccode/denv"
	cunittest "github.com/jurgen-kluft/cunittest/package"
)

// GetPackage returns the package object of 'cglfw'
func GetPackage() *denv.Package {
	// Dependencies
	cunittestpkg := cunittest.GetPackage()
	cbasepkg := cbase.GetPackage()

	// The main (cglfw) package
	mainpkg := denv.NewPackage("cglfw")
	mainpkg.AddPackage(cbasepkg)

	// 'cglfw' library
	mainlib := denv.SetupDefaultCppLibProject("cglfw", "github.com\\jurgen-kluft\\cglfw")
	mainlib.Dependencies = append(mainlib.Dependencies, cbasepkg.GetMainLib())

	// 'cglfw' unittest project
	maintest := denv.SetupDefaultCppTestProject("cglfw"+"test", "github.com\\jurgen-kluft\\cglfw")
	maintest.Dependencies = append(maintest.Dependencies, cunittestpkg.GetMainLib())
	maintest.Dependencies = append(maintest.Dependencies, cbasepkg.GetMainLib())
	maintest.Dependencies = append(maintest.Dependencies, mainlib)

	if denv.OS == "windows" {
		mainlib.AddDefine("_GLFW_WIN32;_GLFW_WGL;WIN32")
		maintest.AddDefine("_GLFW_WIN32;_GLFW_WGL;WIN32")
	} else if denv.OS == "darwin" {
		mainlib.AddDefine("_GLFW_COCOA;MACOSX")
		maintest.AddDefine("_GLFW_COCOA;MACOSX")
	} else if denv.OS == "linux" {
		mainlib.AddDefine("_GLFW_X11;_GLFW_GFX;LINUX")
		maintest.AddDefine("_GLFW_X11;_GLFW_GFX;LINUX")
	}

	mainpkg.AddMainLib(mainlib)
	mainpkg.AddUnittest(maintest)
	return mainpkg
}
