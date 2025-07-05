package main

import (
	"slices"
	"testing"
)

func TestCache(t *testing.T) {
	cache, err := NewLFUCache(3, func(key string) (string, error) {
		return "", nil
	})
	if err != nil {
		t.Fatal(err)
	}

	err = cache.Set("vu", "10")
	err = cache.Set("nghia", "20")
	err = cache.Set("luan", "5")

	value, err := cache.Get("vu")
	if value != "10" {
		t.Errorf("value should be 10, got %s", value)
	}

	value, err = cache.Get("nghia")
	if value != "20" {
		t.Errorf("value should be 20, got %s", value)
	}

	err = cache.Set("xanh", "30")

	keys := cache.GetKeys()
	if slices.Contains(keys, "luan") {
		t.Errorf("keys should not contain luan")
	}
}
