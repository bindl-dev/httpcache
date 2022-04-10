package test_test

import (
	"testing"

	"github.com/bindl-dev/httpcache"
	"github.com/bindl-dev/httpcache/test"
)

func TestMemoryCache(t *testing.T) {
	test.Cache(t, httpcache.NewMemoryCache())
}
