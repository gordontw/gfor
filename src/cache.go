package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"time"
)

var (
	dbfile = "/tmp/hostcheck.db"
	bucket = "HostGroup"
)

func initBolt() {
	db, _ := bolt.Open(dbfile, 0600, nil)

	db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucketIfNotExists([]byte(bucket))
		err := b.Put([]byte("UpdateTime"), []byte(fmt.Sprintf("%d", int32(time.Now().Unix()))))
		return err
	})
	defer db.Close()
}

func updateHostStatus(group string, host string, value string) bool {
	initBolt()
	db, err := bolt.Open(dbfile, 0600, nil)
	if err != nil {
		return false
	}

	key := fmt.Sprintf("%s.%s", group, host)
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		err := b.Put([]byte(key), []byte(value))
		return err
	})
	defer db.Close()
	return true
}

func getHostStatus(group string, host string) string {
	initBolt()
	db, _ := bolt.Open(dbfile, 0600, nil)
	var value []byte

	key := fmt.Sprintf("%s.%s", group, host)
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		value = b.Get([]byte(key))
		return nil
	})
	defer db.Close()
	return string(value[:])
}
