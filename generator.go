package main

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func generate() {
	var err error
	switch {
	case flagFile != "":
		var dir, file string
		if strings.HasPrefix(flagFile, "/") {
			file = flagFile
		} else {
			dir, err = os.Getwd()
			if err != nil {
				logger.Fatalf("_fatal||reason=%+v", err)
			}
			file = dir + "/" + flagFile
		}
		arr := strings.Split(file, "/")
		dir = strings.Join(arr[:len(arr)-1], "/")
		generateForFile(dir, file)

	default:
		var dir string
		if flagDirectory == "." {
			dir, err = os.Getwd()
			if err != nil {
				logger.Fatalf("_fatal||reason=%+v", err)
			}
		} else if strings.HasPrefix(flagDirectory, "/") {
			dir = flagDirectory
		} else {
			wd, err := os.Getwd()
			if err != nil {
				logger.Fatalf("_fatal||reason=%+v", err)
			}
			dir = wd + "/" + flagDirectory
		}
		generateForDirectory(dir)
	}
}

func generateForDirectory(dir string) {
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		logger.Fatalf("_fatal||reason=%+v", err)
	}

	for _, fileInfo := range fileInfos {
		if !strings.HasSuffix(fileInfo.Name(), ".go") {
			continue
		}

		generateForFile(dir, dir+"/"+fileInfo.Name())
	}
}

func generateForFile(dir, file string) {
	arr := strings.Split(dir+file, "/")
	filename := arr[len(arr)-1]

	fileItem := parseSrcFile(file)
	if len(fileItem.InterfaceItems) == 0 {
		return
	}
	reader := buildDistFile(fileItem)

	if flagTarget == "stdout" {
		if _, err := io.Copy(os.Stdout, reader); err != nil {
			logger.Fatalf("_fatal||reason=%+v", err)
		}

		return
	}

	outDir := dir + "/" + flagOutPkg
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
		strings.Replace(file, dir, "", -1)[1:],
		strings.Replace(outDir+"/"+filename, dir, "", -1)[1:],
	)
}
