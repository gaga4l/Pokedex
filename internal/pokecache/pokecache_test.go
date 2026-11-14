package pokecache

import (
	"bytes"
	"testing"
	"time"
)

var interval = 5 * time.Minute

func TestNewCache(t *testing.T) {
	cache := NewCache(interval) // Use a 1-minute interval for testing

	if cache == nil {
		t.Fatal("NewCache() returned nil")
	}

	if len(cache.entry) != 0 {
		t.Fatalf("expected empty cache, got %d entries", len(cache.entry))
	}
}

func TestCacheAddAndGet(t *testing.T) {
	cache := NewCache(interval)
	key := "pikachu"
	val := []byte("data")

	cache.Add(key, val)

	got, ok := cache.Get(key)
	if !ok {
		t.Fatalf("expected key %q to exist", key)
	}

	if !bytes.Equal(got, val) {
		t.Fatalf("expected value %q, got %q", val, got)
	}
}

func TestCacheGetMissingKey(t *testing.T) {
	cache := NewCache(interval)

	got, ok := cache.Get("missing")
	if ok {
		t.Fatal("expected missing key to return ok=false")
	}

	if got != nil {
		t.Fatalf("expected nil value for missing key, got %q", got)
	}
}

func TestCacheOverwriteExistingKey(t *testing.T) {
	cache := NewCache(interval)
	key := "bulbasaur"

	cache.Add(key, []byte("first"))
	cache.Add(key, []byte("second"))

	got, ok := cache.Get(key)
	if !ok {
		t.Fatalf("expected key %q to exist after overwrite", key)
	}

	if string(got) != "second" {
		t.Fatalf("expected overwritten value %q, got %q", "second", got)
	}
}

func TestReapLoop(t *testing.T) {
	cache := NewCache(10 * time.Millisecond)
	cache.Add("charmander", []byte("data"))
	time.Sleep(50 * time.Millisecond)

	if _, ok := cache.Get("charmander"); ok {
		t.Fatal("expected key 'charmander' to be removed after interval")
	}
}
