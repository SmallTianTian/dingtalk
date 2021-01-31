package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiOrgListshortcutRequest() *OapiOrgListshortcutRequest {
	return &OapiOrgListshortcutRequest{}
}

type OapiOrgListshortcutRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiOrgListshortcutResponse
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiOrgListshortcutRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiOrgListshortcutRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiOrgListshortcutRequest) GetApiMethodName() string {
	return "dingtalk.oapi.org.listshortcut"
}
func (this *OapiOrgListshortcutRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiOrgListshortcutRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiOrgListshortcutRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiOrgListshortcutRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiOrgListshortcutRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	return txtParams
}
func (this *OapiOrgListshortcutRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiOrgListshortcutRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiOrgListshortcutRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiOrgListshortcutResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
type OrgShortcutVO struct {
	Type  int64  `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}
