package qq

import (
	"github.com/GrayRabbit-FE/qq-miniprogram-sdk-golang/types"
	"github.com/GrayRabbit-FE/qq-miniprogram-sdk-golang/utils"
)

type QQ struct {
	AppId       string
	AppSecret   string
	Cache       utils.Cache
	AccessToken utils.AccessToken
}

func NewQQ(config *types.QQAppletConfig) *QQ {
	// 如果存在cache组件就实例化一个内存缓存
	if config.Cache == nil {
		config.Cache = utils.NewMemory()
	}
	// 如果未设置token管理 就使用默认的
	if config.AccessToken == nil {
		config.AccessToken = utils.NewAccessTokenManager(config.AppId, config.AppSecret, config.Cache)
	}

	return &QQ{
		AppId:       config.AppId,
		AppSecret:   config.AppSecret,
		Cache:       config.Cache,
		AccessToken: config.AccessToken,
	}
}
