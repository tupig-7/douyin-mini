package dao

import (
	"douyin_service/internal/model"
	"time"
)

// CreateFavorite 创建点赞
func (d *Dao) CreateFavorite(userId, videoId uint) (model.Favorite, error) {
	fvt := model.Favorite{
		Model: &model.Model{
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		},
		UserId:  userId,
		VideoId: videoId,
	}
	err := fvt.Create(d.engine)
	if err != nil {
		return fvt, err
	}
	video, _ := d.QueryVideoInfoById(videoId)
	user, _ := d.GetUserById(userId)
	video.FavoriteCount += 1
	err = d.UpdateFavoriteCnt(video)
	user.FavoriteCount += 1
	err = user.UpdateFavoriteCnt(d.engine)
	return fvt, nil
}

// CancelFavorite 取消点赞
func (d *Dao) CancelFavorite(userId, videoId uint) error {
	fvt := model.Favorite{
		Model: &model.Model{
			CreatedAt: time.Now().UnixMilli(),
			UpdatedAt: time.Now().UnixMilli(),
		},
		UserId:  userId,
		VideoId: videoId,
	}
	err := fvt.Delete(d.engine)
	if err != nil {
		return err
	}
	video, _ := d.QueryVideoInfoById(videoId)
	user, _ := d.GetUserById(userId)
	video.FavoriteCount -= 1
	err = d.UpdateFavoriteCnt(video)
	user.FavoriteCount -= 1
	err = user.UpdateFavoriteCnt(d.engine)
	return nil
}

func (d *Dao) QueryVideoInfoById(videoId uint) (model.Video, error) {
	var video model.Video
	v, err := video.QueryVideoById(videoId, d.engine)
	if err != nil {
		return video, err
	}
	return v, nil
}

// GetFavoritesByUserId 获取所有点赞的视频id
func (d *Dao) GetFavoritesByUserId(userId uint) ([]uint, error) {
	f := model.Favorite{UserId: userId}
	videoIds, err := f.QueryFavoriteByUserId(d.engine) // 获取点赞的所有视频id
	if err != nil {
		return nil, err
	}

	return videoIds, nil
}

func (d *Dao) QueryFavorCnt(videoId uint) (int64, error) {
	f := model.Favorite{VideoId: videoId}
	return f.QueryFavoritedCnt(d.engine)
}

func (d *Dao) IsFavor(userId, videoId uint) (bool, error) {
	f := model.Favorite{UserId: userId, VideoId: videoId}
	return f.IsFavor(d.engine)
}
