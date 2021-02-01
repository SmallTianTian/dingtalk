package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiServiceGetUnactiveCorpRequest() *OapiServiceGetUnactiveCorpRequest {
	return &OapiServiceGetUnactiveCorpRequest{}
}

type OapiServiceGetUnactiveCorpRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiServiceGetUnactiveCorpResponse
	AppId           int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiServiceGetUnactiveCorpRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiServiceGetUnactiveCorpRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiServiceGetUnactiveCorpRequest) SetAppId(appId2 int64) {
	this.AppId = appId2
}
func (this *OapiServiceGetUnactiveCorpRequest) GetAppId() int64 {
	return this.AppId
}
func (this *OapiServiceGetUnactiveCorpRequest) GetApiMethodName() string {
	return "dingtalk.oapi.service.get_unactive_corp"
}
func (this *OapiServiceGetUnactiveCorpRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiServiceGetUnactiveCorpRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiServiceGetUnactiveCorpRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiServiceGetUnactiveCorpRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiServiceGetUnactiveCorpRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["app_id"] = this.AppId
	return txtParams
}
func (this *OapiServiceGetUnactiveCorpRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiServiceGetUnactiveCorpRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiServiceGetUnactiveCorpRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiServiceGetUnactiveCorpResponse struct {
	taobao.TaobaoResponse
	AppId    int64    `json:"app_id,omitempty"`
	CorpList []string `json:"corp_list,omitempty"`

	HasMore bool `json:"has_more,omitempty"`
}
