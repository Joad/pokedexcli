package pokecache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/endpoint",
			val: []byte("endpointdata"),
		},
	}
	for _, testCase := range cases {
		cache := NewCache(interval)
		cache.Add(testCase.key, testCase.val)
		val, ok := cache.Get(testCase.key)
		if !ok {
			t.Errorf("expected to find key")
		}
		if string(val) != string(testCase.val) {
			t.Errorf("expected to find value")
		}
	}
}
