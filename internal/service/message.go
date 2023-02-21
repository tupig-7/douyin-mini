package service

import "douyin_service/pkg/util"

// ChatRequest 聊天请求
type ChatRequest struct {
	Token  string `form:"token" binding:"required"`
	// 对方用户id
	ToUserId uint `form:"to_user_id" binding:"required"`
	// 发送用户id
	FromUserId uint
	// 1-发送消息
	ActionType int `form:"action_type" binding:"required"`
	// 消息内容
	Content string `form:"content" binding:"required"`
}

// ChatResponse 聊天响应
type ChatResponse struct {
	ResponseCommon
}

// MessagesRequest 消息记录请求
type MessagesRequest struct {
	// 发送方ID
	FromUserId uint
	Token  string `form:"token" binding:"required"`
	// 对方用户id
	ToUserId uint `form:"to_user_id" binding:"required"`
}

// MessagesResponse 聊天记录响应
type MessagesResponse struct {
	ResponseCommon
	MessageList []message `form:"message_list" json:"message_list"`
}

type message struct {
	ID uint `json:"id"`
	ToUserId uint `json:"to_user_id"`
	FromUserId uint `json:"from_user_id"`
	Content string `json:"content"`
	// 消息发送时间 yyyy-MM-dd HH:MM:ss
	CreateTime int64 `json:"create_time"`
}

// MessageAction 发送消息
func (svc *Service) MessageAction(param *ChatRequest) (ChatResponse, error) {
	var chatResp ChatResponse
	_, err := svc.dao.CreateMessage(param.ToUserId, param.FromUserId, param.Content)
	if err != nil {
		return chatResp, err
	}
	return chatResp, nil
}

// MessageChat 聊天记录
func (svc *Service) MessageChat(param *MessagesRequest) (MessagesResponse, error) {
	var msgResp MessagesResponse
	msgs, err := svc.dao.GetMsgByToUserId(param.ToUserId, param.FromUserId)
	if err != nil {
		return msgResp, err
	}
	for _, m := range msgs {
		msg := message{
			ID:         m.ID,
			ToUserId:   m.ToUserId,
			FromUserId: m.FromUserId,
			Content:    util.Filter(m.Content),
			CreateTime: m.CreatedAt,
		}
		msgResp.MessageList = append(msgResp.MessageList, msg)
	}
	return msgResp, nil
}

