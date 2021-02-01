package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiProcessApproversForecastRequest() *OapiProcessApproversForecastRequest {
	return &OapiProcessApproversForecastRequest{}
}

type OapiProcessApproversForecastRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessApproversForecastResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiProcessApproversForecastRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessApproversForecastRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessApproversForecastRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiProcessApproversForecastRequest) GetRequest() string {
	return this.Request
}
func (this *OapiProcessApproversForecastRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.approvers.forecast"
}
func (this *OapiProcessApproversForecastRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessApproversForecastRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessApproversForecastRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessApproversForecastRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessApproversForecastRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiProcessApproversForecastRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiProcessApproversForecastRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessApproversForecastRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type FormComponentValueVo struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}
type ProcessForecastRequest struct {
	Agentid             int64                  `json:"agentid,omitempty"`
	FormComponentValues []FormComponentValueVo `json:"form_component_values,omitempty"`
	OriginatorDeptid    int64                  `json:"originator_deptid,omitempty"`
	OriginatorUserid    string                 `json:"originator_userid,omitempty"`
	ProcessCode         string                 `json:"process_code,omitempty"`
}
type OapiProcessApproversForecastResponse struct {
	taobao.TaobaoResponse
	Result []ProcessForecastResponse `json:"result,omitempty"`
}
type UserProfileVo struct {
	Name   string `json:"name,omitempty"`
	Userid string `json:"userid,omitempty"`
}
type ProcessForecastResponse struct {
	Approvers    []UserProfileVo `json:"approvers,omitempty"`
	NodeName     string          `json:"node_name,omitempty"`
	NodeShowName string          `json:"node_show_name,omitempty"`
}
