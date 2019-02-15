package main

import (
    "time"

    "github.com/boltdb/bolt"
)

func main() {
    // Open a connection to the db (my.db) with file mode 0600 and an options struct
    // Specifying that we'd like to timeout our attempted connection if it takes longer
    // than one second to connect.
    db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})

    // Check if there was an error loading the database
    if err != nil {
        panic(err)
    }

    // Doesn't execute until the surrounding function has returned
    // Defers the databse closing until everything else is done
    defer db.Close()
}
