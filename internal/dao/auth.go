package dao

import "gin-template/internal/model"

func (dao *Dao) GetAuth(appKey, appSecret string) (model.Auth, error) {
	auth := model.Auth{AppKey: appKey, AppSecret: appSecret}
	return auth.Get(dao.engine)
}
