package main

import (
	"bufio"
	"github.com/kataras/iris/core/errors"
	"go/build"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func queryGoPath() (pth string) {
	pth = os.Getenv("GOPATH")
	if pth == "" {
		pth = build.Default.GOPATH
	}
	return
}

func queryPkgPath(gopath, f string) (defaultPkg string, err error) {
	absPath, err := filepath.Abs(f)
	if err != nil {
		return
	}
	absPath = filepath.Dir(absPath)
	if !strings.Contains(absPath, gopath) {
		err = errors.New("out gopath")
		return
	}
	defaultPkg = strings.Replace(absPath, gopath+"/src/", "", 1)
	return
}

func queryPkgName(gopath, f string) (pkg string, pkgPath string, err error) {
	pkgPath, err = queryPkgPath(gopath, f)
	if err != nil {
		return
	}
	fp, err := os.Open(f)
	if err != nil {
		return
	}
	defer fp.Close()
	scanner := bufio.NewScanner(fp)
	reg := regexp.MustCompile(`^( |\t)*package( |\t)+[a-zA-Z0-9._]*( |\t)*$`)
	for scanner.Scan() {
		line := scanner.Text()
		ss := reg.FindAllString(line, -1)
		if len(ss) > 0 {
			pkg = strings.Replace(ss[0], "package", "", 1)
			break
		}
	}
	pkg = strings.Split(pkg, ".")[0]
	pkg = strings.TrimSpace(pkg)
	ps := strings.Split(pkgPath, "/")
	ps = append(ps[:len(ps)-1], pkg)
	pkgPath = strings.Join(ps, "/")
	return
}
