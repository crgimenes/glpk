package main

import (
	"fmt"
	"go/build"
	"os"
	"path/filepath"

	"github.com/crgimenes/goConfig"
)

type config struct {
	PackageName string `cfg:"pn"`
	GoPath      string
}

var cfg config

func visit(path string, f os.FileInfo, err error) error {
	if !f.IsDir() {
		return nil
	}

	if cfg.PackageName == f.Name() {
		println(path)
	}

	return nil
}

func main() {

	cfg = config{}
	err := goConfig.Parse(&cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	if cfg.GoPath == "" {
		cfg.GoPath = build.Default.GOPATH
	}

	root := cfg.GoPath + "/src"

	_, err = os.Stat(root)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = filepath.Walk(root, visit)
	if err != nil {
		fmt.Println(err)
		return
	}
}
