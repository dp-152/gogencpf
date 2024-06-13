package main

import (
	"flag"

	"github.com/dp-152/gogencpf/util"
)

var generate bool
var validate string
var format bool
var count int

func init() {
	flag.BoolVar(&generate, "g", true, "")
	flag.BoolVar(&generate, "generate", true, "Generate values (default)")

	// Generate flags
	flag.BoolVar(&format, "f", false, "")
	flag.BoolVar(&format, "format", false, "Format output")
	flag.IntVar(&count, "c", 1, "")
	flag.IntVar(&count, "count", 1, "Amount to generate")

	flag.StringVar(&validate, "v", "", "")
	flag.StringVar(&validate, "validate", "", "")
}

func main() {
	flag.Parse()

	switch {
	case validate != "":
		util.Check(validate)
	case generate:
		util.Gen(format, count)
	}
}
