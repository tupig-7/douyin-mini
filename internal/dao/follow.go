package dao

import "douyin_service/internal/model"

func (d *Dao) IsFollow(follower, followed uint) (flag bool, err error) {
	follow := model.Follow{
		FollowedId: followed,
		FollowerId: follower,
	}
	flag, err = follow.IsExist(d.engine)
	return
}
