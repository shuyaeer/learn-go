package adapters

import (
	"testing"
)

func TestMemoryCacheAdapter_Get(t *testing.T) {
	adapter := NewMemoryCacheAdapter()

	adapter.Set("key", "value")

	if adapter.Get("key") != "value" {
		t.Error("Getメソッドのテストに失敗しました")
	}
}