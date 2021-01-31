package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiUserSeniorWhitelistSetRequest() *OapiUserSeniorWhitelistSetRequest {
	return &OapiUserSeniorWhitelistSetRequest{}
}

type OapiUserSeniorWhitelistSetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                   OapiUserSeniorWhitelistSetResponse
	SeniorWhitelistRequest string
	TopHttpMethod          string
	TopResponseType        string
}

func (this *OapiUserSeniorWhitelistSetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiUserSeniorWhitelistSetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiUserSeniorWhitelistSetRequest) SetSeniorWhitelistRequest(seniorWhitelistRequest2 string) {
	this.SeniorWhitelistRequest = seniorWhitelistRequest2
}
func (this *OapiUserSeniorWhitelistSetRequest) GetSeniorWhitelistRequest() string {
	return this.SeniorWhitelistRequest
}
func (this *OapiUserSeniorWhitelistSetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.user.senior.whitelist.set"
}
func (this *OapiUserSeniorWhitelistSetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiUserSeniorWhitelistSetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiUserSeniorWhitelistSetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiUserSeniorWhitelistSetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiUserSeniorWhitelistSetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["senior_whitelist_request"] = this.SeniorWhitelistRequest
	return txtParams
}
func (this *OapiUserSeniorWhitelistSetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiUserSeniorWhitelistSetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiUserSeniorWhitelistSetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type SeniorWhitelistRequest struct {
	SeniorPermitDeptids []int64  `json:"senior_permit_deptids,omitempty"`
	SeniorPermitRoleids []int64  `json:"senior_permit_roleids,omitempty"`
	SeniorPermitUserids []string `json:"senior_permit_userids,omitempty"`
	SeniorUserid        string   `json:"senior_userid,omitempty"`
}
type OapiUserSeniorWhitelistSetResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Success bool   `json:"success,omitempty"`
}
