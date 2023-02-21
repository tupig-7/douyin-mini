package service

// CommentListRequest 视频评论列表请求
type CommentListRequest struct {
	Token  string `form:"token" binding:"required"`
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
	ID uint `json:"id"`
	User UserInfo `json:"user"`
	Content string `json:"content"`
	// 评论发布日期：mm-dd
	CreateDate string `json:"create_date"`
}

// UserInfo 用户信息
type UserInfo struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	FollowCount int64 `json:"follow_count"`
	FollowerCount int64 `json:"follower_count"`
	IsFollow bool `json:"is_follow"`
	Avatar string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature string `json:"signature"`
	TotalFavorited int64 `json:"total_favorited"`
	WorkCount int64 `json:"work_count"`
	FavoriteCount int64 `json:"favorite_count"`
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
			ID:         c.ID,
			User:       UserInfo{
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
			Content:    c.Content,
			CreateDate: c.CreatedAt.Format("01-02"), // 这里要转化时间
		})
	}
	return cmtListResp, nil
}