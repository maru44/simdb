package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/alexflint/go-arg"
	"github.com/spf13/viper"
)

type commandArgs struct {
	// package name
	Package string `arg:"positional" default:"main"`
	// filename generated
	GeneratedFileName string `arg:"positional" default:"db.go"`
	// dir name
	Dir string `arg:"positional"`
}

func main() {
	// read env first
	viper.AutomaticEnv()

	viper.SetConfigName("simdb")
	viper.AddConfigPath(".")

	// viper.EnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("failed to read config: ", err)
	}

	var args commandArgs
	arg.MustParse(&args)

	pkgDir, err := filepath.Abs(".")
	if err != nil {
		log.Fatal("failed to get package dir: ", err)
	}

	var m *Material
	if err := viper.Unmarshal(&m); err != nil {
		log.Fatal("failed to unmarshal material: ", err)
	}

	// set package name
	m.PackageName = args.Package
	// set key type
	for _, c := range m.Columns {
		if c.IsKey {
			m.KeyType = c.Type
			break
		}
	}

	outputPath := filepath.Join(pkgDir, args.Dir, args.GeneratedFileName)
	data, err := render(outputPath, m)
	if err != nil {
		log.Fatal("failed to render: ", err)
	}

	if err := m.Validate(); err != nil {
		log.Fatal("failed to validate: ", err)
	}

	_, err = os.Stat(filepath.Join(pkgDir, args.Dir))
	if err != nil && os.IsNotExist(err) {
		if err := os.Mkdir(filepath.Dir(outputPath), 0755); err != nil {
			log.Fatal("failed to make dir: ", err)
		}
	} else if err != nil {
		log.Fatal("failed to load dir: ", err)
	}
	if err := os.WriteFile(outputPath, data, 0600); err != nil {
		log.Fatal("failed to write file: ", err)
	}
}
