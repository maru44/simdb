package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"text/template"

	"golang.org/x/tools/imports"
)

//go:embed template.tmpl
var renderTemplate string

func render(filePath string, m *Material) ([]byte, error) {
	tmpl, err := template.New("").Parse(renderTemplate)
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
