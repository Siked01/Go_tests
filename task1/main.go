package main

import (
	"flag"

	"github.com/Siked01/Go_tests/uniq"
)

func main() {

	var opts uniq.flagsMain
	uniq.Init(&opts)
	flag.Parse()

}
