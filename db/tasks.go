package db

import (
	"encoding/binary"
	"errors"

	//"log"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")
var db *bolt.DB

// Task ...
type Task struct {
	Key   int
	Value string
}

//Init ialize Datbase ...
func Init(dbPath string) (*bolt.DB, error) {
	//fmt.Println(dbPath)
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}

	return db, db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

//CreateTask ...
func CreateTask(str string) (int, error) {
	var id int

	err := db.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket(taskBucket)
		if b == nil {
			return errors.New("Bucket not found")
		}
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)
		return b.Put(key, []byte(str))

	})

	if err != nil {
		return -1, err
	}
	return id, nil
}

//AllTasks ...
func AllTasks() ([]Task, error) {
	var tasks []Task

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})

	return tasks, err
}

//DeleteTask ...
func DeleteTask(k int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		if b == nil {
			return errors.New("Bucket not found")
		}
		return b.Delete(itob(k))
	})
}

func itob(z int) (b []byte) {
	b = make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(z))
	return
}

func btoi(s []byte) int {
	return int(binary.BigEndian.Uint64(s))
}
