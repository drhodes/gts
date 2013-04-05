package main

import (
	"flag"
	"os"
	"fmt"
	//"io/ioutil"
)

var (
	pkgNameMsg = "The package name used for the generated code."
	pkgName = flag.String("pkg", "", pkgNameMsg)

	typeNameMsg = "The name the type to be inserted"
	typeName = flag.String("type", "", typeNameMsg)

	genericNameMsg = "The name of the generic template"
	genericName = flag.String("gen", "", genericNameMsg)
)

func ParseFlags() {
	flag.Parse()
	
	if !flag.Parsed() {
		flag.Usage()
	}

	if *pkgName == "" {
		fmt.Println("You must supply a package name")
		os.Exit(1)
	}

	if *typeName == "" {
		fmt.Println("You must supply a type name")
		os.Exit(1)
	}

	if *genericName == "" {
		fmt.Println("You must supply " + genericNameMsg)
		os.Exit(1)
	}
}

func OpenTemplate(s string) {
	//ioutil.ReadFile(./
}


func main() {
	ParseFlags()

	// open a template.
	// replace

	// recursively get the deps
	// deps | sort | uniq 
	// 
}
