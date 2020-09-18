package main

import (
	"io"
	"log"
	"os"
	"strings"
)

// go run main.go
func main() {
	name := "Emannuel Loxt"

	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="utf-8">
		<title>Hello World!</title>
	</head>
	<body>
		<h1>` + name + `</h1>
	</body>
	`

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(tpl))
}
