package main

import (
	"fmt"
)

type (
	Material struct {
		Name        string           `mapstructure:"name"`
		Columns     []ColumnMaterial `mapstructure:"columns"`
		PackageName string           `mapstructure:"package_name"`
		KeyType     string
	}

	ColumnMaterial struct {
		Name  string `mapstructure:"name"`
		Type  string `mapstructure:"type"`
		IsKey bool   `mapstructure:"is_key"`
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
	if m.PackageName == "" {
		return fmt.Errorf("Validation Error: Package name is required")
	}
	return nil
}
