package queue

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

var _ Queue = new(mysqlqueue)

const (
	Initial uint8 = iota
	Finish
)

type mysqlqueue struct {
	db *gorm.DB

	table string
}

func NewMysqlQueue(db *gorm.DB, table string) (*mysqlqueue, error) {
	if !db.Migrator().HasTable(table) {
		if err := db.Table(table).AutoMigrate(&MysqlMessage{}); err != nil {
			return nil, err
		}
	}
	return &mysqlqueue{db: db, table: table}, nil
}

func (mq *mysqlqueue) Offer(bytes []byte) error {
	return mq.db.Table(mq.table).Create(&MysqlMessage{
		UUID:      uuid.NewV4().String(),
		Content:   string(bytes),
		Status:    Initial,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}).Error
}

func (mq *mysqlqueue) Poll() ([]byte, error) {
	tx := mq.db.Begin()
	msg := MysqlMessage{}
	tx.Table(mq.table).Where("status = ?", Initial).First(&msg)
	tx.Table(mq.table).Model(&msg).Update("status", Finish)
	tx.Commit()
	return []byte(msg.Content), nil
}

func (mq *mysqlqueue) Size() int64 {
	var count int64
	mq.db.Table(mq.table).Model(&MysqlMessage{}).Where("status = ?", Initial).Count(&count)
	return count
}

type MysqlMessage struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;comment:主键"`
	UUID      string    `gorm:"type:varchar(256);comment:消息UUID"`
	Content   string    `gorm:"type:varchar(256);comment:消息内容"`
	Status    uint8     `gorm:"comment:消息状态;default:0"`
	CreatedAt time.Time `gorm:"type:timestamp;autoCreateTime;comment:消息创建时间"`
	UpdatedAt time.Time `gorm:"type:timestamp;autoUpdateTime;comment:消息更新时间"`
}
