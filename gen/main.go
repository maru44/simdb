package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/alexflint/go-arg"
)

type commandArgs struct {
	FileName          string `arg:"positional,required"`
	Package           string `arg:"positional,required"`
	GeneratedFileName string `arg:"positional,required"`
}

func main() {
	var args commandArgs
	arg.MustParse(&args)

	pkgDir, err := filepath.Abs(".")
	if err != nil {
		log.Fatal("failed to get package dir: ", err)
	}

	m, err := loadMaterial(args.FileName)
	if err != nil {
		log.Fatal("failed to load material: ", err)
	}
	m.PackageName = args.Package

	outputPath := filepath.Join(pkgDir, args.Package, args.GeneratedFileName)
	data, err := render(outputPath, m)
	if err != nil {
		log.Fatal("failed to render: ", err)
	}

	if _, err := os.Stat(filepath.Join(pkgDir, args.Package)); os.IsNotExist(err) {
		if err := os.Mkdir(filepath.Dir(outputPath), 0755); err != nil {
			log.Fatal("failed to make dir: ", err)
		}
	}
	if err := os.WriteFile(outputPath, data, 0600); err != nil {
		log.Fatal("failed to write file: ", err)
	}
}
