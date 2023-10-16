package main

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/Masterminds/sprig"
)

func main() {
	t, err := template.New("hello").Funcs(sprig.FuncMap()).Parse(`{{if eq .input.version_type "hotfix"}}master{{else if eq .input.version_type "feature"}}develop{{else}}master{{end}}*{{substr 8 -1 .input.version_name}}`)
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(nil)
	t.Execute(buf, map[string]interface{}{
		"input": map[string]interface{}{
			"version_type": "release",
			"version_name": "release_1.0.0",
		},
	})
	fmt.Printf("%q\n", buf.String())
}
