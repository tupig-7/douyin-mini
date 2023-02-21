package model

import (
	"gorm.io/gorm"
)

// 点赞信息数据库
type Favorite struct {
	*Model
	UserId  uint `json:"user_id"`
	VideoId uint `json:"video_id"`
}

// 返回数据库名称
func (f Favorite) TableName() string {
	return "douyin_favorite"
}

// 插入点赞信息数据
func (f Favorite) Create(db *gorm.DB) error {
	return db.Create(&f).Error
}

// 删除具体点赞信息
func (f Favorite) Delete(db *gorm.DB) error {
	return db.Model(&f).Where("user_id = ? AND video_id = ?", f.UserId, f.VideoId).Delete(&f).Error
}

// IsFavor userId是否给videoId点赞
func (f Favorite) IsFavor(db *gorm.DB) (bool, error) {
	var count int64
	err := db.Model(&f).Where("user_id = ? AND video_id = ?", f.UserId, f.VideoId).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, err
}

// QueryFavoritedCnt  查询视频获赞数量
func (f Favorite) QueryFavoritedCnt(db *gorm.DB) (int64, error) {
	var count int64
	var err error
	if err = db.Where("video_id = ?", f.VideoId).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, err
}

// QueryFavoriteByUserId 查询userId点赞的videoId
func (f Favorite) QueryFavoriteByUserId(db *gorm.DB) ([]uint, error) {
	var favorList []uint
	err := db.Model(&Favorite{}).Select("video_id").Where("user_id = ?", f.UserId).Find(&favorList).Error
	if err != nil {
		return nil, err
	}
	return favorList, err
}
