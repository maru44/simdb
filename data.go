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
		KeyType     string           `mapstructure:"key_type"`
	}

	ColumnMaterial struct {
		Name      string `mapstructure:"name"`
		Type      string `mapstructure:"type"`
		IsPrivate bool   `mapstructure:"is_private"`
	}

	nameAndPrivate interface {
		getName() string
		getPrivate() bool
	}
)

func (m *Material) validate() error {
	if m.Name == "" {
		return fmt.Errorf("Validation Error: The table name is required")
	}
	if m.PackageName == "" {
		return fmt.Errorf("Validation Error: Package name is required")
	}
	if m.KeyType == "" {
		return fmt.Errorf("Validation Error: Key type is required")
	}

	for _, c := range m.Columns {
		if c.Name == "" {
			return fmt.Errorf("Validation Error: The column name is required")
		}
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
