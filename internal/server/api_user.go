package server

import (
	"shuaoyoupin/internal/kit/db"
	"shuaoyoupin/internal/model"
)

type ApiUserer interface {
	ApiUserCreate(m *model.ApiUser) error
	//ApiUserUpdate(m *model.ApiUser) error
	ApiUserGet(where map[string]interface{}) (*model.ApiUser, error)
	//ApiUserList(m *model.ApiUser) error
}

type ApiUser struct {
	db.Dber
}

func NewApiUser(db db.Dber) ApiUserer {
	return &ApiUser{db}
}

func (auc *ApiUser) ApiUserCreate(m *model.ApiUser) error {
	return auc.Db().Create(m).Error
}

func (auc *ApiUser) ApiUserUpdate(m *model.ApiUser) error {
	return auc.Db().Updates(m).Error
}

func (auc *ApiUser) ApiUserGet(where map[string]interface{}) (*model.ApiUser, error) {
	var result model.ApiUser
	if err := auc.Db().Where(where).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
