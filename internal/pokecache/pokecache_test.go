package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	cache := NewCache(time.Millisecond)

	cases := []struct {
		inputKey string
		inputVal []byte
		expected []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
			expected: []byte("val1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("val2"),
			expected: []byte("val2"),
		},
	}

	for _, tc := range cases {
		cache.Add(tc.inputKey, []byte(tc.inputVal))
		actual, ok := cache.Get(tc.inputKey)
		if !ok {
			t.Errorf("%s not found", tc.inputKey)
			continue
		}

		if string(tc.expected) != string(actual) {
			t.Errorf("%s doesn't match expected: %s", string(tc.expected), string(tc.expected))
			continue
		}
	}
}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)
	cases := []struct {
		inputKey string
		inputVal []byte
		expected []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
			expected: []byte("val1"),
		},
	}

	for _, tc := range cases {
		cache.Add(tc.inputKey, tc.inputVal)
		time.Sleep(interval + time.Millisecond)
		_, ok := cache.Get(tc.inputKey)
		if ok {
			t.Errorf("%s should not be in the cache", tc.inputKey)
			continue
		}
	}
}

func TestReapFail(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)
	cases := []struct {
		inputKey string
		inputVal []byte
		expected []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
			expected: []byte("val1"),
		},
	}

	for _, tc := range cases {
		cache.Add(tc.inputKey, tc.inputVal)
		time.Sleep(interval / 2)
		_, ok := cache.Get(tc.inputKey)
		if !ok {
			t.Errorf("%s should be in the cache", tc.inputKey)
			continue
		}
	}
}
