package interfaces

type Cache interface {
	Get(key string) string
	Set(key string, value string)
	Delete(key string)
}

func GetData(cache Cache, key string) string {
	return cache.Get(key)
}

func SetData(cache Cache, key string, value string) {
	cache.Set(key, value)
}

func DeleteData(cache Cache, key string) {
	cache.Delete(key)
}
