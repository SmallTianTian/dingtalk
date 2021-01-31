package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiV2DepartmentListparentbydeptRequest() *OapiV2DepartmentListparentbydeptRequest {
	return &OapiV2DepartmentListparentbydeptRequest{}
}

type OapiV2DepartmentListparentbydeptRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiV2DepartmentListparentbydeptResponse
	DeptId          int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiV2DepartmentListparentbydeptRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiV2DepartmentListparentbydeptRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiV2DepartmentListparentbydeptRequest) SetDeptId(deptId2 int64) {
	this.DeptId = deptId2
}
func (this *OapiV2DepartmentListparentbydeptRequest) GetDeptId() int64 {
	return this.DeptId
}
func (this *OapiV2DepartmentListparentbydeptRequest) GetApiMethodName() string {
	return "dingtalk.oapi.v2.department.listparentbydept"
}
func (this *OapiV2DepartmentListparentbydeptRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiV2DepartmentListparentbydeptRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiV2DepartmentListparentbydeptRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiV2DepartmentListparentbydeptRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiV2DepartmentListparentbydeptRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["dept_id"] = this.DeptId
	return txtParams
}
func (this *OapiV2DepartmentListparentbydeptRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.DeptId, "deptId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiV2DepartmentListparentbydeptRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiV2DepartmentListparentbydeptRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiV2DepartmentListparentbydeptResponse struct {
	taobao.TaobaoResponse
	Result DeptListParentByDeptIdResponse `json:"result,omitempty"`
}
type DeptListParentByDeptIdResponse struct {
	ParentIdList []int64 `json:"parent_id_list,omitempty"`
}
