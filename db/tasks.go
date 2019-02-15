package db

import (
    "time"

    "github.com/boltdb/bolt"
)


// Define our bucket
var taskBucket = []byte("tasks")
// Define our db
var db *bolt.DB

type Task struct {
    Key int
    Value string

}

func Init(dbPath string) error {
    var err error
    db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})

    if err != nil {
        return err
    }
    
    // Update the database in the context of a read-write managed transaction
    // If there is an error (the bucket exists) the transaction will get rolled back
    return db.Update(func(tx *bolt.Tx) error {
        // Create a new bucket if it doesn't exist,
        // Throw an error if it does
        _, err := tx.CreateBucketIfNotExists(taskBucket)
        return err
    })
}
