package model

import (
	"fmt"
	"gorm.io/gorm"
)

type Message struct {
	*Model
	ToUserId uint `json:"to_user_id"`
	FromUserId uint `json:"from_user_id"`
	Content string `json:"content"`
}

func (m Message) TableName() string {
	return "douyin_message"
}

func (m *Message) Create(db *gorm.DB) error {
	return db.Create(&m).Error
}

func (m *Message) Update(db *gorm.DB) error {
	return db.Model(&Message{}).Where("id = ?", m.ID).Error
}

func (m *Message) ListByToUserId(db *gorm.DB) ([]Message, error) {
	var msgs []Message
	fmt.Println("111", m.ToUserId)
	var err error
	if err = db.Where("to_user_id = ?", m.ToUserId).Find(&msgs).Error; err != nil {
		return nil, err
	}
	return msgs, nil
}