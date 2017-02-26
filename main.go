package main

import (
	"fmt"
	"go/build"
	"os"
	"path/filepath"
	"strings"

	"github.com/crgimenes/goConfig"
)

type config struct {
	PackageName string `cfg:"name"`
	GoPath      string
	ListAll     bool
	SkipVendor  bool `cfgDefault:"true"`
}

var cfg config

func visit(path string, f os.FileInfo, err error) error {
	if !f.IsDir() {
		return nil
	}

	if cfg.SkipVendor {
		if f.Name() == "vendor" {
			return filepath.SkipDir
		}
	}

	pkgName := strings.ToLower(cfg.PackageName)
	dirName := strings.ToLower(f.Name())

	if pkgName == dirName {
		fmt.Println(path)
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

	if cfg.PackageName == "" {
		println("Package name not defined.")
		goConfig.Usage()
		os.Exit(1)
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
