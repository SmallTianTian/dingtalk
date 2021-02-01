package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiServiceReauthCorpRequest() *OapiServiceReauthCorpRequest {
	return &OapiServiceReauthCorpRequest{}
}

type OapiServiceReauthCorpRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiServiceReauthCorpResponse
	AppId           string
	CorpidList      []string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiServiceReauthCorpRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiServiceReauthCorpRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiServiceReauthCorpRequest) SetAppId(appId2 string) {
	this.AppId = appId2
}
func (this *OapiServiceReauthCorpRequest) GetAppId() string {
	return this.AppId
}
func (this *OapiServiceReauthCorpRequest) SetCorpidList(corpidList2 []string) {
	this.CorpidList = corpidList2
}
func (this *OapiServiceReauthCorpRequest) GetCorpidList() []string {
	return this.CorpidList
}
func (this *OapiServiceReauthCorpRequest) GetApiMethodName() string {
	return "dingtalk.oapi.service.reauth_corp"
}
func (this *OapiServiceReauthCorpRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiServiceReauthCorpRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiServiceReauthCorpRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiServiceReauthCorpRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiServiceReauthCorpRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["app_id"] = this.AppId
	txtParams["corpid_list"] = this.CorpidList
	return txtParams
}
func (this *OapiServiceReauthCorpRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiServiceReauthCorpRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiServiceReauthCorpRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiServiceReauthCorpResponse struct {
	taobao.TaobaoResponse
}
