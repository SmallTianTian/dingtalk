package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiSmartworkHrmEmployeeFieldGrouplistRequest() *OapiSmartworkHrmEmployeeFieldGrouplistRequest {
	return &OapiSmartworkHrmEmployeeFieldGrouplistRequest{}
}

type OapiSmartworkHrmEmployeeFieldGrouplistRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSmartworkHrmEmployeeFieldGrouplistResponse
	Agentid         int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSmartworkHrmEmployeeFieldGrouplistRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSmartworkHrmEmployeeFieldGrouplistRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSmartworkHrmEmployeeFieldGrouplistRequest) SetAgentid(agentid2 int64) {
	this.Agentid = agentid2
}
func (this *OapiSmartworkHrmEmployeeFieldGrouplistRequest) GetAgentid() int64 {
	return this.Agentid
}
func (this *OapiSmartworkHrmEmployeeFieldGrouplistRequest) GetApiMethodName() string {
	return "dingtalk.oapi.smartwork.hrm.employee.field.grouplist"
}
func (this *OapiSmartworkHrmEmployeeFieldGrouplistRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSmartworkHrmEmployeeFieldGrouplistRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSmartworkHrmEmployeeFieldGrouplistRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSmartworkHrmEmployeeFieldGrouplistRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSmartworkHrmEmployeeFieldGrouplistRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentid"] = this.Agentid
	return txtParams
}
func (this *OapiSmartworkHrmEmployeeFieldGrouplistRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiSmartworkHrmEmployeeFieldGrouplistRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSmartworkHrmEmployeeFieldGrouplistRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSmartworkHrmEmployeeFieldGrouplistResponse struct {
	taobao.TaobaoResponse
	Errcode int64           `json:"errcode,omitempty"`
	Errmsg  string          `json:"errmsg,omitempty"`
	Result  []GroupMetaInfo `json:"result,omitempty"`
	Success bool            `json:"success,omitempty"`
}
type FieldMetaInfo struct {
	FieldCode  string `json:"field_code,omitempty"`
	FieldName  string `json:"field_name,omitempty"`
	FieldType  string `json:"field_type,omitempty"`
	OptionText string `json:"option_text,omitempty"`
}
type GroupMetaInfo struct {
	FieldList []FieldMetaInfo `json:"field_list,omitempty"`
	GroupId   string          `json:"group_id,omitempty"`
	HasDetail bool            `json:"has_detail,omitempty"`
}
