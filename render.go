package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"text/template"

	"github.com/stoewer/go-strcase"
	"golang.org/x/tools/imports"
)

//go:embed template.tmpl
var renderTemplate string

func render(filePath string, m *Material) ([]byte, error) {
	funcMap := map[string]interface{}{
		"camel":      camelName,
		"upperCamel": upperCamel,
	}

	tmpl, err := template.New("").Funcs(funcMap).Parse(renderTemplate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse tempalte: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, m); err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}

	out, err := imports.Process(filePath, buf.Bytes(), &imports.Options{
		FormatOnly: false,
		Comments:   true,
		TabIndent:  true,
		TabWidth:   8,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to process goimports: %w", err)
	}
	return out, nil
}

func camelName(m nameAndPrivate) string {
	s := m.getName()
	if len(s) == 0 {
		return s
	}
	if m.getPrivate() {
		return strcase.LowerCamelCase(s)
	}
	return strcase.UpperCamelCase(s)
}

func upperCamel(s string) string {
	if len(s) == 0 {
		return s
	}
	return strcase.UpperCamelCase(s)
}
