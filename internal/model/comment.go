package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserId uint `json:"user_id"`
	VideoId uint `json:"video_id"`
	Content string `json:"content"`
}

func (c *Comment) TableName() string {
	return "douyin_comment"
}

func (c *Comment) Create(db *gorm.DB) error {
	return db.Create(&c).Error
}

func (c *Comment) Delete(db *gorm.DB) error {
	return db.Where("id = ?", c.ID).Delete(&c).Error
}

func (c *Comment) List(db *gorm.DB) ([]*Comment, error) {
	var comments []*Comment
	var err error
	if err = db.Where("video_id = ?", c.VideoId).Order("created_at desc").Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}