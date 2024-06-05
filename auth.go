package qq

import (
	"fmt"

	"github.com/GrayRabbit-FE/qq-miniprogram-sdk-golang/consts"
	"github.com/GrayRabbit-FE/qq-miniprogram-sdk-golang/types"
	"github.com/GrayRabbit-FE/qq-miniprogram-sdk-golang/utils"
)

func (cli *QQ) Login(code string) (*types.LoginResponse, error) {
	api := consts.BaseUrl + consts.Code2SessionEndPoint
	return cli.login(code, api)
}

func (cli *QQ) login(code, api string) (*types.LoginResponse, error) {
	url := fmt.Sprintf("%s?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", consts.BaseUrl+consts.Code2SessionEndPoint, cli.AppId, cli.AppSecret, code)
	fmt.Printf("%s?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", consts.BaseUrl+consts.Code2SessionEndPoint, cli.AppId, cli.AppSecret, code)
	resp := new(types.LoginResponse)
	err := utils.Get(url, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
