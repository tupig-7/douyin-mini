package dao

import (
	"douyin_service/internal/model"

	"time"

	"gorm.io/gorm"
)

// 返回userID发布的所有信息
func (d *Dao) ListVideoByUserId(userId uint) ([]model.Video, error) {
	video := model.Video{AuthorId: userId}
	return video.ListVideoByUserId(d.engine)
}

// 上传视频
func (d *Dao) PublishVideo(authorId uint, playUrl, coverUrl, title string) error {
	now := time.Now()
	video := model.Video{
		Model: gorm.Model{
			CreatedAt: now,
		},
		AuthorId:      authorId,
		PlayUrl:       playUrl,
		CoverUrl:      coverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         title,
		PublishDate:   now,
	}
	user, _ := d.GetUserById(authorId)
	user.WorkCount += 1
	_ = user.UpdateWorkCnt(d.engine)
	return video.Create(d.engine)
}

// 通过视频id返回视频
func (d *Dao) QueryVideoById(videoId uint) (model.Video, error) {
	var video model.Video
	video, err := video.QueryVideoById(videoId, d.engine)
	if err != nil {
		return video, err
	}
	return video, nil
}

// 通过一组视频id返回一组视频
func (d *Dao) QueryBatchVideoById(favorList []uint) ([]model.Video, error) {
	var video model.Video
	videos, err := video.QueryBatchVdieoById(favorList, d.engine)
	if err != nil {
		return nil, err
	}
	return videos, nil
}

// 更新video的点赞量
func (d *Dao) UpdateFavoriteCnt(video model.Video) error {
	return video.UpdateFavoriteCnt(d.engine)
}

// 更新video的comment_count
func (d *Dao) UpdateCommentCnt(video model.Video) error {
	return video.UpdateCommentCnt(d.engine)
}

// 根据Id查询点赞数量
func (d *Dao) QueryFavorCntById(videoId uint) (int64, error) {
	var video model.Video
	video.ID = videoId
	return video.QueryFavorCntById(d.engine)
}

func (d *Dao) QueryAuthorIdByVideoId(videoId uint) (uint, error) {
	var video model.Video
	video.ID = videoId
	return video.QueryAuthorIdByVideoId(d.engine)
}
