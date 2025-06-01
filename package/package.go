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
	mainpkg := denv.NewPackage("github.com\\jurgen-kluft", "cglfw")
	mainpkg.AddPackage(cbasepkg)

	// 'cglfw' library
	mainlib := denv.SetupCppLibProject(mainpkg, "cglfw")
	mainlib.AddDependencies(cbasepkg.GetMainLib()...)

	// 'cglfw' unittest project
	maintest := denv.SetupCppTestProject(mainpkg, "cglfw"+"test")
	maintest.AddDependencies(cunittestpkg.GetMainLib()...)
	maintest.AddDependencies(cbasepkg.GetMainLib()...)
	maintest.AddDependency(mainlib)

	if denv.IsWindows() {
		mainlib.AddDefine("_GLFW_WIN32;_GLFW_WGL;WIN32")
		maintest.AddDefine("_GLFW_WIN32;_GLFW_WGL;WIN32")
	} else if denv.IsMacOS() {
		mainlib.AddDefine("_GLFW_COCOA;MACOSX")
		maintest.AddDefine("_GLFW_COCOA;MACOSX")
	} else if denv.IsLinux() {
		mainlib.AddDefine("_GLFW_X11;_GLFW_GFX;LINUX")
		maintest.AddDefine("_GLFW_X11;_GLFW_GFX;LINUX")
	}

	mainpkg.AddMainLib(mainlib)
	mainpkg.AddUnittest(maintest)
	return mainpkg
}
