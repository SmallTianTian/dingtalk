package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiDepartmentListParentDeptsRequest() *OapiDepartmentListParentDeptsRequest {
	return &OapiDepartmentListParentDeptsRequest{}
}

type OapiDepartmentListParentDeptsRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiDepartmentListParentDeptsResponse
	TopHttpMethod   string
	TopResponseType string
	UserId          string
}

func (this *OapiDepartmentListParentDeptsRequest) GetTopHttpMethod() string {
	return "GET"
}
func (this *OapiDepartmentListParentDeptsRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiDepartmentListParentDeptsRequest) SetUserId(userId2 string) {
	this.UserId = userId2
}
func (this *OapiDepartmentListParentDeptsRequest) GetUserId() string {
	return this.UserId
}
func (this *OapiDepartmentListParentDeptsRequest) GetApiMethodName() string {
	return "dingtalk.oapi.department.list_parent_depts"
}
func (this *OapiDepartmentListParentDeptsRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiDepartmentListParentDeptsRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiDepartmentListParentDeptsRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiDepartmentListParentDeptsRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiDepartmentListParentDeptsRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["userId"] = this.UserId
	return txtParams
}
func (this *OapiDepartmentListParentDeptsRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiDepartmentListParentDeptsRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiDepartmentListParentDeptsRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiDepartmentListParentDeptsResponse struct {
	taobao.TaobaoResponse
	Department string `json:"department,omitempty"`
}
