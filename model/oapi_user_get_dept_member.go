package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiUserGetDeptMemberRequest() *OapiUserGetDeptMemberRequest {
	return &OapiUserGetDeptMemberRequest{}
}

type OapiUserGetDeptMemberRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiUserGetDeptMemberResponse
	DeptId          string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiUserGetDeptMemberRequest) GetTopHttpMethod() string {
	return "GET"
}
func (this *OapiUserGetDeptMemberRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiUserGetDeptMemberRequest) SetDeptId(deptId2 string) {
	this.DeptId = deptId2
}
func (this *OapiUserGetDeptMemberRequest) GetDeptId() string {
	return this.DeptId
}
func (this *OapiUserGetDeptMemberRequest) GetApiMethodName() string {
	return "dingtalk.oapi.user.getDeptMember"
}
func (this *OapiUserGetDeptMemberRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiUserGetDeptMemberRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiUserGetDeptMemberRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiUserGetDeptMemberRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiUserGetDeptMemberRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["deptId"] = this.DeptId
	return txtParams
}
func (this *OapiUserGetDeptMemberRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiUserGetDeptMemberRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiUserGetDeptMemberRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiUserGetDeptMemberResponse struct {
	taobao.TaobaoResponse
	UserIds []string `json:"userIds,omitempty"`
}
