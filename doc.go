//Copyright (c) 2017 Cesar Gimenes - MIT License
//
//A small utility that looks for projects in your
//GOPATH and returns the full path to the project directory.
//
//Parameters
//  Find package/software and stop on the first occurrence
//
//      lpk packageName
//  Lists all occurrences including in the vendor directory
//
//      lpk -list:all packageName
//
//  Lists all occurrences except in the vendor directory
//
//      lpk -list:all,skipvendor packageName
package main
