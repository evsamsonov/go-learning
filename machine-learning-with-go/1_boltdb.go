package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	db, err := bolt.Open("machine-learning-with-go/bolt.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucket([]byte("MyBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %w", err)
		}
	
		if err = bucket.Put([]byte("key1"), []byte("value1")); err != nil {
			return fmt.Errorf("put: %w", err)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("MyBucket"))

		err := bucket.ForEach(func(k, v []byte) error {
			fmt.Printf("%s: %s\n", string(k), string(v))
			return nil
		})
		if err != nil {
			return fmt.Errorf("foreach: %w", err)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

}
