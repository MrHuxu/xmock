package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var (
	versionPattern = regexp.MustCompile("v[0-9]+")
)

// FileItem ...
type FileItem struct {
	ImportItems    []*ImportItem
	InterfaceItems []*InterfaceItem
}

// ImportItem ...
type ImportItem struct {
	Name string
	Path string
}

// InterfaceItem ...
type InterfaceItem struct {
	Name    string
	Methods []*FuncItem
}

// FuncItem ...
type FuncItem struct {
	Name    string
	Params  []*FieldItem
	Results []*FieldItem
}

// FieldItem ...
type FieldItem struct {
	Name       string
	Type       string
	Dependency string
}

func parseFile(filePath string) *FileItem {
	var fileItem FileItem

	src, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, filePath, src, 0)
	if err != nil {
		log.Fatal(err)
	}
	offset := file.Pos()

	ast.Inspect(file, func(n ast.Node) bool {
		if t, ok := n.(*ast.ImportSpec); ok {
			fileItem.ImportItems = append(fileItem.ImportItems, parseImportSpec(t, src, offset))
		}

		if t, ok := n.(*ast.TypeSpec); ok {
			if !t.Name.IsExported() {
				return true
			}

			if _, ok := t.Type.(*ast.InterfaceType); ok {
				fileItem.InterfaceItems = append(fileItem.InterfaceItems, parseInterfaceSpec(t, src, offset))
			}
		}
		return true
	})

	return &fileItem
}

func parseImportSpec(t *ast.ImportSpec, src []byte, offset token.Pos) *ImportItem {
	var importItem ImportItem

	arr := strings.Split(strings.Replace(string(src[t.Pos()-offset:t.End()-offset]), `"`, "", -1), " ")
	importItem.Path = arr[len(arr)-1]
	if t.Name == nil {
		importItem.Name = parseDependencyName(importItem.Path)
	} else {
		importItem.Name = t.Name.Name
	}

	return &importItem
}

func parseDependencyName(path string) string {
	arr := strings.Split(path, "/")
	if versionPattern.Match([]byte(arr[len(arr)-1])) {
		return arr[len(arr)-2]
	}
	return arr[len(arr)-1]
}

func parseInterfaceSpec(t *ast.TypeSpec, src []byte, offset token.Pos) *InterfaceItem {
	interfaceItem := &InterfaceItem{Name: t.Name.Name}
	i := t.Type.(*ast.InterfaceType)
	for _, method := range i.Methods.List {
		fn, ok := method.Type.(*ast.FuncType)
		if !ok {
			continue
		}

		funcItem := &FuncItem{Name: method.Names[0].Name}

		for i, param := range fn.Params.List {
			var paramItem FieldItem
			if len(param.Names) > 0 {
				paramItem.Name = param.Names[0].Name
			} else {
				paramItem.Name = "p" + strconv.Itoa(i)
			}
			paramItem.Type = string(src[param.Type.Pos()-offset : param.Type.End()-offset])
			funcItem.Params = append(funcItem.Params, &paramItem)
		}

		for i, result := range fn.Results.List {
			var resultItem FieldItem
			if len(result.Names) > 0 {
				resultItem.Name = result.Names[0].Name
			} else {
				resultItem.Name = "r" + strconv.Itoa(i)
			}
			resultItem.Type = string(src[result.Type.Pos()-offset : result.Type.End()-offset])
			funcItem.Results = append(funcItem.Results, &resultItem)
		}

		interfaceItem.Methods = append(interfaceItem.Methods, funcItem)
	}
	return interfaceItem
}
