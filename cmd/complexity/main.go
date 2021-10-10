package main

import (
	"github.com/teststudyacc2/CodeParser"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(complexity.Analyzer) }
