package main

import (
	"flag"
	"go/importer"
	"html/template"
	"log"
	"os"
	"strings"
)

var tmplStr = `

// {{.Name}} writes a response to the Responder with the status http.Status{{.Name}}.
func (r Responder) {{.Name}}() {r.write(http.Status{{.Name}})}

`

func main() {
	outFile := flag.String("out-file", "", "file to write output to")
	flag.Parse()

	pkgHTTP, err := importer.Default().Import("net/http")
	if err != nil {
		log.Fatal(err)
		return
	}

	tmpl, err := template.New("helper").Parse(tmplStr)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(*outFile)
	if err != nil {
		log.Fatalf("creating file %+v", err)
	}
	defer f.Close()

	f.WriteString("package apiresponse\n")

	for _, name := range pkgHTTP.Scope().Names() {
		if strings.HasPrefix(name, "Status") {
			strippedName := strings.SplitAfter(name, "Status")[1]
			// skip the StatusText method
			if strippedName == "Text" {
				continue
			}
			err := tmpl.Execute(f, struct{ Name string }{Name: strippedName})
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
