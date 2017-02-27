# lpk
[![Build Status](https://travis-ci.org/crgimenes/lpk.svg?branch=master)](https://travis-ci.org/crgimenes/lpk)
[![Go Report Card](https://goreportcard.com/badge/github.com/crgimenes/lpk)](https://goreportcard.com/report/github.com/crgimenes/lpk)
[![codecov](https://codecov.io/gh/crgimenes/lpk/branch/master/graph/badge.svg)](https://codecov.io/gh/crgimenes/lpk)
[![GoDoc](https://godoc.org/github.com/crgimenes/goConfig?status.png)](https://godoc.org/github.com/crgimenes/lpk)
[![Go project version](https://badge.fury.io/go/github.com%2Fcrgimenes%2Flpk.svg)](https://badge.fury.io/go/github.com%2Fcrgimenes%2Flpk)
[![MIT Licensed](https://img.shields.io/badge/license-MIT-green.svg)](https://tldrlegal.com/license/mit-license)

This is a small application to help you find the exact path of the packages or project in your GOPATH.

## Install

```
go install github.com/crgimenes/lpk
```

## Example of use

### Displays full path to the project

```
lpk packageName
```

### Changes to the project directory automatically

```
cd $(lpk packageName)
```

### Tips

Set up an alias for the cd and jump straight to your project directory.

```
alias gofn="cd $(lpk gofn)"
```

