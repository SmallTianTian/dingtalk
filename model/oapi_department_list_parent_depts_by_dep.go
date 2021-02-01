package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiDepartmentListParentDeptsByDeptRequest() *OapiDepartmentListParentDeptsByDeptRequest {
	return &OapiDepartmentListParentDeptsByDeptRequest{}
}

type OapiDepartmentListParentDeptsByDeptRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiDepartmentListParentDeptsByDeptResponse
	Id              string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiDepartmentListParentDeptsByDeptRequest) GetTopHttpMethod() string {
	return "GET"
}
func (this *OapiDepartmentListParentDeptsByDeptRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiDepartmentListParentDeptsByDeptRequest) SetId(id2 string) {
	this.Id = id2
}
func (this *OapiDepartmentListParentDeptsByDeptRequest) GetId() string {
	return this.Id
}
func (this *OapiDepartmentListParentDeptsByDeptRequest) GetApiMethodName() string {
	return "dingtalk.oapi.department.list_parent_depts_by_dept"
}
func (this *OapiDepartmentListParentDeptsByDeptRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiDepartmentListParentDeptsByDeptRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiDepartmentListParentDeptsByDeptRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiDepartmentListParentDeptsByDeptRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiDepartmentListParentDeptsByDeptRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["id"] = this.Id
	return txtParams
}
func (this *OapiDepartmentListParentDeptsByDeptRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiDepartmentListParentDeptsByDeptRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiDepartmentListParentDeptsByDeptRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiDepartmentListParentDeptsByDeptResponse struct {
	taobao.TaobaoResponse

	ParentIds []int64 `json:"parentIds,omitempty"`
}
