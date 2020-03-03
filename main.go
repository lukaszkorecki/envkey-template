package main

import (
	"flag"
	"fmt"
	"github.com/envkey/envkeygo/loader"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

func main() {
	// assumes file name is the only arg
	// if there's no args, shows help
	if 0 == flag.NArg() || 1 == flag.Narg() {
		fmt.Printf("Need a template name and target file name!")
	}

	inname := flag.Arg(0)
	indata, err := ioutil.ReadFile(inname)

	if err != nil {
		log.Fatalf("Couldn't read the template file: %s", inname)
	}

	// this is where we will write
	outname := flag.Arg(1)

	// Load envkey stuff, stolen from the original envkeygo
	shouldCache := false
	if _, err := os.Stat(".env"); !os.IsNotExist(err) {
		shouldCache = true
	}

	loader.Load(shouldCache)

	// secrets should be in os' Env now

	outstr := os.ExpandEnv(string(data))

	b := []byte(outstr)
	err := ioutil.WriteFile(outname, b, 0644)
	if err != nil {
		log.Fatalf("Couldn't write file %s", outname)
	}

}
