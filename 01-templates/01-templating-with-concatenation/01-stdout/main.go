package main

import "fmt"

// go run main.go > index.html
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

	fmt.Println(tpl)
}
