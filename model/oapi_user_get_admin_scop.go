package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiUserGetAdminScopeRequest() *OapiUserGetAdminScopeRequest {
	return &OapiUserGetAdminScopeRequest{}
}

type OapiUserGetAdminScopeRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiUserGetAdminScopeResponse
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiUserGetAdminScopeRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiUserGetAdminScopeRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiUserGetAdminScopeRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiUserGetAdminScopeRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiUserGetAdminScopeRequest) GetApiMethodName() string {
	return "dingtalk.oapi.user.get_admin_scope"
}
func (this *OapiUserGetAdminScopeRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiUserGetAdminScopeRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiUserGetAdminScopeRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiUserGetAdminScopeRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiUserGetAdminScopeRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiUserGetAdminScopeRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Userid, taobao.DATA_OUTGOING_USER_ID); err != nil {
		return err
	}
	return nil
}
func (this *OapiUserGetAdminScopeRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiUserGetAdminScopeRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiUserGetAdminScopeResponse struct {
	taobao.TaobaoResponse
	DeptIds []int64 `json:"dept_ids,omitempty"`
	Errcode int64   `json:"errcode,omitempty"`
	Errmsg  string  `json:"errmsg,omitempty"`
}
