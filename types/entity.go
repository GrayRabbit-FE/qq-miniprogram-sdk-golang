package types

import "github.com/GrayRabbit-FE/qq-miniprogram-sdk-golang/utils"

type QQAppletConfig struct {
	AppId       string      // 快手小程序的appid
	AppSecret   string      // 快手小程序的app secret
	Cache       utils.Cache // 基础的缓存接口
	AccessToken utils.AccessToken
}
