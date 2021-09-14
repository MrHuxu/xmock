package main

import (
	"flag"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func init() {
	flag.StringVar(&buildOption.File, "file", "", "generate the mock structs for a file")
	flag.StringVar(&buildOption.Directory, "dir", ".", "generate the mock structs for whole directory")
	flag.StringVar(&buildOption.OutPkg, "outpkg", "mock", "package name of mock structs")

	initLogger()
}

func main() {
	flag.Parse()

	logger.Infof("_args||file=%s||directory=%s||outpkg=%s", buildOption.File, buildOption.Directory, buildOption.OutPkg)

	var err error
	switch {
	case buildOption.File != "":
		arr := strings.Split(buildOption.File, "/")
		buildOption.Directory = strings.Join(arr[:len(arr)-1], "/")
		generateForFile(&buildOption)

	default:
		if buildOption.Directory == "." {
			if buildOption.Directory, err = os.Getwd(); err != nil {
				logger.Fatalf("_fatal||reason=%+v", err)
			}
		}

		generateForDirectory(&buildOption)
	}
}

func generateForDirectory(buildOption *BuildOption) {
	fileInfos, err := ioutil.ReadDir(buildOption.Directory)
	if err != nil {
		logger.Fatalf("_fatal||reason=%+v", err)
	}

	for _, fileInfo := range fileInfos {
		if !strings.HasSuffix(fileInfo.Name(), ".go") {
			continue
		}

		generateForFile(&BuildOption{
			File:      buildOption.Directory + "/" + fileInfo.Name(),
			Directory: buildOption.Directory,
			OutPkg:    buildOption.OutPkg,
		})
	}
}

func generateForFile(buildOption *BuildOption) {
	arr := strings.Split(buildOption.File, "/")
	dir := strings.Join(arr[:len(arr)-1], "/")
	filename := arr[len(arr)-1]

	fileItem := parseSrcFile(buildOption.File)
	if len(fileItem.InterfaceItems) == 0 {
		return
	}
	reader := buildDistFile(fileItem, buildOption)

	outDir := dir + "/" + buildOption.OutPkg
	os.MkdirAll(outDir, 0755)

	distFile, err := os.Create(outDir + "/" + filename)
	if err != nil {
		logger.Fatalf("_fatal||reason=%+v", err)
	}
	defer distFile.Close()

	if _, err := io.Copy(distFile, reader); err != nil {
		logger.Fatalf("_fatal||reason=%+v", err)
	}

	logger.Infof(
		"_generate_done||src=%s||dist=%s",
		strings.Replace(buildOption.File, buildOption.Directory, "", -1)[1:],
		strings.Replace(outDir+"/"+filename, buildOption.Directory, "", -1)[1:],
	)
}
