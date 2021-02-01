package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiDepartmentListIdsRequest() *OapiDepartmentListIdsRequest {
	return &OapiDepartmentListIdsRequest{}
}

type OapiDepartmentListIdsRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiDepartmentListIdsResponse
	Id              string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiDepartmentListIdsRequest) GetTopHttpMethod() string {
	return "GET"
}
func (this *OapiDepartmentListIdsRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiDepartmentListIdsRequest) SetId(id2 string) {
	this.Id = id2
}
func (this *OapiDepartmentListIdsRequest) GetId() string {
	return this.Id
}
func (this *OapiDepartmentListIdsRequest) GetApiMethodName() string {
	return "dingtalk.oapi.department.list_ids"
}
func (this *OapiDepartmentListIdsRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiDepartmentListIdsRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiDepartmentListIdsRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiDepartmentListIdsRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiDepartmentListIdsRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["id"] = this.Id
	return txtParams
}
func (this *OapiDepartmentListIdsRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiDepartmentListIdsRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiDepartmentListIdsRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiDepartmentListIdsResponse struct {
	taobao.TaobaoResponse

	SubDeptIdList []int64 `json:"sub_dept_id_list,omitempty"`
}
