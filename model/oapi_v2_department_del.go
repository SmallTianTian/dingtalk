package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiV2DepartmentDeleteRequest() *OapiV2DepartmentDeleteRequest {
	return &OapiV2DepartmentDeleteRequest{}
}

type OapiV2DepartmentDeleteRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiV2DepartmentDeleteResponse
	DeptId          int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiV2DepartmentDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiV2DepartmentDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiV2DepartmentDeleteRequest) SetDeptId(deptId2 int64) {
	this.DeptId = deptId2
}
func (this *OapiV2DepartmentDeleteRequest) GetDeptId() int64 {
	return this.DeptId
}
func (this *OapiV2DepartmentDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.v2.department.delete"
}
func (this *OapiV2DepartmentDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiV2DepartmentDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiV2DepartmentDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiV2DepartmentDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiV2DepartmentDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["dept_id"] = this.DeptId
	return txtParams
}
func (this *OapiV2DepartmentDeleteRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.DeptId, "deptId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiV2DepartmentDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiV2DepartmentDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiV2DepartmentDeleteResponse struct {
	taobao.TaobaoResponse
}
