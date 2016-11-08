package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"

	"github.com/Joker/jade"
)

var	funcMap = template.FuncMap{
	"include": func(includePath string) (template.HTML, error) {
		include_tpl, err := jade.ParseFile(includePath)
		if err != nil {
			log.Printf("\nParseFile error: %v", err)
		}
		log.Printf("%s\n\n", include_tpl)

		go_partial_tpl, _ := template.New("partial").Parse(include_tpl)

		buf := new(bytes.Buffer)
		go_partial_tpl.Execute(buf, "")
		return template.HTML(buf.String()), nil

	},
}

func handler(w http.ResponseWriter, r *http.Request) {
	layout, err := jade.ParseFile("layout.jade")
	if err != nil {
		log.Printf("\nParseFile error: %v", err)
	}
	log.Printf("%s\n\n", layout)

	index, err := jade.ParseFile("index.jade")
	if err != nil {
		log.Printf("\nParseFile error: %v", err)
	}
	log.Printf("%s\n\n", index)

	//

	go_tpl, err := template.New("layout").Funcs(funcMap).Parse(layout)
	go_tpl.New("index").Parse(index)
	if err != nil {
		log.Printf("\nTemplate parse error: %v", err)
	}

	err = go_tpl.Execute(w, "")
	if err != nil {
		log.Printf("\nExecute error: %v", err)
	}
}

func main() {
	log.Println("open  http://localhost:8080/")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
