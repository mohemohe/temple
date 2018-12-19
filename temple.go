package temple

import (
	"strings"
	"text/template"
)

func Execute(templateString string, variable map[string]interface{}) (resultString string, err error) {
	var tmpl *template.Template
	tmpl, err = template.New("test").Parse(templateString)
	if err != nil {
		return resultString, err
	}

	var builder strings.Builder
	if err = tmpl.Execute(&builder, variable); err != nil {
		return resultString, err
	}
	resultString = builder.String()

	return resultString, err
}