package queue

import (
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestMysqlQueue(t *testing.T) {
	dsn := "user:pass@tcp(ip:port)/database?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	q, err := NewMysqlQueue(db, "msg_17")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	q.Offer([]byte("shit"))
	t.Log(q.Size())
	msg, err := q.Poll()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(string(msg))
	t.Log(q.Size())
	for {
		t.Log(q.Size())
		if q.Size() <= 0 {
			break
		}
		msg, err := q.Poll()
		if err != nil {
			t.Log(err)
			t.FailNow()
		}
		t.Log(string(msg))
	}
}
