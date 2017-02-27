package main

import (
	"os"
	"testing"
)

func TestParseListPart(t *testing.T) {
	cfg.List = "all,skipvendor"
	err := parseListPar()
	if err != nil {
		t.Fatal(err)
	}

	if !cfg.SkipVendor {
		t.Fatal("SkipVendor false, expected true")
	}

	if !cfg.ListAll {
		t.Fatal("ListAll false, expected true")
	}

	cfg.SkipVendor = false
	cfg.List = "all"
	err = parseListPar()
	if err != nil {
		t.Fatal(err)
	}

	if cfg.SkipVendor {
		t.Fatal("SkipVendor true, expected false")
	}

	cfg.List = "all,error"
	err = parseListPar()
	if err == nil {
		t.Fatal("error expected")
	}
}

func TestFind(t *testing.T) {
	err := find()
	cfg.PackageName = ""
	if err != errNameNotDefined {
		t.Fatal("errNameNotDefined expected")
	}

	cfg.PackageName = "testpackage"
	cfg.GoPath = "./testdata"
	cfg.List = "skipvendor"
	err = find()
	if err != nil {
		t.Fatal(err)
	}

	cfg.PackageName = "testpackage"
	cfg.GoPath = "./testdata"
	cfg.List = "all,error"
	err = find()
	if err == nil {
		t.Fatal("error expected")
	}

	nameFound = false
	cfg.SkipVendor = true
	cfg.ListAll = false
	cfg.PackageName = "error name"
	cfg.GoPath = "./testdata"
	cfg.List = "skipvendor"
	err = find()
	if err == nil {
		t.Fatal("error expected")
	}

	cfg.PackageName = "testpackage"
	cfg.GoPath = ""
	cfg.List = "skipvendor"
	err = find()
	if err != nil {
		t.Fatal(err)
	}

	cfg.PackageName = "testpackage"
	cfg.GoPath = "./error dir"
	cfg.List = "skipvendor"
	err = find()
	if err == nil {
		t.Fatal("error expected")
	}
}

func TestConfigAndFind(t *testing.T) {
	os.Setenv("GOPAHT", "./testdata")
	err := configAndFind()
	if err != nil {
		t.Fatal(err)
	}
}
