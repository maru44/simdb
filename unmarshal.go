package main

import (
	"encoding/json"
	"os"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

var unmarshalerByExt = map[string]func(fileName string, material **Material) error{
	".yaml": unmarshalYaml,
	".yml":  unmarshalYaml,
	".toml": unmarshalToml,
	".json": unmarshalJSON,
}

func unmarshalYaml(fileName string, material **Material) error {
	b, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(b, material)
}

func unmarshalToml(fileName string, material **Material) error {
	b, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	return toml.Unmarshal(b, material)
}

func unmarshalJSON(fileName string, material **Material) error {
	b, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, material)
}
