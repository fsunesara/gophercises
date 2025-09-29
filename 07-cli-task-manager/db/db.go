package db

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"time"

	bolt "go.etcd.io/bbolt"
)

var db *bolt.DB

type Task struct {
	Id             uint64
	Text           string
	CompletionTime *time.Time
}

func (t *Task) ToJSONString() ([]byte, error) {
	s, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func InitDB() {
	var err error
	db, err = bolt.Open("tasks.db", 0600, nil)
	if err != nil {
		panic(err)
	}
	db.Logger().Debug("db initialized")
}

func uint64ToBytes(i uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, i)
	return b
}

func AddTask(t Task) {
	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("Tasks"))
		if err != nil {
			return err
		}

		id, _ := b.NextSequence()

		t.Id = id
		str, err := t.ToJSONString()
		if err != nil {
			return err
		}

		return b.Put(uint64ToBytes(id), str)
	})
}

func CompleteTask(taskId uint64) {
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tasks"))
		taskStr := b.Get(uint64ToBytes(taskId))
		var task Task
		err := json.Unmarshal(taskStr, &task)
		if err != nil {
			return err
		}

		now := time.Now()
		task.CompletionTime = &now
		taskStr, err = task.ToJSONString()

		if err != nil {
			return err
		}

		return b.Put(uint64ToBytes(taskId), taskStr)
	})
}

func ListTasks() []Task {
	tasks := make([]Task, 0)
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tasks"))
		if b == nil {
			return fmt.Errorf("no tasks found")
		}
		c := b.Cursor()

		for id, txt := c.First(); id != nil; id, txt = c.Next() {
			var task Task
			err := json.Unmarshal(txt, &task)
			if err != nil {
				return err
			}
			if task.CompletionTime == nil {
				tasks = append(tasks, task)
			}
		}
		return nil
	})
	return tasks
}

func CloseDB() {
	db.Close()
}
