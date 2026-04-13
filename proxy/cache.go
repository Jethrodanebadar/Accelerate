package proxy

import (
	"net/http"
)

type CacheEntry struct {
	StatusCode int
	Headers    http.Header
	Body       []byte
}
var MEMORY_CACHE = make(map[string]CacheEntry)
