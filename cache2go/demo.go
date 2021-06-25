package cache2go

import (
	"fmt"
	"time"

	"github.com/muesli/cache2go"
	"github.com/spf13/cast"
)

// Keys & values in cache2go can be of arbitrary types, e.g. a struct.
type myStruct struct {
	title     string
	content interface{}
}

func Demo() {
	table := "xiao"
	key := "sai0556"
	// 新建 CacheTable
	cache := cache2go.Cache(table)

	val := myStruct{"what", "hello 13sai"}

	// 添加一个kv，过期时间 2s
	cache.Add(key, 2*time.Second, &val)

	// 获取value
	res, err := cache.Value(key)
	if err == nil {
		fmt.Println("Found value in cache:", res.Data().(*myStruct))
	} else {
		fmt.Println("Error retrieving value from cache:", err)
	}

	time.Sleep(3 * time.Second)
	res, err = cache.Value(key)
	if err != nil {
		fmt.Println("Item is not cached (anymore).")
	} else {
		fmt.Println("cache", res)
	}

	cache.Add(key, 0, &val)

	// 增加删除回调函数
	cache.SetAboutToDeleteItemCallback(func(e *cache2go.CacheItem) {
		fmt.Println("Deleting:", e.Key(), e.Data().(*myStruct).title, e.CreatedOn())
	})

	// 删除
	cache.Delete(key)

	// 清除所有item
	cache.Flush()
}

func SetDataLoader(table string) {
	cache := cache2go.Cache(table)
	// 在尝试访问不存在的key时将调用该回调
	cache.SetDataLoader(func(key interface{}, args ...interface{}) *cache2go.CacheItem {
		val := "this is " + key.(string) + ", 13sai"
		fmt.Println(key, "--", val)
		return cache2go.NewCacheItem(key, 0, val)
	})

	for i:=0; i < 3; i++ {
		_, err := cache.Value("generated_" + cast.ToString(i))
		if err != nil {
			fmt.Println("err", err)
		} else {
			fmt.Println(cast.ToString(i), "in cache")
		}
	}
	res, _ := cache.Value("generated_1")

	fmt.Println("generated_1-----", res.Data().(string))
}


func Callback(table, key string) {
	cache := cache2go.Cache(table)

	// 设置添加回调
	cache.SetAddedItemCallback(func(entry *cache2go.CacheItem){
		fmt.Println("callback 1:", entry.Key(), entry.Data(), entry.CreatedOn())
	})

	// 可以设置多个，但会被覆盖
	cache.SetAddedItemCallback(func(entry *cache2go.CacheItem){
		fmt.Println("callback 2:", entry.Key(), entry.Data(), entry.CreatedOn())
	})

	cache.SetAboutToDeleteItemCallback(func(entry *cache2go.CacheItem){
		fmt.Println("del 1:", entry.Key(), entry.Data(), entry.CreatedOn())
	})

	cache.SetAboutToDeleteItemCallback(func(entry *cache2go.CacheItem){
		fmt.Println("del 2:", entry.Key(), entry.Data(), entry.CreatedOn())
	})

	cache.Add(key, 0, "13sai")

	res, err := cache.Value(key)
	if err != nil {
		fmt.Println("err", err)
		return 
	}

	cache.Delete(key)

	fmt.Println(key, res.Data())

	cache.Delete(key)

	cache.RemoveAddedItemCallbacks()

	res = cache.Add(key + "2", time.Second, "sai0556")

	res.SetAboutToExpireCallback(func(i interface{}) {
		fmt.Println("expire", i.(string))
	})

	time.Sleep(2*time.Second)
}