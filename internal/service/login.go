package service

type LoginRequest struct {
	UserName string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	LoginIP string `form:"login_ip"`
}

type LoginResponse struct {
	ResponseCommon
	UserID uint   `json:"user_id"`
	Token  string `json:"token"`
}

type RegisterRequest struct {
	UserName string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	LoginIP string  `form:"login_ip"`
}

type RegisterResponse struct {
	ResponseCommon
	UserID uint   `json:"user_id"`
	Token  string `json:"token"`
}

func (svc *Service) Login(param *LoginRequest) (uint, bool, error) {
	return svc.dao.CheckUser(param.UserName, param.Password, param.LoginIP)
}