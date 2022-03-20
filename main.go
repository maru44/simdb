package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/alexflint/go-arg"
)

type commandArgs struct {
	// Paresed file name yaml
	FileName string `arg:"positional,required"`
	// package name
	Package string `arg:"positional,required"`
	// filename generated
	GeneratedFileName string `arg:"positional,required"`
	// dir name
	Dir string `arg:"positional"`
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
	// set package name
	m.PackageName = args.Package

	outputPath := filepath.Join(pkgDir, args.Dir, args.GeneratedFileName)
	data, err := render(outputPath, m)
	if err != nil {
		log.Fatal("failed to render: ", err)
	}

	if _, err := os.Stat(filepath.Join(pkgDir, args.Dir)); os.IsNotExist(err) {
		if err := os.Mkdir(filepath.Dir(outputPath), 0755); err != nil {
			log.Fatal("failed to make dir: ", err)
		}
	}
	if err := os.WriteFile(outputPath, data, 0600); err != nil {
		log.Fatal("failed to write file: ", err)
	}
}
