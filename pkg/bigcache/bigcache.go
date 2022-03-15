package bigcache

import (
	"github.com/allegro/bigcache"
	"time"
)

type Cache struct {
	cache *bigcache.BigCache
}

func NewBigCache() *Cache {
	cache,err:=bigcache.NewBigCache(bigcache.DefaultConfig(30*time.Hour))
	if err!=nil{
		panic(err)
	}
	return &Cache{cache:cache}
}

func (c *Cache) Get(cookie string)([]byte,error){
	bytes,err:=c.cache.Get(cookie)
	return bytes,err
}

func (c *Cache) Set(key string,bytes []byte)error{
	return c.cache.Set(key,bytes)
}