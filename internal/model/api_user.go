package model

import "gorm.io/gorm"

type ApiUser struct {
	gorm.Model
	Mobile         string `gorm:"column:mobile;comment:'手机号码';type:varchar(255)"`
	PassWord       string `gorm:"column:password;comment:'密码';type:varchar(255)"`
	AssetsPassWord string `gorm:"column:assets_pass_word;comment:'资金密码';type:varchar(255)"`
	RealName       string `gorm:"column:real_name;comment:'真实姓名';type:varchar(255)"`
	CardNo         string `gorm:"column:card_no;comment:'身份证号码';type:varchar(255)"`
	CardPositive   string `gorm:"column:card_positive;comment:'身份证正面';type:varchar(255)"`
	CardReverse    string `gorm:"column:card_reverse;comment:'身份证反面';type:varchar(255)"`
	RealNameAuth   int    `gorm:"column:real_name_auth;comment:'实名认证状态';type:tinyint"`
	Status         bool   `gorm:"column:status;comment:'启停用';type:tinyint"`
}
