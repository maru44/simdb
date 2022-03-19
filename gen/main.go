package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/alexflint/go-arg"
)

type commandArgs struct {
	FileName          string `arg:"positional,required"`
	Package           string `arg:"positional,reqired"`
	GeneratedFileName string `arg:"positional,required"`
	// Targets     []string `arg:"positional"`
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

	outputPath := filepath.Join(pkgDir, args.Package, args.GeneratedFileName+".go")
	data, err := render(outputPath, m)
	if err != nil {
		log.Fatal("failed to render: ", err)
	}

	if err := os.Mkdir(filepath.Dir(outputPath), 0755); err != nil {
		log.Fatal("failed to make dir: ", err)
	}
	if err := os.WriteFile(outputPath, data, 0600); err != nil {
		log.Fatal("failed to write file: ", err)
	}
}
