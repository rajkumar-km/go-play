package main

import (
	"fmt"
	"sync"
)

// Unnamed struct that embeds a sync.Mutex
// Generally, we can not add methods to unnamed structs
// But embedding an indirect way of adding methods to unnamed struct.
var cache = struct {
	sync.Mutex // embed Mutex to guard data
	data       map[string]string
}{
	data: make(map[string]string),
}

func GetCache(key string) string {
	// Lock and Unlock are the methods derived from sync.Mutex
	cache.Lock()
	defer cache.Unlock()
	return cache.data[key]
}

func SetCache(key string, val string) {
	cache.Lock()
	defer cache.Unlock()
	cache.data[key] = val
}

func DemoUnnamedStructEmbed() {
	SetCache("a", "97")
	fmt.Println(GetCache("a"))
}
