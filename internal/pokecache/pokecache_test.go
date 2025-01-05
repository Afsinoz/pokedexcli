package pokecache

import (
	"fmt"
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
			key: "https://example.com/",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/another",
			val: []byte("moretestdata"),
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)

			if !ok {
				t.Errorf("Expected to find a key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("Expected to find a val")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Second
	const waitTime = baseTime + 5*time.Millisecond

	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdataaa"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("Expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok1 := cache.Get("https://example.com")
	if ok1 {
		t.Errorf("Expected to find no key")
		return
	}

}
