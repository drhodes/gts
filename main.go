package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

const URL = "https://raw.github.com:443/drhodes/gts/master/%s/%s.go"

func BuildUrl(s string) string {
	url := fmt.Sprintf(URL, s, s)
	return url
}

func PageOk(url string) error {
	cmd := exec.Command(
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
		msg := "Template Found: " + url + "\n"
		os.Stderr.Write([]byte(msg))
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
	// TODO: build multi parameter generics when that's needed	
	return strings.Replace(template, `α`, *typeName, -1)
}

func ReplacePackageName(template string) (string, error) {
	// golang regex tool
	// http://regoio.herokuapp.com/ <- highly recommended :)
	pattern := `package ([\pL_])+([\pL\pN_])*`
	re := regexp.MustCompile(pattern)
	ends := re.FindStringIndex(template)

	if ends == nil {
		msg := "package declaration not found in template"
		return "", errors.New(msg)
	}

	head := template[0:ends[0]]
	tail := template[ends[1]:len(template)]
	mid := "package " + (*pkgName)

	return head + mid + tail, nil
}

// remove every line in the dummy section
// what is the dummy section you ask?
// it's just a part of the template that implements
// a fake type so the template can compile and be tested.
func RemoveDummySection(template string) string {
	lines := strings.Split(template, "\n")
	result := []string{}

	capturing := true
	for n := range lines {
		if strings.Contains(lines[n], "dummy start") {
			capturing = false
		}

		if capturing {
			result = append(result, lines[n])
		}

		if strings.Contains(lines[n], "dummy end") {
			capturing = true
		}
	}
	return strings.Join(result, "\n")
}

func main() {
	ParseFlags()

	// template, err := GetTemplate(*genericName)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// template = Generify(template)

	// template, err = ReplacePackageName(template)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// template = RemoveDummySection(template)
	// fmt.Println(template)

	err := ManglePackage(manglePath)
	log.Println(manglePath)
	if err != nil {
		log.Fatal(err)
	}
}



















