package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiV2DepartmentListsubidRequest() *OapiV2DepartmentListsubidRequest {
	return &OapiV2DepartmentListsubidRequest{}
}

type OapiV2DepartmentListsubidRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiV2DepartmentListsubidResponse
	DeptId          int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiV2DepartmentListsubidRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiV2DepartmentListsubidRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiV2DepartmentListsubidRequest) SetDeptId(deptId2 int64) {
	this.DeptId = deptId2
}
func (this *OapiV2DepartmentListsubidRequest) GetDeptId() int64 {
	return this.DeptId
}
func (this *OapiV2DepartmentListsubidRequest) GetApiMethodName() string {
	return "dingtalk.oapi.v2.department.listsubid"
}
func (this *OapiV2DepartmentListsubidRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiV2DepartmentListsubidRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiV2DepartmentListsubidRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiV2DepartmentListsubidRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiV2DepartmentListsubidRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["dept_id"] = this.DeptId
	return txtParams
}
func (this *OapiV2DepartmentListsubidRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiV2DepartmentListsubidRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiV2DepartmentListsubidRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiV2DepartmentListsubidResponse struct {
	taobao.TaobaoResponse
	Result DeptListSubIdResponse `json:"result,omitempty"`
}
type DeptListSubIdResponse struct {
	DeptIdList []int64 `json:"dept_id_list,omitempty"`
}
