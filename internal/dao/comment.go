package dao

import (
	"douyin_service/internal/model"
	"gorm.io/gorm"
	"time"
)

// CreateComment 创建新评论
func (d *Dao) CreateComment(userId, videoId uint, content string) (model.Comment, error) {
	cmt := model.Comment{
		Model:   gorm.Model{
			CreatedAt: time.Now(),
		},
		UserId:  userId,
		VideoId: videoId,
		Content: content,
	}
	err := cmt.Create(d.engine)
	if err != nil {
		return cmt, err
	}
	return cmt, nil
}

func (d *Dao) DeleteComment(commentId uint) error {
	cmt := model.Comment{
		Model:   gorm.Model{ID: commentId},
		UserId:  0,
		VideoId: 0,
		Content: "",
	}
	err := cmt.Delete(d.engine)
	return err
}

// GetCommentsByVideoId 根据视频id获取视频评论，按发布时间降序
func (d *Dao) GetCommentsByVideoId(videoId uint) ([]*model.Comment, error) {
	c := model.Comment{VideoId: videoId }
	comments, err := c.List(d.engine)
	if err != nil {
		return nil, err
	}
	return comments, nil
}