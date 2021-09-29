/*
 * @Author: your name
 * @Date: 2021-09-29 17:01:22
 * @LastEditTime: 2021-09-29 17:11:33
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \shuaoyoupin\internal\kit\db\mysql_test.go
 */
package db

import "testing"

func TestConnect(t *testing.T) {
	InstanceDb().Connect("root:123456@tcp(127.0.0.1:3306)/game?charset=utf8mb4&parseTime=True&loc=Local")
	InstanceDb().Db()
}
