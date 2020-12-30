package db

import (
	"fmt"
	"log"
	"time"

	"github.com/marcelovicentegc/kontrolio-cli/src/config"
	bolt "go.etcd.io/bbolt"
)

const BUCKET_NAME = "KontrolioBucket"
const (
	PUNCHED_IN  = "PUNCHED_IN"
	PUNCHED_OUT = "PUNCHED_OUT"
)

func getBucket(transaction *bolt.Tx) *bolt.Bucket {
	bucket := transaction.Bucket([]byte(BUCKET_NAME))
	if bucket == nil {
		fmt.Println("Bucket doesn't exist, creating it...")
		newBucket, err := transaction.CreateBucketIfNotExists([]byte(BUCKET_NAME))

		if err != nil {
			log.Fatal(err)
		}

		return newBucket
	}

	return bucket
}

func getDb() *bolt.DB {
	dbPath := config.GetLocalDataStorePath()

	db, err := bolt.Open(dbPath, 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func SaveRecord() {
	db := getDb()

	err := db.Update(func(transaction *bolt.Tx) error {
		bucket := getBucket(transaction)

		_, value := bucket.Cursor().Last()

		recordType := string(value)

		if recordType == PUNCHED_IN {
			bucket.Put([]byte(time.Now().String()), []byte(PUNCHED_OUT))
		} else if recordType == PUNCHED_OUT {
			bucket.Put([]byte(time.Now().String()), []byte(PUNCHED_IN))
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}

func cursorIterator(key []byte, value []byte, iterator int) (k []byte, v []byte, i int) {
	return key, value, iterator
}
