/*
 * @Author: your name
 * @Date: 2021-09-29 13:15:34
 * @LastEditTime: 2021-09-29 13:16:34
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \shuaoyoupin\internal\kit\config\config_test.go
 */
package config

import "testing"

func TestReadConfig(t *testing.T) {
	var cf = &Config{}
	cf.readConfig()
	t.Log(cf.YingJiShuJuYunV2.Port)
}
