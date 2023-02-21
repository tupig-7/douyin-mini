package service

import (
	"douyin_service/internal/model"
)

type FavoriteActionRequest struct {
	Token string `form:"token" binding:"required"`
	// 用户Id
	UserId uint
	// 视频id
	VideoId uint `form:"video_id" binding:"required"`
	// 1-点赞，2-取消点赞
	ActionType uint `form:"action_type" binding:"required"`
}

type FavoriteActionResponse struct {
	ResponseCommon
}

// FavoriteListRequest 视频评论列表请求
type FavoriteListRequest struct {
	Token string `form:"token" binding:"required"`
	// 视频id
	UserId uint `form:"user_id" binding:"required"`
}

type FavoriteListResponse struct {
	ResponseCommon
	VideoList []VideoInfo `json:"video_list"`
}

// CreateFavorite 点赞操作
func (svc *Service) CreateFavorite(param *FavoriteActionRequest) error {
	_, err := svc.dao.CreateFavorite(param.UserId, param.VideoId)
	if err != nil {
		return err
	}
	return nil
}

// CancelFavorite 取消点赞
func (svc *Service) CancelFavorite(param *FavoriteActionRequest) error {
	err := svc.dao.CancelFavorite(param.UserId, param.VideoId)
	if err != nil {
		return err
	}
	return nil
}

// FavoriteList 查询视频列表
func (svc *Service) FavoriteList(param *FavoriteListRequest) (FavoriteListResponse, error) {
	var fvtResp FavoriteListResponse
	videoIds, err := svc.dao.GetFavoritesByUserId(param.UserId)
	var videos []VideoInfo
	if err != nil {
		return fvtResp, err
	}
	for _, vid := range videoIds {
		video := model.Video{}
		video, err := svc.dao.QueryVideoInfoById(vid)
		if err != nil {
			return fvtResp, err
		}
		user, err := svc.dao.GetUserById(video.AuthorId)
		if err != nil {
			return fvtResp, err
		}
		isFollow, err := svc.dao.IsFollow(video.AuthorId, user.ID)
		videoInfo := VideoInfo{
			Id: video.ID,
			Author: UserInfo{
				ID:              user.ID,
				Name:            user.UserName,
				FollowCount:     user.FollowCount,
				FollowerCount:   user.FollowerCount,
				IsFollow:        isFollow,
				Avatar:          user.Avatar,
				Signature:       user.Signature,
				BackgroundImage: user.BackgroundImage,
				WorkCount:       user.WorkCount,
				LoginIP:         user.LoginIP,
				TotalFavorited:  user.TotalFavorited,
				FavoriteCount:   user.FavoriteCount,
			},
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    true,
			Title:         video.Title,
		}
		videos = append(videos, videoInfo)
	}
	fvtResp.VideoList = videos
	return fvtResp, nil
}

// IsFavor 查询是否点赞的功能
// 由于用户对哪些视频点赞使用bitmap存储到redis中，因此直接在redis查询。
func (svc *Service) IsFavor(userId uint, videoId uint) (bool, error) {
	return svc.dao.IsFavor(userId, videoId)
}

func (svc *Service) QueryFavorCnt(videoId uint) (int64, error) {
	return svc.dao.QueryFavorCnt(videoId)
}
