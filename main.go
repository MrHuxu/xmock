package main

import (
	"flag"
)

var (
	flagFile      string
	flagDirectory string
	flagTarget    string
	flagSrcPkg    string
	flagOutPkg    string
)

func init() {
	flag.StringVar(&flagFile, "file", "", "generate the mock structs for a file")
	flag.StringVar(&flagDirectory, "dir", ".", "generate the mock structs for whole directory")
	flag.StringVar(&flagTarget, "target", "file", "output to file or stdout")
	flag.StringVar(&flagSrcPkg, "srcpkg", "", "full package path of the file and directory")
	flag.StringVar(&flagOutPkg, "outpkg", "mock", "package name of mock structs")

	initLogger()
}

func main() {
	flag.Parse()
	logger.Infof("_args||file=%s||directory=%s||target=%s||srcpkg=%s||outpkg=%s", flagFile, flagDirectory, flagTarget, flagSrcPkg, flagOutPkg)

	generate()
}
