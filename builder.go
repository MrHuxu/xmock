package main

import (
	"bytes"
	"go/format"
	"html/template"
	"io"
	"io/ioutil"
)

var (
	tmplBytes, _ = ioutil.ReadFile("./dist.tmpl")
	tmpl         = template.Must(template.New("dist").Parse(string(tmplBytes)))
)

// BuildOption ...
type BuildOption struct {
	Directory string
	File      string
	OutPkg    string
}

type buildArgs struct {
	FileItem    *FileItem
	BuildOption *BuildOption
}

var buildOption BuildOption

func buildDistFile(fileItem *FileItem, buildOption *BuildOption) io.Reader {
	var tmp bytes.Buffer
	if err := tmpl.Execute(&tmp, buildArgs{
		FileItem:    fileItem,
		BuildOption: buildOption,
	}); err != nil {
		logger.Fatalf("_fatal||reason=%+v", err)
	}

	bs, err := format.Source(tmp.Bytes())
	if err != nil {
		logger.Fatalf("_fatal||reason=%+v", err)
	}

	var buf bytes.Buffer
	if _, err := buf.Write(bs); err != nil {
		logger.Fatalf("_fatal||reason=%+v", err)
	}
	return &buf
}
