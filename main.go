package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/alexflint/go-arg"
	"github.com/spf13/viper"
)

type commandArgs struct {
	/* generated */
	// filename generated
	GeneratedFileName string `arg:"positional" default:"db.go"`
	// dir name
	Dir string `arg:"positional"`
	// package name
	Package string `arg:"positional" default:"main"`

	/* config */
	// config file
	ConfigName string `arg:"positional" default:"simdb"`
}

func main() {
	// read env first
	viper.AutomaticEnv()

	var args commandArgs
	arg.MustParse(&args)

	viper.SetConfigName(args.ConfigName)
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("failed to read config: ", err)
	}

	var m *Material
	if err := viper.Unmarshal(&m); err != nil {
		log.Fatal("failed to unmarshal material: ", err)
	}

	// set package name if blank
	if m.PackageName == "" {
		m.PackageName = args.Package
	}

	if err := m.validate(); err != nil {
		log.Fatal("failed to validate: ", err)
	}

	pkgDir, err := filepath.Abs(".")
	if err != nil {
		log.Fatal("failed to get package dir: ", err)
	}
	outputPath := filepath.Join(pkgDir, args.Dir, args.GeneratedFileName)
	data, err := render(outputPath, m)
	if err != nil {
		log.Fatal("failed to render: ", err)
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
