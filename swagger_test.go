package main

import (
	"github.com/influxdata/influxdb/pkg/testing/assert"
	"testing"
)

func TestParseRouter(t *testing.T) {
	line := "get,post, /v1/class/detail, base, desc"
	r := parseRouter(line)
	assert.Equal(t, r.methods[0], "get")
	assert.Equal(t, r.methods[1], "post")
	assert.Equal(t, r.path, "/v1/class/detail")
	assert.Equal(t, r.tag, "base")
	assert.Equal(t, r.desc, "desc")
}
