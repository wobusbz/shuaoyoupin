package model

import "gorm.io/gorm"

type Relation struct {
	gorm.Model
	UserId         int `gorm:"column:user_id;comment:'用户ID';type:int(10)"`
	ParentId       int `gorm:"column:user_parent_id;comment:'父级ID';type:int(10)"`
	InvitationCode int `gorm:"column:user_parent_id;comment:'父级ID';type:int(10)"`
	Level          int `gorm:"column:level;comment:'层级';type:int(10)"`
	Left           int `gorm:"column:left;comment:'左节点';type:int(10)"`
	Right          int `gorm:"column:right;comment:'右节点';int(10)"`
}
