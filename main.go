package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/alexflint/go-arg"
)

// command positional arguments
type commandArgs struct {
	/* generated */

	// filename generated
	File string `arg:"positional" default:"db.go"`
	// dir name generated
	Dir string
	// package name generated
	Package string `default:"main"`

	/* config */

	// config file without extension
	// The priority is as follows.
	// yaml, yml, toml, json
	Config string `default:"simdb"`
}

var configExtensions = []string{
	".yaml", ".yml", ".toml", ".json",
}

func main() {
	var args commandArgs
	arg.MustParse(&args)

	var m *Material
	var parsed bool
	fileByExt, err := getFileByExtensionFromFileName(".", args.Config)
	for _, e := range configExtensions {
		if file, ok := fileByExt[e]; ok {
			fn := unmarshalerByExt[e]
			fn(file, &m)
			parsed = true
			break
		}
	}
	if !parsed {
		log.Fatal("failed to read config file")
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
	outputPath := filepath.Join(pkgDir, args.Dir, args.File)
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

func getFileByExtensionFromFileName(dirPath, fileNameWithoutExt string) (map[string]string, error) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	out := make(map[string]string)
	for _, file := range files {
		fileName := file.Name()
		if strings.TrimSuffix(fileName, filepath.Ext(fileName)) == fileNameWithoutExt {
			out[filepath.Ext(fileName)] = fileName
		}
	}
	return out, nil
}
