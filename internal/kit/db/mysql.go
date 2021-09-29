/*
 * @Author: your name
 * @Date: 2021-09-29 16:49:12
 * @LastEditTime: 2021-09-29 17:09:13
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \shuaoyoupin\internal\kit\db\mysql.go
 */
package db

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Dber interface {
	Db() *gorm.DB
	Connect(dsn string)
}

type Db struct {
	mysqlDb *gorm.DB
}

var (
	_instance Dber
	once      sync.Once
)

func InstanceDb() Dber {
	once.Do(func() {
		_instance = &Db{}
	})
	return _instance
}

func (d *Db) Connect(dsn string) {
	var err error

	d.mysqlDb, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Sprintf("init mysql connect failed error %v", err))
	}
}

func (d *Db) Db() *gorm.DB {
	if d.mysqlDb == nil {
		panic("mysql db onject is nil")
	}
	return d.mysqlDb
}
