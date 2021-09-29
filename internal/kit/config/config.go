/*
 * @Author: your name
 * @Date: 2021-09-29 12:50:34
 * @LastEditTime: 2021-09-29 13:20:01
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \shuaoyoupin\internal\kit\config\config.go
 */
package config

import (
	"fmt"
	"sync"

	"github.com/BurntSushi/toml"
)

type (
	yingJiShuJuYunV2 struct {
		DbConnect string
		Port      uint
		AdminJwt  string `toml:"admin"`
		ApiJwt    string `toml:"api"`
	}
	Config struct {
		YingJiShuJuYunV2 yingJiShuJuYunV2 `toml:"yingjishujuyunV2"`
	}
)

var (
	_instance *Config
	once sync.Once
)

func InstanceConfig() *Config {
	once.Do(func() {
		_instance = newConfig()
	})
	return _instance
}

func newConfig() *Config {
	return &Config{YingJiShuJuYunV2: yingJiShuJuYunV2{
		Port:     8080,
		AdminJwt: "adminJwt",
		ApiJwt:   "apiJwt",
	}}
}

func (cf *Config) readConfig() {
	if _, err := toml.DecodeFile("../../../config/config.toml", cf); err != nil {
		panic(fmt.Sprintf("read config notfound file error %v", err))
	}
}
