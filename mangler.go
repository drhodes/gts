package main

// mangle private names
// 
// __gts_pkgname_name

import (
	"fmt"
	"go/parser"
	"go/token"
	"go/ast"
	"go/printer"
	"os"
	"strings"
	//	"log"	
)

func MangleIdent(pkgname, ident string) string {
	return fmt.Sprintf("__gts_%s_%s", pkgname, ident)
}

var InTopLevel func(s string) bool
var PkgName func() string

type mangler struct{}

func (m mangler) Visit(n ast.Node) ast.Visitor {
	switch n.(type) {
	case *ast.Ident:		
		ident := n.(*ast.Ident)
		// only mangle names that are declared in the toplevel
		// and aren't exported.
		if !ident.IsExported() && InTopLevel(ident.Name) {			
			ident.Name = MangleIdent(PkgName(), ident.Name)
		}
	}
	return m
}


func FixInTopLevel(pkgs map[string]*ast.Package) {
	// get a list of the top level names.
	toplevel := []string{}
	for _, pkg := range pkgs {
		for _, f := range pkg.Files {			
			for _, o := range f.Scope.Objects {
				toplevel = append(toplevel, o.Name)
			}
		}
	}

	// fix InTopLevel
	InTopLevel = func (s string) bool {
		for _, x := range toplevel { 
			if s == x {
				return true
			}
		}		
		return false
	}
}

func ManglePackage(pkgName string) error {
	fs := new(token.FileSet)
	filter := func(info os.FileInfo) bool {
		return strings.HasSuffix(info.Name(), ".go")
	}

	pkgs, err := parser.ParseDir(fs, pkgName, filter, parser.AllErrors)
	if err != nil {
		return Err(err, "Mangler failed to parse")
	}

	FixInTopLevel(pkgs)

	for name, pkg := range pkgs {
		PkgName = func() string { return name }
		ast.Walk(mangler{}, pkg)

		// extract the imports
		// sort | uniq < imports

		// merge the files

		// pretty print the new AST 
		fmt.Println(name, fs, pkg)
		for _, f := range pkg.Files { 		
			printer.Fprint(os.Stdout, fs, f)
		}
	}



	// include a seperator
	// a new file called <container>_pgkname.go
	// it will include tests.

	return nil
}


















