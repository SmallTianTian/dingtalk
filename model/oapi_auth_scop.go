package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAuthScopesRequest() *OapiAuthScopesRequest {
	return &OapiAuthScopesRequest{}
}

type OapiAuthScopesRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAuthScopesResponse
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAuthScopesRequest) GetTopHttpMethod() string {
	return "GET"
}
func (this *OapiAuthScopesRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAuthScopesRequest) GetApiMethodName() string {
	return "dingtalk.oapi.auth.scopes"
}
func (this *OapiAuthScopesRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAuthScopesRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAuthScopesRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAuthScopesRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAuthScopesRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	return txtParams
}
func (this *OapiAuthScopesRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAuthScopesRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAuthScopesRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAuthScopesResponse struct {
	taobao.TaobaoResponse
	AuthOrgScopes  AuthOrgScopes `json:"auth_org_scopes,omitempty"`
	AuthUserField  []string      `json:"auth_user_field,omitempty"`
	ConditionField []string      `json:"condition_field,omitempty"`
	Errcode        int64         `json:"errcode,omitempty"`
	Errmsg         string        `json:"errmsg,omitempty"`
}
type AuthOrgScopes struct {
	AuthedDept []int64  `json:"authed_dept,omitempty"`
	AuthedUser []string `json:"authed_user,omitempty"`
}
