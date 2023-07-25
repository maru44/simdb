package main

import (
	"fmt"
)

type (
	// material of table
	Material struct {
		// table name: required
		Name string `yaml:"name" toml:"name"`
		// table struct should be private or not
		IsPrivate bool `yaml:"is_private" toml:"is_private"`
		// map's value is pointer
		IsPointer bool `yaml:"is_pointer" toml:"is_pointer"`
		// package name of generated file name: required
		PackageName string `yaml:"package_name" toml:"package_name"`
		// type of key (this table's primary and unique key): required
		KeyType string `yaml:"key_type" toml:"key_type"`
		// columns
		Columns []ColumnMaterial `yaml:"columns" toml:"columns"`
	}

	// material of column
	ColumnMaterial struct {
		// column name: required
		Name string `yaml:"name" toml:"name"`
		// type of column: required
		Type string `yaml:"type" toml:"type"`
		// the field should be private or not
		IsPrivate bool `yaml:"is_private" toml:"is_private"`
	}

	// interface have name and isPrivate
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
