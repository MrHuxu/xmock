package main

import "log"

func main() {
	fileItem := parseFile("./tmp/usage.go")
	log.Println(fileItem)
}
