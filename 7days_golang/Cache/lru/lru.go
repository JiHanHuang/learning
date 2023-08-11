package lru

import (
	"errors"
	"sync"
	"time"

	lru "github.com/hashicorp/golang-lru"
)

type LruCacheLoadFunc func(key interface{}) (interface{}, bool, error)

type Opt struct {
	Name string
	//expire of item's data
	Exp time.Duration
	//Number of data in lru.Cache
	CacheSize    int
	LoadFromSave LruCacheLoadFunc
}

type LruCache struct {
	Opt
	items *lru.Cache
}

type lruCacheData struct {
	//thread-safe fo lruCacheData
	mu             sync.RWMutex
	createOrUpdate time.Time
	Value          interface{}
}

const DefaultLruCacheSize = 32
const DefaultLruCacheExpire = 3 * time.Second

func NewCache(opt Opt) (*LruCache, error) {
	if opt.LoadFromSave == nil {
		return nil, errors.New("Please input LoadFromSave function in Opt")
	}
	if opt.Exp < time.Millisecond {
		opt.Exp = DefaultLruCacheExpire
	}
	if opt.CacheSize <= 0 {
		opt.CacheSize = DefaultLruCacheSize
	}
	lc := &LruCache{
		Opt: opt,
	}
	lc.items, _ = lru.New(opt.CacheSize)
	return lc, nil
}

//Get value should be a pointer
func (lc *LruCache) Get(key interface{}) (interface{}, bool, error) {
	//lc.items.Get(key) will be used Lock instead of RLock.
	v, ok := lc.items.Peek(key)
	if !ok {
		return lc.loadFromDB(key)
	}
	cacheData := v.(*lruCacheData)
	cacheData.mu.RLock()
	active := cacheData.createOrUpdate.Add(lc.Exp).After(time.Now())
	cacheData.mu.RUnlock()
	if !active {
		lc.del(key)
		return lc.loadFromDB(key)
	}
	return cacheData.Value, true, nil
}

//loadFromSave value should be a pointer
func (lc *LruCache) loadFromDB(key interface{}) (value interface{}, ok bool, err error) {
	value, ok, err = lc.LoadFromSave(key)
	if ok {
		lc.set(key, value)
	}
	return value, ok, err
}

func (lc *LruCache) set(key, value interface{}) {
	cacheData := &lruCacheData{
		createOrUpdate: time.Now(),
		Value:          value,
	}
	lc.items.Add(key, cacheData)
	return
}

func (lc *LruCache) del(key interface{}) {
	lc.items.Remove(key)
	return
}
