package interfaces

import (
	"testing"
)

func TestCache_Get(t *testing.T) {
	// Create a mock implementation of the Cache interface
	cache := &mockCache{
		data: map[string]string{
			"key": "value",
		},
	}

	// Call the Get method
	result := cache.Get("key")

	// Check if the result matches the expected value
	if result != "value" {
		t.Errorf("Get method returned incorrect value. Expected: %s, Got: %s", "value", result)
	}
}

func TestCache_Set(t *testing.T) {
	// Create a mock implementation of the Cache interface
	cache := &mockCache{
		data: make(map[string]string),
	}

	// Call the Set method
	cache.Set("key", "value")

	// Check if the value was set correctly
	if cache.data["key"] != "value" {
		t.Errorf("Set method did not set the value correctly. Expected: %s, Got: %s", "value", cache.data["key"])
	}
}

// Mock implementation of the Cache interface for testing
type mockCache struct {
	data map[string]string
}

func (m *mockCache) Get(key string) string {
	return m.data[key]
}

func (m *mockCache) Set(key string, value string) {
	m.data[key] = value
}
