package main

import (
	"fmt"
)

type (
	Material struct {
		Name        string           `mapstructure:"name"`
		Columns     []ColumnMaterial `mapstructure:"columns"`
		IsPrivate   bool             `mapstructure:"is_private"`
		PackageName string           `mapstructure:"package_name"`
		KeyType     string
	}

	ColumnMaterial struct {
		Name      string `mapstructure:"name"`
		Type      string `mapstructure:"type"`
		IsPrivate bool   `mapstructure:"is_private"`
		IsKey     bool   `mapstructure:"is_key"`
	}

	nameAndPrivate interface {
		getName() string
		getPrivate() bool
	}
)

func (m *Material) validate() error {
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

func (m *Material) getName() string {
	return m.Name
}
func (m *Material) getPrivate() bool {
	return m.IsPrivate
}

func (c *ColumnMaterial) getName() string {
	return c.Name
}
func (c *ColumnMaterial) getPrivate() bool {
	return c.IsPrivate
}
