package service

import "douyin_service/pkg/util"

// CommentActionRequest 评论操作请求
type CommentActionRequest struct {
	// 用户id
	UserId uint
	Token  string `form:"token" binding:"required"`
	// 视频id
	VideoId uint `form:"video_id" binding:"required"`
	// 1-发布评论，2-删除评论
	ActionType uint `form:"action_type" binding:"required"`
	// 用户填写的评论内容，在action_type=1的时候使用
	CommentText string `form:"comment_text"`
	// 用户填写的评论内容，在action_type=1的时候使用
	CommentId uint `form:"comment_id"`
}

// CommentActionResponse 评论操作响应
type CommentActionResponse struct {
	ResponseCommon
	CommentList []CommentInfo `json:"comment"`
}

// CommentListRequest 视频评论列表请求
type CommentListRequest struct {
	Token string `form:"token" binding:"required"`
	// 视频id
	VideoId uint `form:"video_id" binding:"required"`
}

// CommentListResponse 视频评论列表响应
type CommentListResponse struct {
	ResponseCommon
	CommentList []CommentInfo `json:"comment_list"`
}

// CommentInfo 单条评论信息
type CommentInfo struct {
	ID      uint     `json:"id"`
	User    UserInfo `json:"user"`
	Content string   `json:"content"`
	// 评论发布日期：mm-dd
	CreateDate string `json:"create_date"`
}

// UserInfo 用户信息
//type UserInfo struct {
//	ID uint `json:"id"`
//	Name string `json:"name"`
//	FollowCount int64 `json:"follow_count"`
//	FollowerCount int64 `json:"follower_count"`
//	IsFollow bool `json:"is_follow"`
//	Avatar string `json:"avatar"`
//	BackgroundImage string `json:"background_image"`
//	Signature string `json:"signature"`
//	TotalFavorited int64 `json:"total_favorited"`
//	WorkCount int64 `json:"work_count"`
//	FavoriteCount int64 `json:"favorite_count"`
//}

// CreateComment 创建评论
func (svc *Service) CreateComment(param *CommentActionRequest) (CommentActionResponse, error) {
	var cmtResp CommentActionResponse
	cmt, err := svc.dao.CreateComment(param.UserId, param.VideoId, param.CommentText)
	if err != nil {
		return cmtResp, err
	}
	user, err := svc.dao.GetUserById(uint(cmt.UserId))
	if err != nil {
		return CommentActionResponse{}, err
	}
	isFollow, err := svc.dao.IsFollow(user.ID, param.VideoId)
	if err != nil {
		return CommentActionResponse{}, err
	}
	cmtResp.CommentList = append(cmtResp.CommentList, CommentInfo{
		ID: cmt.ID,
		User: UserInfo{
			ID:              user.ID,
			Name:            user.UserName,
			FollowCount:     user.FollowCount,
			FollowerCount:   user.FollowerCount,
			IsFollow:        isFollow,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature:       user.Signature,
			TotalFavorited:  user.TotalFavorited,
			WorkCount:       user.WorkCount,
			FavoriteCount:   user.FavoriteCount,
		},
		Content:    cmt.Content,
		CreateDate: cmt.CreatedAt.Format("01-02"), // 这里要转化时间
	})
	return cmtResp, nil
}

// DeleteComment 删除评论
func (svc *Service) DeleteComment(param *CommentActionRequest) (error) {
	//var cmtResp CommentActionResponse
	err := svc.dao.DeleteComment(param.CommentId)
	if err != nil {
		return err
	}
	return  nil
	//user, err := svc.dao.GetUserById(uint(cmt.UserId))
	//if err != nil {
	//	return CommentActionResponse{}, err
	//}
	//isFollow, err := svc.dao.IsFollow(user.ID, param.VideoId)
	//if err != nil {
	//	return CommentActionResponse{}, err
	//}
	//cmtResp.CommentList = append(cmtResp.CommentList, CommentInfo{
	//	ID: cmt.ID,
	//	User: UserInfo{
	//		ID:              user.ID,
	//		Name:            user.UserName,
	//		FollowCount:     user.FollowCount,
	//		FollowerCount:   user.FollowerCount,
	//		IsFollow:        isFollow,
	//		Avatar:          user.Avatar,
	//		BackgroundImage: user.BackgroundImage,
	//		Signature:       user.Signature,
	//		TotalFavorited:  user.TotalFavorited,
	//		WorkCount:       user.WorkCount,
	//		FavoriteCount:   user.FavoriteCount,
	//	},
	//	Content:    cmt.Content,
	//	CreateDate: cmt.CreatedAt.Format("01-02"), // 这里要转化时间
	//})
	//return cmtResp, nil
}

// CommentList 视频评论列表
func (svc *Service) CommentList(param *CommentListRequest) (CommentListResponse, error) {
	var cmtListResp CommentListResponse
	cmts, err := svc.dao.GetCommentsByVideoId(param.VideoId)
	if err != nil {
		return cmtListResp, err
	}
	for _, c := range cmts {

		user, err := svc.dao.GetUserById(uint(c.UserId))
		if err != nil {
			return CommentListResponse{}, err
		}
		isFollow, err := svc.dao.IsFollow(user.ID, param.VideoId)
		if err != nil {
			return CommentListResponse{}, err
		}
		cmtListResp.CommentList = append(cmtListResp.CommentList, CommentInfo{
			ID: c.ID,
			User: UserInfo{
				ID:              user.ID,
				Name:            user.UserName,
				FollowCount:     user.FollowCount,
				FollowerCount:   user.FollowerCount,
				IsFollow:        isFollow,
				Avatar:          user.Avatar,
				BackgroundImage: user.BackgroundImage,
				Signature:       user.Signature,
				TotalFavorited:  user.TotalFavorited,
				WorkCount:       user.WorkCount,
				FavoriteCount:   user.FavoriteCount,
			},
			Content:    util.Filter(c.Content),
			CreateDate: c.CreatedAt.Format("01-02"), // 这里要转化时间
		})
	}
	return cmtListResp, nil
}
