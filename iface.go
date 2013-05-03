package main

import (
	// "log"
	// "os"
	"flag"
	"errors"
)

var (
	// -------------------------------------------------------
	// 3 input options

	// -cu
	// -custom-url
	customUrl string

	// -cf
	// -custom-file
	customFile string

	// -stdin
	stdin bool

	// -------------------------------------------------------
	// type-params checks the gts file in the repository
	// to ensure that the correct number of type parameters were
	// provided.  If not, gts emits an error message and exits
	// with failure code
	// -type-params = Type1 Type2 ... TypeN
	typeParams string

	// -list
	// list the available containers
	// list bool

	// -mangle=pkgpath	
	manglePath string


	pkgNameMsg = "The package name used for the generated code."
	pkgName    = flag.String("pkg", "", pkgNameMsg)

	typeNameMsg = "The name the type to be inserted"
	typeName    = flag.String("type", "", typeNameMsg)

	genericNameMsg = "The name of the generic template container type"
	genericName    = flag.String("gen", "", genericNameMsg)
)

func initFlags() {
	const customUrlUsage = 
		"Use a custom url to specify the path of a template package"
	flag.StringVar(&customUrl, "custom-url", "", customUrlUsage)
	flag.StringVar(&customUrl, "cu", "", customUrlUsage+" (shorthand)")

	const customFileUsage = 
		"custom file use a custom file pointing to the path of a template"
	flag.StringVar(&customFile, "custom-file", "", customFileUsage)
	flag.StringVar(&customFile, "cf", "", customFileUsage+" (shorthand)")

	const stdinUsage = 
		"custom file use a custom file pointing to the path of a template"
	flag.BoolVar(&stdin, "stdin", false, stdinUsage)

	const typeParamsUsage = 
		"specify the types, seperated by spaces to insert in the template"
	flag.StringVar(&typeParams, "type-params", "", typeParamsUsage)

	const manglePathUsage =
		"specify the path of a package to mangle, this operation is non destructive"

	flag.StringVar(&manglePath, "mangle", "", manglePathUsage)
}

func verifyFlags() error {
	bs := 0
	
	if customUrl != "" { bs++ }
	if stdin { bs++ }
	if customFile != "" { bs++ }
	
	switch {
	case bs == 0:
		// no input options were provided.
		return errors.New("You must supply an input option: -cu, -cf, or -stdin")
	case bs > 1:
		// ensure that only one of the input options are specified.
		return errors.New("Too many input options were provided")
	}	
	// a host of other possible issues should be handled here.
	return nil
}

func ParseFlags() {
	initFlags()
	flag.Parse()
	
	

	// if *pkgName == "" {
	// 	log.Println("You must supply a package name")
	// 	os.Exit(1)
	// }

	// if *typeName == "" {
	// 	log.Println("You must supply a type name")
	// 	os.Exit(1)
	// }

	// if *genericName == "" {
	// 	log.Println("You must supply " + genericNameMsg)
	// 	os.Exit(1)
	// }	
}

