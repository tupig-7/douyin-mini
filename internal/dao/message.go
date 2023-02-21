package dao

import (
	"douyin_service/internal/model"
	"douyin_service/pkg/errcode"
)

// CreateMessage 创建一条新消息
func (d *Dao) CreateMessage(toUserId, fromUserId uint, content string) (uint, error) {
	msg := model.Message{
		ToUserId:   toUserId,
		FromUserId: fromUserId,
		Content:    content,
	}
	err := msg.Create(d.engine)
	if err != nil {
		return errcode.ErrorUserID, err
	}
	return msg.ID, nil
}

// GetMsgByToUserId 查询对方用户id to_user_id为id的消息记录
func (d *Dao) GetMsgByToUserId(to_user_id, from_user_id uint) ([]model.Message, error) {
	var msg model.Message
	msg.ToUserId = to_user_id
	msg.FromUserId = from_user_id
	msgs, err := msg.ListByToUserId(d.engine)
	if err != nil {
		return nil, err
	}
	return msgs, nil
}