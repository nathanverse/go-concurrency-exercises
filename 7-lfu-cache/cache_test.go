package main

import (
	"fmt"
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

func TestCache1(t *testing.T) {
	cache, err := NewLFUCache(3, func(key string) (string, error) {
		return "", nil
	})
	if err != nil {
		t.Fatal(err)
	}

	err = cache.Set("vu", "10")
	err = cache.Set("nghia", "20")
	err = cache.Set("luan", "5")

	for i := 0; i < 10; i++ {
		cache.Get("vu")
	}

	for i := 0; i < 9; i++ {
		cache.Get("nghia")
	}

	for i := 0; i < 8; i++ {
		cache.Get("luan")
	}

	i := 8
	for e := cache.GetBuckets().Front(); e != nil; e = e.Next() {
		fmt.Printf("Value: %v (Type: %T)\n", e.Value, e.Value)
		bucketFreq := cache.GetFreq(e)
		if bucketFreq != i {
			t.Errorf("bucketFreq should be %d, got %d", i, bucketFreq)
		}
		i += 1
	}

	err = cache.Set("xanh", "30")
	keys := cache.GetKeys()
	if slices.Contains(keys, "luan") {
		t.Errorf("keys should not contain luan")
	}
}
