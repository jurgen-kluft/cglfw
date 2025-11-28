package cglfw

import (
	"github.com/jurgen-kluft/ccode/denv"
	ccore "github.com/jurgen-kluft/ccore/package"
	cunittest "github.com/jurgen-kluft/cunittest/package"
)

const (
	repo_path = "github.com\\jurgen-kluft"
	repo_name = "cglfw"
)

// GetPackage returns the package object of 'cglfw'
func GetPackage() *denv.Package {
	// Dependencies
	cunittestpkg := cunittest.GetPackage()
	ccorepkg := ccore.GetPackage()

	// The main (cglfw) package
	mainpkg := denv.NewPackage(repo_path, repo_name)
	mainpkg.AddPackage(ccorepkg)

	// 'cglfw' library
	mainlib := denv.SetupCppLibProject(mainpkg, repo_name)
	mainlib.AddDependencies(ccorepkg.GetMainLib())

	// 'cglfw' library for testing
	testlib := denv.SetupCppTestLibProject(mainpkg, repo_name)
	testlib.AddDependencies(ccorepkg.GetTestLib())

	// 'cglfw' unittest project
	maintest := denv.SetupCppTestProject(mainpkg, repo_name+"_test")
	maintest.AddDependencies(cunittestpkg.GetMainLib())
	maintest.AddDependency(testlib)

	if denv.IsWindows() {
		mainlib.AddDefine("_GLFW_WIN32", "_GLFW_WGL", "WIN32")
		maintest.AddDefine("_GLFW_WIN32", "_GLFW_WGL", "WIN32")
	} else if denv.IsMacOS() {
		mainlib.AddDefine("_GLFW_COCOA", "MACOSX")
		maintest.AddDefine("_GLFW_COCOA", "MACOSX")
	} else if denv.IsLinux() {
		mainlib.AddDefine("_GLFW_X11", "_GLFW_GFX", "LINUX")
		maintest.AddDefine("_GLFW_X11", "_GLFW_GFX", "LINUX")
	}

	mainpkg.AddMainLib(mainlib)
	mainpkg.AddTestLib(testlib)
	mainpkg.AddUnittest(maintest)
	return mainpkg
}
