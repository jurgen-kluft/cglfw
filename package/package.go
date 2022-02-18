package libglfw

import (
	xbase "github.com/jurgen-kluft/xbase/package"
	"github.com/jurgen-kluft/xcode/denv"
	xentry "github.com/jurgen-kluft/xentry/package"
	xunittest "github.com/jurgen-kluft/xunittest/package"
)

// GetPackage returns the package object of 'libglfw'
func GetPackage() *denv.Package {
	// Dependencies
	xunittestpkg := xunittest.GetPackage()
	xentrypkg := xentry.GetPackage()
	xbasepkg := xbase.GetPackage()

	// The main (libglfw) package
	mainpkg := denv.NewPackage("libglfw")
	mainpkg.AddPackage(xunittestpkg)
	mainpkg.AddPackage(xentrypkg)
	mainpkg.AddPackage(xbasepkg)

	// 'libglfw' library
	mainlib := denv.SetupDefaultCppLibProject("libglfw", "github.com\\jurgen-kluft\\libglfw")
	mainlib.Dependencies = append(mainlib.Dependencies, xbasepkg.GetMainLib())

	// 'libglfw' unittest project
	maintest := denv.SetupDefaultCppTestProject("libglfw_test", "github.com\\jurgen-kluft\\libglfw")
	maintest.Dependencies = append(maintest.Dependencies, xunittestpkg.GetMainLib())
	maintest.Dependencies = append(maintest.Dependencies, xentrypkg.GetMainLib())
	maintest.Dependencies = append(maintest.Dependencies, xbasepkg.GetMainLib())
	maintest.Dependencies = append(maintest.Dependencies, mainlib)

	if denv.XCodeOS == "windows" {
		mainlib.AddDefine("_GLFW_WIN32")
		maintest.AddDefine("_GLFW_WIN32")
	} else if denv.XCodeOS == "darwin" {
		mainlib.AddDefine("_GLFW_COCOA")
		maintest.AddDefine("_GLFW_COCOA")
	} else if denv.XCodeOS == "linux" {
		mainlib.AddDefine("_GLFW_X11")
		maintest.AddDefine("_GLFW_X11")
	}

	mainpkg.AddMainLib(mainlib)
	mainpkg.AddUnittest(maintest)
	return mainpkg
}
