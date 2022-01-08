package service

import "errors"

type AuthRequest struct {
	Appkey    string `form:"app_key" binding:"required"`
	Appsecret string `form:"app_secret" binding:"required"`
}

func (svg *Service) CheckAuth(params *AuthRequest) error {
	auth, err := svg.dao.GetAuth(params.Appkey, params.Appsecret)
	if err != nil {
		return err
	}
	if auth.ID > 0 {
		return nil
	}
	return errors.New("auth info does not exist")
}
