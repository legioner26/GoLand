package result

import (
	"sync"
	"time"
)

type CacheResultSetT struct {
	Data ResultSetT
	Time time.Time
	sync.Mutex
}

var cache CacheResultSetT

func GetFromCache() (ResultSetT, bool) {
	cache.Lock()
	cacheData := cache.Data
	cacheTime := cache.Time
	cache.Unlock()

	if cacheTime.Add(30 * time.Second).After(time.Now()) {
		return cacheData, true
	}

	return cacheData, false
}

func SetToCache(data ResultSetT, time time.Time) {
	cache.Lock()
	cache.Data = data
	cache.Time = time
	cache.Unlock()
}
