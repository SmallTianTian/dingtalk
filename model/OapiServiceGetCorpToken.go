package model

import (
	"encoding/json"

	"github.com/SmallTianTian/dingtalk/model/taobao"
)

type OapiServiceGetCorpTokenRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	cropID string

	Resp OapiServiceGetCorpTokenResponse
}

func (req *OapiServiceGetCorpTokenRequest) SetAuthCorpid(cropID string) {
	req.cropID = cropID
}

func (req *OapiServiceGetCorpTokenRequest) GetApiMethodName() string {
	return "dingtalk.oapi.service.get_corp_token"
}

func (req *OapiServiceGetCorpTokenRequest) GetTopApiCallType() string {
	return "dingtalk"
}

func (req *OapiServiceGetCorpTokenRequest) GetTextParams() map[string]string {
	return map[string]string{
		"auth_corpid": req.cropID,
	}
}
func (req *OapiServiceGetCorpTokenRequest) GetRespInstance() interface{} {
	return &req.Resp
}

func (req *OapiServiceGetCorpTokenRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &req.Resp.TaobaoResponse
}

func NewOapiServiceGetCorpTokenRequest() *OapiServiceGetCorpTokenRequest {
	return &OapiServiceGetCorpTokenRequest{}
}

type OapiServiceGetCorpTokenResponse struct {
	taobao.TaobaoResponse
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func (resp *OapiServiceGetCorpTokenResponse) GetBody() string {
	bs, _ := json.Marshal(resp)
	return string(bs)
}
