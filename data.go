package main

import (
	"fmt"
)

type (
	Material struct {
		Name        string           `yaml:"name"`
		Columns     []ColumnMaterial `yaml:"columns"`
		PackageName string
		KeyType     string
	}

	ColumnMaterial struct {
		Name  string `yaml:"name"`
		Type  string `yaml:"type"`
		IsKey bool   `yaml:"is_key"`
	}
)

func (m *Material) Validate() error {
	var countPK int
	for _, c := range m.Columns {
		if c.IsKey {
			m.KeyType = c.Type
			countPK++
		}
	}
	if countPK != 1 {
		return fmt.Errorf("Validation Error: The number of primary key must be one, but there are %d primary keys.", countPK)
	}
	return nil
}
