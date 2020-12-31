package db

import (
	"fmt"
	"log"
	"time"

	"github.com/marcelovicentegc/kontrolio-cli/src/config"
	"github.com/marcelovicentegc/kontrolio-cli/src/utils"

	bolt "go.etcd.io/bbolt"
)

const BUCKET_NAME = "KontrolioBucket"
const (
	PUNCHED_IN  = "punched in"
	PUNCHED_OUT = "punched out"
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

func SaveOfflineRecord() string {
	db := getDb()
	var record utils.Record

	err := db.Update(func(transaction *bolt.Tx) error {
		bucket := getBucket(transaction)

		_, recordType := bucket.Cursor().Last()
		record = utils.Record{Time: time.Now().In(time.Local), Type: string(recordType)}
		key, _ := utils.ByteSerializeOfflineRecord(record)

		if record.Type == PUNCHED_IN {
			bucket.Put(key, []byte(PUNCHED_OUT))
		} else {
			bucket.Put(key, []byte(PUNCHED_IN))
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if record.Type == PUNCHED_IN {
		return PUNCHED_OUT
	} else {
		return PUNCHED_IN
	}
}

func GetOfflineRecords() []string {
	db := getDb()
	var records []string

	err := db.View(func(transaction *bolt.Tx) error {
		bucket := getBucket(transaction)

		cursor := bucket.Cursor()

		for key, value := cursor.First(); key != nil; key, value = cursor.Next() {
			records = append(records, utils.SerializeOfflineRecord(key, value))
		}

		return nil
	})

	defer db.Close()

	if err != nil {
		log.Fatalf("failure : %s\n", err)
	}

	return records
}
