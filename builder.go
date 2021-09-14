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

type buildArgs struct {
	OutPkg   string
	FileItem *FileItem
}

func buildDistFile(fileItem *FileItem) io.Reader {
	var tmp bytes.Buffer
	if err := tmpl.Execute(&tmp, buildArgs{
		OutPkg:   flagOutPkg,
		FileItem: fileItem,
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
