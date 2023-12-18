package adapters

import (
	"testing"
)

func TestRedisAdapter_Get(t *testing.T) {
	adapter := NewRedisAdapter()

	adapter.Set("key", "value")
	want := "value"
	got := adapter.Get("key")

	if got != want {
		t.Error("RedisAdapter.Get failed to get the value")
	}
}