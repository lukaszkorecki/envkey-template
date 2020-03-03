package main

import (
	"flag"
	"github.com/envkey/envkeygo/loader"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var inName string
var outName string
var debug bool

func init() {
	flag.StringVar(&inName, "in", "", "Input template")
	flag.StringVar(&outName, "out", "", "Output file")
	flag.BoolVar(&debug, "debug", false, "Debug")
	flag.Parse()

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var pathErr error
	inName, pathErr = filepath.Abs(inName)
	check(pathErr)

	indata, err := ioutil.ReadFile(inName)

	if err != nil {
		log.Fatalf("Couldn't read the template file: %s", inName)
	}

	// this is where we will write
	outName, err = filepath.Abs(outName)

	if err != nil {
		log.Fatalf("Invalid path %s", err)
	}

	if debug {
		log.Printf("%s -> %s", inName, outName)
	}

	// Load envkey stuff, stolen from the original envkeygo
	shouldCache := false
	if _, err := os.Stat(".env"); !os.IsNotExist(err) {
		shouldCache = true
	}

	loader.Load(shouldCache)

	// secrets should be in os' Env now

	outstr := os.ExpandEnv(string(indata))

	b := []byte(outstr)

	f, e := os.Create(outName)
	defer f.Close()

	if e != nil {
		panic(e)
	}

	_, err = f.Write(b)
	if err != nil {
		log.Printf("Couldn't write file %s", outName)
		panic(err)
	}

}
