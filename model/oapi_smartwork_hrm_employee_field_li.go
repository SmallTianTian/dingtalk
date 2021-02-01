package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiSmartworkHrmEmployeeFieldListRequest() *OapiSmartworkHrmEmployeeFieldListRequest {
	return &OapiSmartworkHrmEmployeeFieldListRequest{}
}

type OapiSmartworkHrmEmployeeFieldListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartworkHrmEmployeeFieldListResponse
	Agentid         int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartworkHrmEmployeeFieldListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartworkHrmEmployeeFieldListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartworkHrmEmployeeFieldListRequest) SetAgentid(agentid2 int64) {
	this.Agentid = agentid2
}
func (this *OapiSmartworkHrmEmployeeFieldListRequest) GetAgentid() int64 {
	return this.Agentid
}
func (this *OapiSmartworkHrmEmployeeFieldListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartwork.hrm.employee.field.list"
}
func (this *OapiSmartworkHrmEmployeeFieldListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartworkHrmEmployeeFieldListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartworkHrmEmployeeFieldListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartworkHrmEmployeeFieldListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartworkHrmEmployeeFieldListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentid"] = this.Agentid
	return txtParams
}
func (this *OapiSmartworkHrmEmployeeFieldListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiSmartworkHrmEmployeeFieldListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartworkHrmEmployeeFieldListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartworkHrmEmployeeFieldListResponse struct {
	taobao.TaobaoResponse
	Result  []FieldMetaInfo `json:"result,omitempty"`
	Success bool            `json:"success,omitempty"`
}
