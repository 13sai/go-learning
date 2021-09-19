package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

func main() {
	// 文件不存在，会新建文件
	db, err := bolt.Open("sai.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 创建一个bucket，可以理解成一个table
	db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucketIfNotExists([]byte("sai0556"))
		// 新增
		b.Put([]byte("a"), []byte("11"))
		b.Put([]byte("b"), []byte("22"))
		b.Put([]byte("c"), []byte("33"))
		b.Put([]byte("d"), []byte("44"))
		b.Put([]byte("e"), []byte("55"))
		b.Put([]byte("f"), []byte("66"))

		// 取
		v := b.Get([]byte("a"))
		fmt.Printf("The a is: %s\n", v)

		// 删除
		b.Delete([]byte("a"))
		return nil
	})

	// 游标遍历
	/*
	   First()  Move to the first key.
	   Last()   Move to the last key.
	   Seek()   Move to a specific key.
	   Next()   Move to the next key.
	   Prev()   Move to the previous key
	*/
	db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucketIfNotExists([]byte("sai0556"))

		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("Cursor key=%s, value=%s\n", k, v)
		}

		// ForEach 遍历
		b.ForEach(func(k, v []byte) error {
			fmt.Printf("ForEach key=%s, value=%s\n", k, v)
			return nil
		})
		return nil
	})

	// 删除bucket
	db.Update(func(tx *bolt.Tx) error {
		tx.DeleteBucket([]byte("sai0556"))
		return nil
	})
}
