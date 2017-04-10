# lpk
[![Build Status](https://travis-ci.org/crgimenes/lpk.svg?branch=master)](https://travis-ci.org/crgimenes/lpk)
[![Go Report Card](https://goreportcard.com/badge/github.com/crgimenes/lpk)](https://goreportcard.com/report/github.com/crgimenes/lpk)
[![codecov](https://codecov.io/gh/crgimenes/lpk/branch/master/graph/badge.svg)](https://codecov.io/gh/crgimenes/lpk)
[![GoDoc](https://godoc.org/github.com/crgimenes/lpk?status.png)](https://godoc.org/github.com/crgimenes/lpk)
[![Go project version](https://badge.fury.io/go/github.com%2Fcrgimenes%2Flpk.svg)](https://badge.fury.io/go/github.com%2Fcrgimenes%2Flpk)
[![MIT Licensed](https://img.shields.io/badge/license-MIT-green.svg)](https://tldrlegal.com/license/mit-license)

A small utility that looks for package in your GOPATH and returns the full path to the package directory.

## The problem and the solution

All my projects are in GOPATH and I got tired of typing *cd* and the full path to the package/software directory.

So I created this little utility that looks in GOPATH and returns the full path to the project, so I can create an alias to change to the project directory and also can search by the project by name, etc.

## Install

```
go get github.com/crgimenes/lpk
```

## Example of use

### Displays full path to the project

Search for the *project* in GOPATH and displays the full path type *lpk project* where *project* is the name of the package or software you are looking for. By default, lpk stops searching after it encounters the first occurrence.

```
lpk project 
```

To list all *project* occurrences use the parameter **-list** with "all" string

```
lpk -list=all project
```

This command will list all the *project* occurrences including those found in the vendor directly, to ignore the vendor add *skipvendor* to the **-list** parameter

```
lpk -list=all,skipvendor project
```

### Tips

Changes to the project directory automatically

```
cd $(lpk project)
```
---

Set up an alias for the *cd* and jump straight to your project directory.
Example, let's say you want to create an alias to jump directly to the *project*  directory. Just change project by *project* name in your GOPATH

```
alias aliasname="cd $(lpk project)"
```
---

Another tip, if you are using **macOS** the following command creates an alias to open the Finder in the *project* directory. Remember to change the word project by the name of your *project* in GOPATH

```
alias aliasname="open $(lpk project)"
```
---

## Contributing

- Fork the repo on GitHub
- Clone the project to your own machine
- Create a *branch* with your modifications `git checkout -b fantastic-feature`.
- Then _commit_ your changes `git commit -m 'Implementation of new fantastic feature'`
- Make a _push_ to your _branch_ `git push origin fantastic-feature`.
- Submit a **Pull Request** so that we can review your changes
