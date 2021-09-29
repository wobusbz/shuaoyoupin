package server

import (
	"shuaoyoupin/internal/kit/db"
	"shuaoyoupin/internal/model"

	"gorm.io/gorm"
)

type Relationer interface {
	RelationCreate(m *model.Relation) error
	//ApiUserUpdate(m *model.Relation) error
	RelationGet(where map[string]interface{}) (*model.Relation, error)
	//ApiUserList(m *model.Relation) error
}

type relation struct {
	db.Dber
}

func NewRelation(db db.Dber) Relationer {
	return &relation{db}
}

func (r *relation) RelationCreate(m *model.Relation) error {
	return r.Db().Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("`left` >= ?", m.Left).Update("left", gorm.Expr("`left` + ?", 2)).Error; err != nil {
			return err
		}
		if err := tx.Where("`right` >= ?", m.Right).Update("right", gorm.Expr("`right` + ?", 2)).Error; err != nil {
			return err
		}
		return r.Db().Create(&model.Relation{
			UserId:         m.UserId,
			ParentId:       m.ParentId,
			InvitationCode: m.InvitationCode,
			Level: 		  m.Level+1,
			Left:         m.Right,
			Right:        m.Right + 1,
		}).Error
	})
}

//func (r *relation) RelationUpdate(m *model.ApiUser) error {
//	return r.Db().Updates(m).Error
//}

func (r *relation) RelationGet(where map[string]interface{}) (*model.Relation, error) {
	var result model.Relation
	if err := r.Db().Where(where).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
