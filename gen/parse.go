package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func loadMaterial(fileName string) (*Material, error) {
	switch filepath.Ext(fileName) {
	case ".yaml":
		var material *Material
		b, err := os.ReadFile(fileName)
		if err != nil {
			return nil, err
		}
		if err := yaml.Unmarshal(b, &material); err != nil {
			return nil, err
		}
		if err := material.Validate(); err != nil {
			return nil, err
		}
		material.ToUpperCamel()
		return material, nil
	default:
		return nil, fmt.Errorf("File type '%s' not yet supported: %s", filepath.Ext(fileName), fileName)
	}
}
