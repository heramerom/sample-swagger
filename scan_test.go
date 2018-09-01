package main

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {

	a := "get, /v, hello world, this is a test."

	t.Log(strings.SplitN(a, ",", 3))

	t.Log(strings.SplitN(a, ",", -1))

	t.Log(strings.SplitN(a, ",", 4))

	t.Log(strings.SplitN(a, ",", 1))

	t.Log(strings.SplitN(a, ",", 2))

}

func TestFile(t *testing.T) {

	f, err := os.Open("/home/riki/go/src/github.com/heramerom/sample-swagger/cmd/example.txt")

	if err != nil {
		panic(err)
	}

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		t.Log(sc.Text())
	}

}
