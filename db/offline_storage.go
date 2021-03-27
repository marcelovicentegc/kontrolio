package db

import (
	"fmt"
	"log"
	"time"

	"github.com/marcelovicentegc/kontrolio-cli/config"
	"github.com/marcelovicentegc/kontrolio-cli/messages"
	"github.com/marcelovicentegc/kontrolio-cli/utils"

	bolt "go.etcd.io/bbolt"
)

const BucketName = "KontrolioBucket"

type recordTypeRegistry struct {
	In  string
	Out string
}

func newRecordTypeRegistry() *recordTypeRegistry {
	return &recordTypeRegistry{
		In:  "IN",
		Out: "OUT",
	}
}

var RecordTypeRegistry = newRecordTypeRegistry()

func getBucket(transaction *bolt.Tx) *bolt.Bucket {
	bucket := transaction.Bucket([]byte(BucketName))
	if bucket == nil {
		fmt.Println(messages.CreatingBucket)
		newBucket, err := transaction.CreateBucketIfNotExists([]byte(BucketName))

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

		if record.Type == RecordTypeRegistry.In {
			bucket.Put(key, []byte(RecordTypeRegistry.Out))
		} else {
			bucket.Put(key, []byte(RecordTypeRegistry.In))
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if record.Type == RecordTypeRegistry.In {
		return RecordTypeRegistry.Out
	} else {
		return RecordTypeRegistry.In
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
		log.Fatal(err)
	}

	return records
}
