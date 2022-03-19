package main

import "fmt"

type (
	ColumnType     string
	ColumnOperator string

	IntLike interface {
		int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
	}

	Material struct {
		Name           string           `yaml:"name"`
		Columns        []ColumnMaterial `yaml:"columns"`
		PackageName    string
		PrimaryKeyType ColumnType
	}

	ColumnMaterial struct {
		Name         string     `yaml:"name"`
		Type         ColumnType `yaml:"type"`
		IsPrimaryKey bool       `yaml:"isPK"`
	}
)

const (
	ColumnTypeInt        = ColumnType("int")
	ColumnTypeInt8       = ColumnType("int8")
	ColumnTypeInt16      = ColumnType("int16")
	ColumnTypeInt32      = ColumnType("int32")
	ColumnTypeInt64      = ColumnType("int64")
	ColumnTypeUpperInt   = ColumnType("uint")
	ColumnTypeUpperInt8  = ColumnType("uint8")
	ColumnTypeUpperInt16 = ColumnType("uint16")
	ColumnTypeUpperInt32 = ColumnType("uint32")
	ColumnTypeUpperInt64 = ColumnType("uint64")
	// ColumnTypeBool       = ColumnType("bool")

	ColumnOperatorEqual              = ColumnOperator("EQ")
	ColumnOperatorNotEqual           = ColumnOperator("NEQ")
	ColumnOperatorLessThan           = ColumnOperator("LT")
	ColumnOperatorLessThanOrEqual    = ColumnOperator("LTE")
	ColumnOperatorGreaterThan        = ColumnOperator("GT")
	ColumnOperatorGreaterThanOrEqual = ColumnOperator("GTE")
	// ColumnOperatorNone = ColumnOperator
)

func (m *Material) Validate() error {
	var countPK int
	for _, c := range m.Columns {
		if c.IsPrimaryKey {
			m.PrimaryKeyType = c.Type
			countPK++
		}
		if err := c.Validate(); err != nil {
			return err
		}
	}
	if countPK != 1 {
		return fmt.Errorf("Validation Error: The number of primary key must be one, but there are %d primary keys.", countPK)
	}
	return nil
}

func (c *ColumnMaterial) Validate() error {
	switch c.Type {
	case ColumnTypeInt:
		return nil
	case ColumnTypeInt8:
		return nil
	case ColumnTypeInt16:
		return nil
	case ColumnTypeInt32:
		return nil
	case ColumnTypeInt64:
		return nil
	case ColumnTypeUpperInt:
		return nil
	case ColumnTypeUpperInt8:
		return nil
	case ColumnTypeUpperInt16:
		return nil
	case ColumnTypeUpperInt32:
		return nil
	case ColumnTypeUpperInt64:
		return nil
	default:
		return fmt.Errorf("Validation Error: Type not supported: %s", c.Type)
	}
}
