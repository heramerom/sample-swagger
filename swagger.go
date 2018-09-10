package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/heramerom/sample-swagger/template"
	"log"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"
)

var verbose = flag.Bool("v", false, "verbose")

var out = flag.String("o", "sample-swagger", "out put dir")

func debug(msg ...interface{}) {
	if *verbose {
		fmt.Println(msg...)
	}
}

func main() {

	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("please input path")
		os.Exit(1)
	}
	api := NewApi()

	for _, path := range args {
		b, err := isDirectory(path)
		if err != nil {
			continue
		}
		if b {
			filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
				if !isSourceFile(path) {
					return nil
				}
				parseFile(path, api)
				return nil
			})
		} else {
			if !isSourceFile(path) {
				continue
			}
			parseFile(path, api)
		}
	}

	dumpFile(api)

	fmt.Println("success!")

}

func dumpFile(api *Api) {

	err := os.MkdirAll(*out, 0777)
	if err != nil {
		fmt.Printf("mkdir error: %s", err.Error())
		os.Exit(1)
	}

	js := api.Json()

	tm, _ := base64.StdEncoding.DecodeString(templateModelBase64)
	err = writeFile("model.go", tm)
	if err != nil {
		fmt.Println("write file error:", err.Error())
		os.Exit(1)
	}

	pm, _ := base64.StdEncoding.DecodeString(templateParseBase64)
	err = writeFile("parse.go", pm)
	if err != nil {
		fmt.Println("write file error:", err.Error())
		os.Exit(1)
	}

	sm, _ := base64.StdEncoding.DecodeString(templateServerBase64)
	err = writeFile("server.go", sm)
	if err != nil {
		fmt.Println("write file error:", err.Error())
		os.Exit(1)
	}

	s2m, _ := base64.StdEncoding.DecodeString(templateServer2Base64)
	err = writeFile("server2.go", s2m)
	if err != nil {
		fmt.Println("write file error:", err.Error())
		os.Exit(1)
	}

	str := templateVars
	str = strings.Replace(str, "{{GeneratorJson}}", "`"+js+"`", 1)
	str = strings.Replace(str, "{{Imports}}", api.DefinitionImports(), 1)
	str = strings.Replace(str, "{{GeneratorModels}}", api.DefinitionObjects(), 1)
	err = writeFile("vars.go", []byte(str))
	if err != nil {
		fmt.Println("write file error: ", err.Error())
		os.Exit(1)
	}
}

func writeFile(name string, data []byte) error {
	f, err := os.OpenFile(path.Join(*out, name), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteAt(data, 0)
	if err != nil {
		return err
	}
	return f.Sync()
}

func isDirectory(f string) (bool, error) {
	fi, err := os.Stat(f)
	if err != nil {
		return false, err
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		return true, nil
	case mode.IsRegular():
		return false, nil
	}
	return false, nil
}

func isSourceFile(f string) bool {
	if strings.HasSuffix(f, "_test.go") {
		return false
	}
	if strings.HasSuffix(f, ".go") {
		return true
	}
	return false
}

func parseFile(f string, api *Api) {
	file, err := os.Open(f)
	if err != nil {
		log.Printf("open file error: %s, error: %s", f, err.Error())
		return
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	var currentRouter = emptyRouter
	num := 0
	for fileScanner.Scan() {
		num++
		line := fileScanner.Text()
		if !strings.Contains(line, "@sw:") {
			continue
		}
		line = strings.TrimLeft(line, " \t")
		if !strings.HasPrefix(line, "//") {
			continue
		}
		line = strings.Replace(line, "//", "", 1)
		scanner := newScanner(line, f, num)
		cmd := scanner.nextString(' ', '\t')
		cmd = strings.Replace(cmd, "@sw:", "", 1)
		switch strings.ToLower(cmd) {
		case "router", "r":
			if !reflect.DeepEqual(currentRouter, emptyRouter) {
				api.AddRouters(currentRouter)
			}
			currentRouter = parseRouter(scanner)
		case "param", "p":
			if reflect.DeepEqual(currentRouter, emptyRouter) {
				continue
			}
			currentRouter.params = append(currentRouter.params, parseParam(scanner))
		case "response", "resp", "res", "re":
			if reflect.DeepEqual(currentRouter, emptyRouter) {
				continue
			}
			currentRouter.response = append(currentRouter.response, parseResponse(scanner))
		case "model", "m":
			def := parseModel(scanner)
			if reflect.DeepEqual(def, definition{}) {
				continue
			}
			api.AddDefinitions(def)
		case "swagger":
			api.swagger.Swagger = scanner.nextString()
		case "info", "i":
			parseInfo(api, scanner)
		case "basepath":
			api.swagger.BasePath = scanner.nextString()
		case "host":
			api.swagger.Host = scanner.nextString()

		default:
			debug("file:", f, "line:", num, "unsupport command:", cmd)
		}
	}
	if !reflect.DeepEqual(currentRouter, emptyRouter) {
		api.AddRouters(currentRouter)
	}
}
func parseInfo(api *Api, s *Scanner) {
	if api.swagger.Info == nil {
		api.swagger.Info = &template.Info{}
	}
	next := s.nextString(',')
	line := s.nextString()
	switch strings.ToLower(next) {
	case "description", "desc":
		api.swagger.Info.Description = line
	case "version", "v":
		api.swagger.Info.Version = line
	case "title":
		api.swagger.Info.Title = line
	case "termsOfService":
		api.swagger.Info.TermsOfService = line
	case "contact.email":
		api.swagger.Info.Contact.Email = line
	case "license.name":
		api.swagger.Info.License.Name = line
	case "license.url":
		api.swagger.Info.License.URL = line
	}
}
