package service

type FollowActionRequest struct {
	UserId uint
	Token      string `form:"token" binding:"required"`
	ToUserId   uint   `form:"to_user_id" binding:"required"`
	ActionType int64  `form:"action_type" binding:"required"`
}

type FollowActionResponse struct {
	ResponseCommon
}

type FollowListRequest struct {
	UserId uint   `form:"user_id"  binding:"required"`
	Token  string `form:"token" binding:"required"`
}

type FollowListResponse struct {
	ResponseCommon
	UserList []UserInfo `json:"user_list" binding:"required"`
}

type FollowerListRequest struct {
	UserId uint   `form:"user_id"  binding:"required"`
	Token  string `form:"token" binding:"required"`
}

type FollowerListResponse struct {
	ResponseCommon
	UserList []UserInfo `json:"user_list" binding:"required"`
}

type FriendListRequest struct {
	UserId uint   `form:"user_id"  binding:"required"`
	Token  string `form:"token" binding:"required"`
}

type FriendListResponse struct {
	ResponseCommon
	UserList []UserInfo `json:"user_list" binding:"required"`
}

// IsFollowRequest 判断A是否关注B
type IsFollowRequest struct {
	A uint
	B uint
}

// CreateFollow 关注操作
func (svc *Service) CreateFollow(param *FollowActionRequest) error  {
	_, err := svc.dao.CreateFollow(param.UserId, param.ToUserId)
	if err != nil {
		return err
	}
	return nil
}

// CancelFollow 取消关注
func (svc *Service) CancelFollow(param *FollowActionRequest) error {
	_, err := svc.dao.CancelFollow(param.UserId, param.ToUserId)
	if err != nil {
		return  err
	}
	return nil
}
func (svc *Service) FollowList(param *FollowListRequest) (res FollowListResponse, err error) {
	follows, err := svc.dao.FollowList(param.UserId)
	if err != nil {
		return
	}
	for i := range follows {
		f := follows[i]
		id := f.FollowedId
		var userInfo UserInfo
		userM, UserErr := svc.dao.GetUserById(uint(id))
		if UserErr != nil {
			err = UserErr
			return
		}
		userInfo.ID = userM.ID
		userInfo.FollowCount = userM.FollowerCount
		userInfo.FollowerCount = userM.FollowerCount
		userInfo.Name = userM.UserName
		userInfo.IsFollow = true
		res.UserList = append(res.UserList, userInfo)
	}
	res.StatusCode = 0
	res.StatusMsg = "success"
	return
}

func (svc *Service) FollowerList(param *FollowerListRequest) (res FollowerListResponse, err error) {
	follows, err := svc.dao.FollowerList(param.UserId)
	if err != nil {
		return
	}
	for i := range follows {
		f := follows[i]
		id := f.FollowerId
		var userInfo UserInfo
		userM, UserErr := svc.dao.GetUserById(uint(id))
		if UserErr != nil {
			err = UserErr
			return
		}
		userInfo.ID = userM.ID
		userInfo.FollowCount = userM.FollowerCount
		userInfo.FollowerCount = userM.FollowerCount
		userInfo.Name = userM.UserName
		res.UserList = append(res.UserList, userInfo)
	}
	res.StatusCode = 0
	res.StatusMsg = "success"
	return
}


func (svc *Service) FriendList(param *FriendListRequest) (res FriendListResponse, err error) {
	follows, err := svc.dao.FriendList(param.UserId)
	if err != nil {
		return
	}
	for i := range follows {
		f := follows[i]
		id := f.FollowerId
		var userInfo UserInfo
		userM, UserErr := svc.dao.GetUserById(uint(id))
		if UserErr != nil {
			err = UserErr
			return
		}
		userInfo.ID = userM.ID
		userInfo.FollowCount = userM.FollowerCount
		userInfo.FollowerCount = userM.FollowerCount
		userInfo.Name = userM.UserName
		res.UserList = append(res.UserList, userInfo)
	}
	res.StatusCode = 0
	res.StatusMsg = "success"
	return
}
