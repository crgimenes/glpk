package main

import "testing"

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

	cfg.List = "all,error"
	err = parseListPar()
	if err == nil {
		t.Fatal("error expected")
	}

}
