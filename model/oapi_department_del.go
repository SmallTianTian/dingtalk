package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiDepartmentDeleteRequest() *OapiDepartmentDeleteRequest {
	return &OapiDepartmentDeleteRequest{}
}

type OapiDepartmentDeleteRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiDepartmentDeleteResponse
	Id              string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiDepartmentDeleteRequest) GetTopHttpMethod() string {
	return "GET"
}
func (this *OapiDepartmentDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiDepartmentDeleteRequest) SetId(id2 string) {
	this.Id = id2
}
func (this *OapiDepartmentDeleteRequest) GetId() string {
	return this.Id
}
func (this *OapiDepartmentDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.department.delete"
}
func (this *OapiDepartmentDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiDepartmentDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiDepartmentDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiDepartmentDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiDepartmentDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["id"] = this.Id
	return txtParams
}
func (this *OapiDepartmentDeleteRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiDepartmentDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiDepartmentDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiDepartmentDeleteResponse struct {
	taobao.TaobaoResponse
}
