package main

import (
	"log"
	"flag"
	"os/exec"
	"os"
	"fmt"
	"errors"
	"bytes"
	"strings"
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
		log.Println("You must supply a package name")
		os.Exit(1)
	}

	if *typeName == "" {
		log.Println("You must supply a type name")
		os.Exit(1)
	}

	if *genericName == "" {
		log.Println("You must supply " + genericNameMsg)
		os.Exit(1)
	}
}

func BuildUrl(s string) string {
	url := "https://raw.github.com/drhodes/gts/master/%s/%s.go"
	url = fmt.Sprintf(url, s, s)
	return url
}

func PageOk(url string) error {
	cmd := exec.Command (		
		"curl",
		"-4",
		`--write-out`,
		"%{http_code}",
		`--silent`,
		`--output`,
		`/dev/null`,
		url,
	)

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		msg := fmt.Sprintf("Couldn't evaluate if template is available\n")
		return errors.New(msg + err.Error())
	}
	if out.String() != "200" {
		msg := fmt.Sprintf("curl encountered error: %s", out.String())
		return errors.New(msg)
	}
	return nil
}


func GetTemplate(s string) (string, error) {
	url := BuildUrl(s)
	
	err := PageOk(url) 
	if err != nil {
		msg := fmt.Sprintf("Fetching template failed:\n")
		return "", errors.New(msg + err.Error())
	} else {
		log.Println("Template Found: ", url)
	}
	
	log.Println("Fetching: " + url)
	// so I think github has misconfigured ipv6, it seems to be timing out
	// I don't know how to force http.Get() to resolve only ipv4
	// without doing this, the reponse time is multiseconds with redirect.
	cmd := exec.Command("curl", "-4", url)

	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		msg := fmt.Sprintf("Could not be download template: %s\n", s)
		return "", errors.New(msg + err.Error())
	}
	return out.String(), nil
}


func Generify(template string) string {
	// TODO: buildin for multi parameter generics when that's needed
	return strings.Replace(template, `𝞃`, *typeName, -1)
}

func main() {
	ParseFlags()

	template, err := GetTemplate(*genericName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(Generify(template))

	// need to change the package name
	// need to remove the dummy lines


}
