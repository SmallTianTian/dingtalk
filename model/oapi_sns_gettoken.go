package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiSnsGettokenRequest() *OapiSnsGettokenRequest {
	return &OapiSnsGettokenRequest{}
}

type OapiSnsGettokenRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSnsGettokenResponse
	Appid           string
	Appsecret       string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSnsGettokenRequest) GetTopHttpMethod() string {
	return "GET"
}
func (this *OapiSnsGettokenRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSnsGettokenRequest) SetAppid(appid2 string) {
	this.Appid = appid2
}
func (this *OapiSnsGettokenRequest) GetAppid() string {
	return this.Appid
}
func (this *OapiSnsGettokenRequest) SetAppsecret(appsecret2 string) {
	this.Appsecret = appsecret2
}
func (this *OapiSnsGettokenRequest) GetAppsecret() string {
	return this.Appsecret
}
func (this *OapiSnsGettokenRequest) GetApiMethodName() string {
	return "dingtalk.oapi.sns.gettoken"
}
func (this *OapiSnsGettokenRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSnsGettokenRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSnsGettokenRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSnsGettokenRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSnsGettokenRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["appid"] = this.Appid
	txtParams["appsecret"] = this.Appsecret
	return txtParams
}
func (this *OapiSnsGettokenRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiSnsGettokenRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSnsGettokenRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSnsGettokenResponse struct {
	taobao.TaobaoResponse
	AccessToken string `json:"access_token,omitempty"`
}
